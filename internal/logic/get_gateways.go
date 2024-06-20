package logic

import (
	"context"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities/models"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/daodto"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/log"
)

func GetGateways(ctx context.Context) (*models.GatewaysResponse, error) {
	resp := &models.GatewaysResponse{}

	gateways, err := entities.GatewayDao.GetsAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, gateway := range gateways {
		dtoGateway, err := daodto.GatewayDao2Dto(ctx, gateway)
		if err != nil {
			log.Warn("failed GatewayDao2Dto")
			continue
		}
		resp.Gateways = append(resp.Gateways, *dtoGateway)

	}

	return resp, nil
}
