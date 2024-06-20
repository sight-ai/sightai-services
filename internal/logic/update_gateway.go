package logic

import (
	"context"
	"github.com/capybaralabs-xyz/sightai-services/internal/constants"
	"github.com/capybaralabs-xyz/sightai-services/internal/constants/types"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities/models"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/comm_utils"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/rest"
	"github.com/jinzhu/gorm"
)

func UpdateGateway(ctx context.Context, req *models.UpsertGatewayRequest) (*models.SimpleMessageResponse, error) {
	addr, err := comm_utils.ToEthAddress(req.Address)
	if err != nil {
		return nil, rest.ErrBadRequest("invalid gateway address")
	}

	// handle delete
	if req.Id != 0 && req.Deleted {
		err = entities.GatewayDao.Delete(ctx, uint(req.Id))
		if err != nil {
			return nil, rest.ErrFromGormError(ctx, err, "GatewayDao.Delete failed")
		}

		return &models.SimpleMessageResponse{
			Message: constants.OKResponse,
		}, nil
	}

	gateway := &entities.Gateway{
		Model:    gorm.Model{ID: uint(req.Id)},
		Address:  addr,
		Endpoint: req.Endpoint,
		Name:     req.Name,
	}

	// handle update/add
	err = entities.GatewayDao.Upsert(ctx, gateway)
	if err != nil {
		return nil, rest.ErrFromGormError(ctx, err, "GatewayDao.Upsert failed")
	}

	account, _ := entities.AccountDao.GetByAddress(ctx, addr)
	if account == nil {
		err = entities.AccountDao.Create(ctx, &entities.Account{
			Address: addr,
			Nonce:   0,
			Role:    types.AccountRoleGateway.String(),
		})
		if err != nil {
			return nil, rest.ErrFromGormError(ctx, err, "AccountDao.Create failed")
		}
	} else {
		account.Role = types.AccountRoleGateway.String()
		err = entities.AccountDao.Update(ctx, account)
		if err != nil {
			return nil, rest.ErrFromGormError(ctx, err, "AccountDao.Create failed")
		}
	}

	return &models.SimpleMessageResponse{
		Message: constants.OKResponse,
	}, nil
}
