package logic

import (
	"context"
	"fmt"
	"github.com/capybaralabs-xyz/sightai-services/internal/constants"
	"github.com/capybaralabs-xyz/sightai-services/internal/constants/types"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities/models"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/comm_utils"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/rest"
	"github.com/shopspring/decimal"
	"gopkg.in/guregu/null.v4"
)

// UploadReceipt
// 1. deduct allowance
// 2. decrease user hold balance
// 3. increase gateway hold balance
func UploadReceipt(ctx context.Context, req *models.CreateReceiptRequest) (*models.SimpleMessageResponse, error) {
	gatewayAddr, err := comm_utils.ToEthAddress(req.GatewayAddress)
	if err != nil {
		return nil, rest.ErrBadRequest("invalid gateway address")
	}
	gateway, err := entities.AccountDao.GetByAddress(ctx, gatewayAddr)
	if err != nil {
		return nil, rest.ErrBadRequest("gateway account not exists")
	}

	userAddr, err := comm_utils.ToEthAddress(req.UserAddress)
	if err != nil {
		return nil, rest.ErrBadRequest("invalid user address")
	}
	user, err := entities.AccountDao.GetByAddress(ctx, userAddr)
	if err != nil {
		return nil, rest.ErrBadRequest("user account not exists")
	}

	cost, err := decimal.NewFromString(req.Cost)
	if err != nil {
		return nil, rest.ErrBadRequest("invalid cost")
	}

	status, err := types.NewReceiptStatusFromString(req.Status)
	if err != nil {
		return nil, rest.ErrBadRequest("invalid status")
	}

	receipt := &entities.Receipt{
		UserId:     user.ID,
		GatewayId:  gateway.ID,
		FinishedAt: null.TimeFrom(req.FinishedAt),
		Cost:       cost,
		Proof:      req.Proof,
		TxnId:      req.TxnId,
		Status:     status.String(),
	}

	err = entities.Txn.Tx(ctx, func(txnCtx context.Context) error {
		err = entities.ReceiptDao.Create(ctx, receipt)

		allowance, err := entities.AllowanceDao.GetByFromToAccountID(txnCtx, user.ID, gateway.ID)
		if err != nil {
			return rest.ErrFromGormError(txnCtx, err, "failed AllowanceDao.GetByFromToAccountID")
		}
		if allowance.Allowance.LessThan(cost) {
			return rest.ErrBadRequest("allowance less than cost")
		}
		_, err = entities.AllowanceDao.UpsertAllowance(txnCtx, user.ID, gateway.ID, allowance.Version+1, allowance.Allowance.Sub(cost))
		if err != nil {
			return rest.ErrFromGormError(txnCtx, err, "AllowanceDao.UpsertAllowance failed")
		}

		err = entities.AccountDao.DecreaseHoldBalance(txnCtx, user, cost, types.TransactionTypePay, fmt.Sprintf("pay receipt %d - %s", receipt.ID, receipt.TxnId))
		if err != nil {
			return rest.ErrFromGormError(txnCtx, err, "failed AccountDao.DecreaseHoldBalance")
		}

		err = entities.AccountDao.IncreaseHoldBalance(txnCtx, gateway, cost, types.TransactionTypeReceive, fmt.Sprintf("receive receipt %d - %s", receipt.ID, receipt.TxnId))
		if err != nil {
			return rest.ErrFromGormError(txnCtx, err, "failed AccountDao.IncreaseHoldBalance")
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
