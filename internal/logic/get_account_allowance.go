package logic

import (
	"context"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities/models"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/daodto"
)

func GetAccountAllowances(ctx context.Context, accountId uint) (*models.AllowancesResponse, error) {
	resp := &models.AllowancesResponse{}

	toAllowances, err := entities.AllowanceDao.GetsByToAccountID(ctx, accountId)
	if err != nil {
		return nil, err
	}

	for _, toAllowance := range toAllowances {
		dtoAllowance := daodto.AllowanceDao2Dto(toAllowance)
		resp.ToAllowances = append(resp.FromAllowances, *dtoAllowance)
	}

	fromAllowances, err := entities.AllowanceDao.GetsByFromAccountID(ctx, accountId)
	if err != nil {
		return nil, err
	}

	for _, fromAllowance := range fromAllowances {
		dtoAllowance := daodto.AllowanceDao2Dto(fromAllowance)
		resp.FromAllowances = append(resp.FromAllowances, *dtoAllowance)
	}

	return resp, nil
}
