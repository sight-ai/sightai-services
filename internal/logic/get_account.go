package logic

import (
	"context"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities/models"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/daodto"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/rest"
)

func GetAccount(ctx context.Context, accountId uint) (*models.Account, error) {
	account, err := entities.AccountDao.Get(ctx, accountId)
	if err != nil {
		return nil, rest.ErrFromGormError(ctx, err, "failed to get account")
	}

	return daodto.AccountDao2Dto(account), nil
}
