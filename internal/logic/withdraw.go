package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/capybaralabs-xyz/sightai-services/internal/constants"
	"github.com/capybaralabs-xyz/sightai-services/internal/constants/types"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities/models"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/lib_signature"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/log"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/rest"
	"github.com/shopspring/decimal"
)

// Withdraw
// 1. deduct available amount
// 2. sign & issue ticket
func Withdraw(ctx context.Context, account *entities.Account, sig string, req *models.WithdrawRequest) (*models.WithdrawResponse, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, rest.ErrBadRequest("cannot marshal req")
	}
	valid, err := lib_signature.VerifyDataSig(account.Address, sig, b)
	if !valid {
		return nil, rest.ErrBadRequest("invalid signature")
	}
	if err != nil {
		return nil, rest.ErrInternalServerError(ctx, err)
	}

	var signedTicket string
	var sizeInWei decimal.Decimal
	var nonce uint

	amount, err := decimal.NewFromString(req.Amount)
	if err != nil {
		return nil, rest.ErrBadRequest(fmt.Sprintf("invalid withdraw amount %s", req.Amount))
	}

	err = entities.Txn.Tx(ctx, func(txnCtx context.Context) error {
		nonce, err = entities.AccountDao.GetAndIncreaseUserNonce(ctx, account.ID)
		if err != nil {
			return rest.ErrFromGormError(txnCtx, err, "failed UserDao.GetAndIncreaseUserNonce")
		}

		account, err = entities.AccountDao.Get(txnCtx, account.ID)
		if err != nil {
			log.Error(ctx).Err(err).Msg("AccountDao.Get failed")
			return err
		}

		// validate and hold account balance
		err = entities.AccountDao.HoldBalance(txnCtx, account, amount, types.TransactionTypeWithdraw, fmt.Sprintf("withdraw nonce %d", nonce))
		if err != nil {
			return rest.ErrFromGormError(txnCtx, err, "failed AccountDao.HoldBalance")
		}

		sizeInWei = amount.Mul(decimal.New(1, int32(constants.SightTokenDecimal)))

		signedTicket, err = lib_signature.GetSignedWithdrawTicket(nonce, sizeInWei.String())
		if err != nil {
			return rest.ErrFromGormError(txnCtx, err, "failed withdraw.GetSignedWithdrawTicket")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &models.WithdrawResponse{
		Sig:    signedTicket,
		Nonce:  int32(nonce),
		Amount: sizeInWei.String(),
	}, nil
}
