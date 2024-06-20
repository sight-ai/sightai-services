package daodto

import (
	"context"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities/models"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/rest"
)

func GatewayDao2Dto(ctx context.Context, daoGateway *entities.Gateway) (*models.Gateway, error) {
	account, err := entities.AccountDao.GetByAddress(ctx, daoGateway.Address)
	if err != nil {
		return nil, err
	}

	dtoGateway := &models.Gateway{
		Id:        int64(daoGateway.ID),
		AccountId: int64(account.ID),
		Address:   daoGateway.Address,
		Endpoint:  daoGateway.Endpoint,
		Name:      daoGateway.Name,
	}

	if daoGateway.DeletedAt != nil {
		dtoGateway.DeletedAt = *daoGateway.DeletedAt
	}

	return dtoGateway, nil
}

func AccountDao2Dto(daoAccount *entities.Account) *models.Account {
	return &models.Account{
		Id:        int64(daoAccount.ID),
		CreatedAt: daoAccount.CreatedAt,
		Address:   daoAccount.Address,
		Hold:      daoAccount.Hold.String(),
		Available: daoAccount.Available.String(),
		Nonce:     int64(daoAccount.Nonce),
		Role:      daoAccount.Role,
	}
}

func TransactionDao2Dto(daoTransaction *entities.Transaction) *models.Transaction {
	return &models.Transaction{
		Id:             int64(daoTransaction.ID),
		CreatedAt:      daoTransaction.CreatedAt,
		AccountId:      int64(daoTransaction.AccountId),
		AvailableDelta: daoTransaction.Available.String(),
		HoldDelta:      daoTransaction.Hold.String(),
		Type:           daoTransaction.Type,
		Notes:          daoTransaction.Notes.String,
	}
}

func ReceiptDao2Dto(ctx context.Context, daoReceipt *entities.Receipt) (*models.Receipt, error) {
	gateway, err := entities.AccountDao.Get(ctx, daoReceipt.GatewayId)
	if err != nil {
		return nil, rest.ErrBadRequest("gateway account not exists")
	}
	user, err := entities.AccountDao.Get(ctx, daoReceipt.UserId)
	if err != nil {
		return nil, rest.ErrBadRequest("user account not exists")
	}

	return &models.Receipt{
		Id:             int64(daoReceipt.ID),
		UserAddress:    user.Address,
		GatewayAddress: gateway.Address,
		FinishedAt:     daoReceipt.FinishedAt.Time,
		Cost:           daoReceipt.Cost.String(),
		Proof:          daoReceipt.Proof,
		TxnId:          daoReceipt.TxnId,
		Status:         daoReceipt.Status,
	}, nil
}

func AllowanceDao2Dto(daoAllowance *entities.Allowance) *models.Allowance {
	return &models.Allowance{
		Id:          int64(daoAllowance.ID),
		FromAccount: int64(daoAllowance.FromAccountID),
		ToAccount:   int64(daoAllowance.ToAccountID),
		Allowance:   daoAllowance.Allowance.String(),
		Version:     int64(daoAllowance.Version),
	}
}
