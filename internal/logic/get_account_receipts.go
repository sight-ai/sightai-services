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

func GetAccountReceipts(ctx context.Context, gatewayID, userID *uint, page, pageSize uint, rs *types.ReceiptStatus, before, after *time.Time) (*models.ReceiptsResponse, error) {
	resp := &models.ReceiptsResponse{}
	params := &entities.GetsReceiptsParams{
		Limit:         pageSize,
		Offset:        page * pageSize,
		GatewayId:     gatewayID,
		UserId:        userID,
		ReceiptStatus: rs,
		Before:        before,
		After:         after,
	}

	receipts, err := entities.ReceiptDao.GetsByFilters(ctx, params)
	if err != nil {
		return nil, rest.ErrFromGormError(ctx, err, "TransactionDao.GetsByFilters error")
	}

	for _, receipt := range receipts {
		dtoReceipt, err := daodto.ReceiptDao2Dto(ctx, receipt)
		if err != nil {
			log.Error().Err(err).Msgf("failed parse receipt to dto %d", receipt.ID)
		} else if dtoReceipt != nil {
			resp.Receipts = append(resp.Receipts, *dtoReceipt)
		}
	}

	return resp, nil
}
