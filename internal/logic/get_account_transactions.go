package logic

import (
	"context"
	"github.com/capybaralabs-xyz/sightai-services/internal/constants/types"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities/models"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/daodto"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/log"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/rest"
	"time"
)

func GetAccountTransactions(ctx context.Context, accountID, page, pageSize uint, tt *types.TransactionType, before, after *time.Time) (*models.TransactionsResponse, error) {
	resp := &models.TransactionsResponse{}
	params := &entities.GetsTransactionsParams{
		Limit:           pageSize,
		Offset:          page * pageSize,
		AccountId:       &accountID,
		TransactionType: tt,
		Before:          before,
		After:           after,
	}

	transactions, err := entities.TransactionDao.GetsByFilters(ctx, params)
	if err != nil {
		return nil, rest.ErrFromGormError(ctx, err, "TransactionDao.GetsByFilters error")
	}

	for _, transaction := range transactions {
		dtoTransaction := daodto.TransactionDao2Dto(transaction)
		if err != nil {
			log.Error().Err(err).Msgf("failed parse transaction to dto %d", transaction.ID)
		} else if dtoTransaction != nil {
			resp.Transactions = append(resp.Transactions, *dtoTransaction)
		}
	}

	return resp, nil
}
