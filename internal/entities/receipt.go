package entities

import (
	"context"
	"github.com/capybaralabs-xyz/sightai-services/internal/constants/types"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/data/mysql"
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
	"gopkg.in/guregu/null.v4"
	"time"
)

const receiptTableName = "receipt"

// ReceiptDao is the exported data access object for queries, e.g., CRUD
var ReceiptDao receiptDao

// receiptDao is the interface lists all functions of the entity
type receiptDao interface {
	Create(ctx context.Context, receipt *Receipt) error
	BatchCreate(ctx context.Context, l []*Receipt) error
	Update(ctx context.Context, e *Receipt) error
	GetsByFilters(ctx context.Context, filters *GetsReceiptsParams) ([]*Receipt, error)
	Gets(ctx context.Context, ids []int64) ([]*Receipt, error)
}

type receiptDaoImpl struct {
	*mysql.Dao
}

// Receipt is a struct
type Receipt struct {
	gorm.Model
	UserId     uint
	GatewayId  uint
	FinishedAt null.Time
	Cost       decimal.Decimal `sql:"type:decimal(36,18);"`
	Proof      string
	TxnId      string
	Status     string
}

// TableName get/set the table name of the entity
func (e *Receipt) TableName() string {
	return receiptTableName
}

func (i *receiptDaoImpl) Create(ctx context.Context, receipt *Receipt) error {
	return i.Db(ctx).Create(receipt).Error
}

func (i *receiptDaoImpl) BatchCreate(ctx context.Context, l []*Receipt) error {
	for _, receipt := range l {
		if err := i.Db(ctx).Create(receipt).Error; err != nil {
			return err
		}
	}

	return nil
}

func (i *receiptDaoImpl) Update(ctx context.Context, e *Receipt) error {
	return i.Db(ctx).Save(e).Error
}

type GetsReceiptsParams struct {
	Limit         uint // 0 means no limit
	Offset        uint // 0 means no limit
	ReceiptStatus *types.ReceiptStatus
	GatewayId     *uint
	UserId        *uint
	Before        *time.Time
	After         *time.Time
	OrderBy       *types.OrderBy
}

func (i *receiptDaoImpl) GetsByFilters(ctx context.Context, filters *GetsReceiptsParams) ([]*Receipt, error) {
	res := []*Receipt{}
	db := i.Db(ctx).Table(receiptTableName)

	if filters.GatewayId != nil {
		db = db.Where("gateway_id = ?", filters.GatewayId)
	}

	if filters.UserId != nil {
		db = db.Where("user_id = ?", filters.UserId)
	}

	if filters.ReceiptStatus != nil {
		db = db.Where("`status` = ?", filters.ReceiptStatus)
	}

	if filters.Before != nil {
		db = db.Where("created_at <= ?", filters.Before)
	}

	if filters.After != nil {
		db = db.Where("created_at >= ?", filters.After)
	}

	if filters.Offset != 0 {
		db.Offset(filters.Offset)
	}

	if filters.Limit != 0 {
		db.Limit(filters.Limit)
	}

	if filters.OrderBy != nil {
		switch *filters.OrderBy {
		case types.OrderByCreateTimeNew:
			db.Order("created_at desc")
		case types.OrderByCreateTimeOld:
			db.Order("created_at")
		}
	} else {
		db.Order("created_at desc")
	}

	err := db.Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (i *receiptDaoImpl) Gets(ctx context.Context, ids []int64) ([]*Receipt, error) {
	res := []*Receipt{}
	err := i.Db(ctx).Find(&res, ids).Error
	return res, err
}
