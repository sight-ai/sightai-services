package logic

import (
	"context"
	"encoding/json"
	"github.com/capybaralabs-xyz/sightai-services/internal/constants"
	"github.com/capybaralabs-xyz/sightai-services/internal/constants/types"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities/models"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/lib_signature"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/rest"
	"github.com/shopspring/decimal"
)

// GatewayWithdraw
// 1. fetch all receipts
// 2. calculate withdraw amount
// 3. update receipt status to paid
// 4. sign & issue ticket
func GatewayWithdraw(ctx context.Context, gateway *entities.Account, sig string, req *models.GatewayWithdrawRequest) (*models.WithdrawResponse, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, rest.ErrBadRequest("cannot marshal req")
	}
	valid, err := lib_signature.VerifyDataSig(gateway.Address, sig, b)
	if !valid {
		return nil, rest.ErrBadRequest("invalid signature")
	}
	if err != nil {
		return nil, rest.ErrInternalServerError(ctx, err)
	}

	var signedTicket string
	var sizeInWei decimal.Decimal
	var nonce uint
	err = entities.Txn.Tx(ctx, func(txnCtx context.Context) error {
		receipts, err := entities.ReceiptDao.Gets(ctx, req.ReceiptIds)
		if err != nil {
			return rest.ErrFromGormError(txnCtx, err, "ReceiptDao.Gets failed")
		}

		var totalAmount decimal.Decimal
		for _, receipt := range receipts {
			totalAmount = totalAmount.Add(receipt.Cost)
			receipt.Status = types.ReceiptStatusPaid.String()
			err = entities.ReceiptDao.Update(txnCtx, receipt)
			if err != nil {
				return rest.ErrFromGormError(txnCtx, err, "ReceiptDao.Update failed")
			}
		}

		nonce, err = entities.AccountDao.GetAndIncreaseUserNonce(txnCtx, gateway.ID)
		if err != nil {
			return rest.ErrFromGormError(txnCtx, err, "failed UserDao.GetAndIncreaseUserNonce")
		}

		gateway, err = entities.AccountDao.Get(txnCtx, gateway.ID)
		if err != nil {
			return err
		}
		// validate and hold account balance
		if gateway.Hold.LessThan(totalAmount) {
			return rest.ErrFromGormError(txnCtx, err, "hold balance not enough")
		}

		sizeInWei = totalAmount.Mul(decimal.New(1, int32(constants.SightTokenDecimal)))

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
