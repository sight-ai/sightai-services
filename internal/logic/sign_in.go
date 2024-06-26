package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/capybaralabs-xyz/sightai-services/internal/constants"
	"github.com/capybaralabs-xyz/sightai-services/internal/constants/types"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities/models"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/comm_utils"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/jwt_auth"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/lib_signature"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/rest"
)

// SignIn - give `NewUserBonus` balance to all new users
func SignIn(ctx context.Context, req *models.SignInRequest, sig string) (*models.SignInResponse, error) {
	addr, err := comm_utils.ToEthAddress(req.Address)
	if err != nil {
		return nil, rest.ErrBadRequest("invalid wallet address")
	}

	b, err := json.Marshal(req)
	if err != nil {
		return nil, rest.ErrBadRequest("cannot marshal req")
	}
	valid, err := lib_signature.VerifyDataSig(addr, sig, b)
	if !valid {
		return nil, rest.ErrBadRequest("invalid signature")
	}
	if err != nil {
		return nil, rest.ErrInternalServerError(ctx, err)
	}

	account := &entities.Account{
		Address: addr,
		Role:    types.AccountRoleUser.String(),
	}

	isNewUser, err := entities.AccountDao.CreateOrGetByAddress(ctx, account)
	if err != nil {
		return nil, rest.ErrFromGormError(ctx, err, "AccountDao.CreateOrGetByAddress failed")
	}
	if isNewUser {
		err = entities.Txn.Tx(ctx, func(txnCtx context.Context) error {
			err = entities.AccountDao.AddBalance(txnCtx, account, constants.NewUserBonus, types.TransactionTypeDeposit, "new user bonus")
			if err != nil {
				return rest.ErrFromGormError(ctx, err, "AccountDao.AddBalance failed")
			}
			// TODO: do not give gateway allowance by default in the future
			gateway, err := entities.GatewayDao.GetById(txnCtx, constants.DefaultGatewayID)
			if err != nil {
				return rest.ErrFromGormError(txnCtx, err, "GatewayDao.GetById failed")
			}
			gatewayAccount, err := entities.AccountDao.GetByAddress(txnCtx, gateway.Address)
			if err != nil {
				return rest.ErrFromGormError(txnCtx, err, "AccountDao.GetByAddress failed")
			}

			_, err = entities.AllowanceDao.UpsertAllowance(txnCtx, account.ID, gatewayAccount.ID, 1, constants.NewUserBonus)
			if err != nil {
				return rest.ErrFromGormError(txnCtx, err, "AccountDao.UpsertAllowance failed")
			}

			// increase allowance -> hold from_account balance, increase to_account hold
			err = entities.AccountDao.HoldBalance(txnCtx, account, constants.NewUserBonus, types.TransactionTypeAllowance, fmt.Sprintf("allowance from %d to %d version %d", account.ID, gatewayAccount.ID, 1))
			if err != nil {
				return rest.ErrFromGormError(txnCtx, err, "failed AccountDao.HoldBalance")
			}

			err = entities.AccountDao.IncreaseHoldBalance(txnCtx, gatewayAccount, constants.NewUserBonus, types.TransactionTypeAllowance, fmt.Sprintf("allowance from %d to %d version %d", account.ID, gatewayAccount.ID, 1))
			if err != nil {
				return rest.ErrFromGormError(txnCtx, err, "failed AccountDao.HoldBalance")
			}

			return nil
		})
		if err != nil {
			return nil, err
		}
	}

	jwtToken, err := jwt_auth.GenerateJwtFromAccount(account, req.Domain)
	if err != nil {
		return nil, rest.ErrInternalServerError(ctx, err)
	}

	return &models.SignInResponse{
		UserToken: jwtToken,
	}, nil
}
