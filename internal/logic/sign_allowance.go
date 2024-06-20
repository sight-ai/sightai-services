package logic

import (
	"context"
	"fmt"
	"github.com/capybaralabs-xyz/sightai-services/internal/constants"
	"github.com/capybaralabs-xyz/sightai-services/internal/constants/types"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities/models"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/rest"
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

// SignAllowance
//  1. if increase allowance
//     1.1. deduct available amount
//     1.2. increase hold amount
//     1.3. upsert allowance
//  2. if increase allowance
//     2.1. deduct hold amount
//     2.2. increase available amount
//     2.3. upsert allowance
func SignAllowance(ctx context.Context, account *entities.Account, req *models.SignAllowanceRequest) (*models.SimpleMessageResponse, error) {
	newAllowance, err := decimal.NewFromString(req.Allowance)
	if err != nil {
		return nil, rest.ErrBadRequest("invalid allowance")
	}

	err = entities.Txn.Tx(ctx, func(txnCtx context.Context) error {
		toAccount, err := entities.AccountDao.Get(txnCtx, uint(req.ToAccountId))
		if err != nil {
			return rest.ErrFromGormError(txnCtx, err, "AccountDao.Get failed")
		}

		existingAllowance, err := entities.AllowanceDao.GetByFromToAccountID(txnCtx, account.ID, uint(req.ToAccountId))
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				// create new Allowance
				_, err = entities.AllowanceDao.UpsertAllowance(txnCtx, account.ID, uint(req.ToAccountId), uint(req.Version), newAllowance)
				if err != nil {
					return rest.ErrFromGormError(txnCtx, err, "AccountDao.UpsertAllowance failed")
				}

				// increase allowance -> hold from_account balance, increase to_account hold
				err = entities.AccountDao.HoldBalance(txnCtx, account, newAllowance, types.TransactionTypeAllowance, fmt.Sprintf("allowance from %d to %d version %d", account.ID, req.ToAccountId, req.Version))
				if err != nil {
					return rest.ErrFromGormError(txnCtx, err, "failed AccountDao.HoldBalance")
				}

				err = entities.AccountDao.IncreaseHoldBalance(txnCtx, toAccount, newAllowance, types.TransactionTypeAllowance, fmt.Sprintf("allowance from %d to %d version %d", account.ID, req.ToAccountId, req.Version))
				if err != nil {
					return rest.ErrFromGormError(txnCtx, err, "failed AccountDao.HoldBalance")
				}

			} else {
				return rest.ErrFromGormError(txnCtx, err, "AllowanceDao.GetByFromToAccountID failed")
			}
		} else {
			// update allowance
			if newAllowance.GreaterThan(existingAllowance.Allowance) {
				// increase allowance -> hold from_account balance, increase to_account hold
				err = entities.AccountDao.HoldBalance(txnCtx, account, newAllowance, types.TransactionTypeAllowance, fmt.Sprintf("allowance from %d to %d version %d", account.ID, req.ToAccountId, req.Version))
				if err != nil {
					return rest.ErrFromGormError(txnCtx, err, "failed AccountDao.HoldBalance")
				}

				err = entities.AccountDao.IncreaseHoldBalance(txnCtx, toAccount, newAllowance, types.TransactionTypeAllowance, fmt.Sprintf("allowance from %d to %d version %d", account.ID, req.ToAccountId, req.Version))
				if err != nil {
					return rest.ErrFromGormError(txnCtx, err, "failed AccountDao.HoldBalance")
				}

				_, err = entities.AllowanceDao.UpsertAllowance(txnCtx, account.ID, uint(req.ToAccountId), uint(req.Version), newAllowance)
				if err != nil {
					return rest.ErrFromGormError(txnCtx, err, "AllowanceDao.UpsertAllowance failed")
				}

			} else if newAllowance.LessThan(existingAllowance.Allowance) {
				// decrease allowance -> unhold from_account balance, decrease to_account hold
				err = entities.AccountDao.UnholdBalance(txnCtx, account, newAllowance, types.TransactionTypeAllowance, fmt.Sprintf("allowance from %d to %d version %d", account.ID, req.ToAccountId, req.Version))
				if err != nil {
					return rest.ErrFromGormError(txnCtx, err, "failed AccountDao.HoldBalance")
				}

				err = entities.AccountDao.DecreaseHoldBalance(txnCtx, toAccount, newAllowance, types.TransactionTypeAllowance, fmt.Sprintf("allowance from %d to %d version %d", account.ID, req.ToAccountId, req.Version))
				if err != nil {
					return rest.ErrFromGormError(txnCtx, err, "failed AccountDao.HoldBalance")
				}

				_, err = entities.AllowanceDao.UpsertAllowance(txnCtx, account.ID, uint(req.ToAccountId), uint(req.Version), newAllowance)
				if err != nil {
					return rest.ErrFromGormError(txnCtx, err, "AllowanceDao.UpsertAllowance failed")
				}
			} else {
				// keep untouched
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &models.SimpleMessageResponse{
		Message: constants.OKResponse,
	}, nil
}
