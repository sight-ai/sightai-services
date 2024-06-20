package services

import (
	"context"
	"fmt"
	"github.com/capybaralabs-xyz/sightai-services/internal/constants"
	"github.com/capybaralabs-xyz/sightai-services/internal/constants/types"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/comm_utils"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/data/mysql"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/lib_subgraph"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/log"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/rest"
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
	"strconv"
	"time"
)

func PullSubgraph(ctx context.Context) {
	err := pullDepositEvents(ctx)
	if err != nil {
		log.Error(ctx).Err(err).Msg("failed PullSubgraph pullDepositEvents")
	}

	// wait at least a block time
	// make sure all deposit are executed before withdraws
	time.Sleep(time.Millisecond * 500)

	err = pullWithdrawEvents(ctx)
	if err != nil {
		log.Error(ctx).Err(err).Msg("failed PullSubgraph pullWithdrawEvents")
	}
}

func pullDepositEvents(ctx context.Context) error {
	subgraphEvents := []lib_subgraph.DepositEvent{}

	for _, c := range lib_subgraph.Clients {
		ts, err := entities.DepositEventDao.GetLatestTimestamp(ctx)
		if err != nil {
			log.Error(ctx).Err(err).Msg("failed pullDepositEvents")
			continue
		}

		newDepositEvents, err := lib_subgraph.GetDepositEventsAfter(ctx, c, ts)
		if err != nil {
			log.Error(ctx).Err(err).Msg("failed GetDepositEventsAfter")
			continue
		}

		subgraphEvents = append(subgraphEvents, newDepositEvents...)

		log.Info(ctx).Msgf("pulled %d new deposit event(s) from %s", len(newDepositEvents), c.Node.ChainName)
	}

	for _, e := range subgraphEvents {
		fromAddr, err := comm_utils.ToEthAddress(e.From)
		if err != nil {
			log.Error(ctx).Err(err).Msgf("failed ToEthAddress %s", e.From)
			continue
		}
		toAddr, err := comm_utils.ToEthAddress(e.To)
		if err != nil {
			log.Error(ctx).Err(err).Msgf("failed ToEthAddress %s", e.To)
			continue
		}

		timestamp, _ := strconv.Atoi(e.Timestamp)

		event := &entities.DepositEvent{
			TxnHash:        e.Id,
			BlockTimestamp: time.Unix(int64(timestamp), 0),
			FromAddr:       fromAddr,
			ToAddr:         toAddr,
			Amount:         e.Amount,
		}
		err = entities.Txn.Tx(ctx, func(txnCtx context.Context) error {
			err = entities.DepositEventDao.Create(txnCtx, event)
			if err != nil {
				if mysql.GetGormErrorCode(err) == mysql.GormErrorCodeDuplicateEntry || err == mysql.ErrDuplicatedKey {
					return nil
				}
				log.Error(txnCtx).Err(err).Msg("failed DepositEventDao.BatchCreate")
				return err
			}

			// increase account balance
			account := &entities.Account{
				Address: toAddr,
			}

			wei, err := decimal.NewFromString(e.Amount)
			if err != nil {
				log.Error(txnCtx).Err(err).Msgf("parse amount failed %s", e.Amount)
				return err
			}

			depositAmount := wei.Mul(decimal.New(1, int32(-1*constants.SightTokenDecimal)))
			err = entities.AccountDao.AddBalance(txnCtx, account, depositAmount, types.TransactionTypeDepositEvent, fmt.Sprintf("deposit event %d", event.ID))
			if err != nil {
				log.Error(txnCtx).Err(err).Msgf("failed AccountDao.AddBalance, account: %+v, event: %+v", account, e)
				return err
			}
			account, _ = entities.AccountDao.GetByAddress(txnCtx, account.Address)

			// TODO: do not give gateway allowance by default in the future
			gateway, err := entities.GatewayDao.GetById(txnCtx, constants.DefaultGatewayID)
			if err != nil {
				return rest.ErrFromGormError(txnCtx, err, "GatewayDao.GetById failed")
			}
			gatewayAccount, err := entities.AccountDao.GetByAddress(txnCtx, gateway.Address)
			if err != nil {
				return rest.ErrFromGormError(txnCtx, err, "AccountDao.GetByAddress failed")
			}

			existingAllowance, err := entities.AllowanceDao.GetByFromToAccountID(txnCtx, account.ID, gatewayAccount.ID)
			if err != nil {
				if gorm.IsRecordNotFoundError(err) {
					existingAllowance, err = entities.AllowanceDao.UpsertAllowance(txnCtx, account.ID, gatewayAccount.ID, 1, depositAmount)
					if err != nil {
						return rest.ErrFromGormError(txnCtx, err, "AccountDao.UpsertAllowance failed")
					}
				} else {
					return rest.ErrFromGormError(txnCtx, err, "AccountDao.GetByFromToAccountID failed")
				}
			} else {
				existingAllowance, err = entities.AllowanceDao.UpsertAllowance(txnCtx, account.ID, gatewayAccount.ID, existingAllowance.Version+1, existingAllowance.Allowance.Add(depositAmount))
				if err != nil {
					return rest.ErrFromGormError(txnCtx, err, "AccountDao.UpsertAllowance failed")
				}
			}

			// increase allowance -> hold from_account balance, increase to_account hold
			err = entities.AccountDao.HoldBalance(txnCtx, account, depositAmount, types.TransactionTypeAllowance, fmt.Sprintf("allowance from %d to %d version %d", account.ID, gatewayAccount.ID, existingAllowance.Version+1))
			if err != nil {
				return rest.ErrFromGormError(txnCtx, err, "failed AccountDao.HoldBalance")
			}

			err = entities.AccountDao.IncreaseHoldBalance(txnCtx, gatewayAccount, depositAmount, types.TransactionTypeAllowance, fmt.Sprintf("allowance from %d to %d version %d", account.ID, gatewayAccount.ID, existingAllowance.Version+1))
			if err != nil {
				return rest.ErrFromGormError(txnCtx, err, "failed AccountDao.HoldBalance")
			}

			return nil
		})

		if err != nil {
			log.Error(ctx).Err(err).Msgf("failed exec DepositEvent: %+v", e)
		}
	}

	return nil
}

