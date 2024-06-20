package logic

import (
	"context"
	"fmt"
	"github.com/capybaralabs-xyz/sightai-services/internal/constants/types"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities/models"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/rest"
	"github.com/shopspring/decimal"
)

func Deposit(ctx context.Context, req *models.DepositRequest) (*models.SimpleMessageResponse, error) {
	amount, err := decimal.NewFromString(req.Amount)
	if err != nil {
		return nil, rest.ErrBadRequest(fmt.Sprintf("invalid deposit amount %s", req.Amount))
	}

	toAccount := &entities.Account{
		Address: req.Address,
	}

	err = entities.AccountDao.AddBalance(ctx, toAccount, amount, types.TransactionTypeDeposit, "admin deposit")
	if err != nil {
		return nil, rest.ErrFromGormError(ctx, err, "AccountDao.AddBalance failed")
	}

	return &models.SimpleMessageResponse{Message: "OK"}, nil
}