func pullWithdrawEvents(ctx context.Context) error {
	subgraphEvents := []lib_subgraph.WithdrawEvent{}

	for _, c := range lib_subgraph.Clients {

		ts, err := entities.WithdrawEventDao.GetLatestTimestamp(ctx)
		if err != nil {
			log.Error(ctx).Err(err).Msg("failed pullWithdrawEvents")
			continue
		}

		newWithdrawEvents, err := lib_subgraph.GetWithdrawEventsAfter(ctx, c, ts)
		if err != nil {
			log.Error(ctx).Err(err).Msg("failed GetWithdrawEventsAfter")
			continue
		}

		subgraphEvents = append(subgraphEvents, newWithdrawEvents...)
		log.Info(ctx).Msgf("pulled %d new withdraw event(s) from %s", len(newWithdrawEvents), c.Node.ChainName)
	}

	for _, e := range subgraphEvents {
		toAddr, err := comm_utils.ToEthAddress(e.To)
		if err != nil {
			log.Error(ctx).Err(err).Msgf("failed ToEthAddress %s", e.To)
			continue
		}

		err = entities.Txn.Tx(ctx, func(txnCtx context.Context) error {
			timestamp, _ := strconv.Atoi(e.Timestamp)

			event := &entities.WithdrawEvent{
				TxnHash:        e.Id,
				ToAddr:         toAddr,
				Amount:         e.Amount,
				Nonce:          uint(e.Nonce),
				BlockTimestamp: time.Unix(int64(timestamp), 0),
			}

			err = entities.WithdrawEventDao.Create(txnCtx, event)
			if err != nil {
				if mysql.GetGormErrorCode(err) == mysql.GormErrorCodeDuplicateEntry || err == mysql.ErrDuplicatedKey {
					return nil
				}
				log.Error(txnCtx).Err(err).Msg("failed WithdrawEventDao.BatchCreate")
				return err
			}

			// decrease hold balance
			account := &entities.Account{
				Address: toAddr,
			}
			_, err = entities.AccountDao.CreateOrGetByAddress(txnCtx, account)

			wei, err := decimal.NewFromString(e.Amount)
			if err != nil {
				log.Error(txnCtx).Err(err).Msgf("parse amount failed %s", e.Amount)
				return err
			}

			err = entities.AccountDao.DecreaseHoldBalance(txnCtx, account, wei.Mul(decimal.New(1, int32(-1*constants.SightTokenDecimal))), types.TransactionTypeWithdrawEvent, fmt.Sprintf("withdraw event %d", event.ID))
			if err != nil {
				log.Error(txnCtx).Err(err).Msgf("failed AccountDao.DecreaseHoldBalance, account: %+v, event: %+v", account, e)
				return err
			}

			return nil
		})
		if err != nil {
			log.Error(ctx).Err(err).Msgf("failed exec WithdrawEvent: %+v", e)
		}
	}

	return nil

}
