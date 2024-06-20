package entities

import (
	"context"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/data/mysql"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/log"
	"github.com/jinzhu/gorm"
	"time"
)

const depositEventTableName = "deposit_event"

// DepositEventDao is the exported data access object for queries, e.g., CRUD
var DepositEventDao depositEventDao

// depositEventDao is the interface lists all functions of the entity
type depositEventDao interface {
	Create(context.Context, *DepositEvent) error
	BatchCreate(context.Context, []*DepositEvent) error
	GetLatestTimestamp(context.Context) (int64, error)
}

type depositEventDaoImpl struct {
	*mysql.Dao
}

// DepositEvent is a struct
type DepositEvent struct {
	gorm.Model
	TxnHash        string
	BlockTimestamp time.Time
	FromAddr       string
	ToAddr         string
	Amount         string
}

// TableName get/set the table name of the entity
func (e *DepositEvent) TableName() string {
	return depositEventTableName
}

func (i *depositEventDaoImpl) Create(ctx context.Context, e *DepositEvent) error {
	r := i.Db(ctx).Where("txn_hash = ?", e.TxnHash).
		Find(&DepositEvent{})
	if r.RowsAffected > 0 {
		return mysql.ErrDuplicatedKey
	}
	return i.Db(ctx).Create(e).Error
}

func (i *depositEventDaoImpl) BatchCreate(ctx context.Context, l []*DepositEvent) error {
	for _, e := range l {
		err := i.Create(ctx, e)
		if mysql.GetGormErrorCode(err) == mysql.GormErrorCodeDuplicateEntry || err == mysql.ErrDuplicatedKey {
			log.Error(ctx).Err(err).Msgf("failed create depositEvent %+v", e)
			continue
		} else if err != nil {
			return err
		}
	}

	return nil
}

func (i *depositEventDaoImpl) GetLatestTimestamp(ctx context.Context) (int64, error) {
	res := &DepositEvent{}
	err := i.Db(ctx).
		Order("block_timestamp desc").
		Limit(1).
		Find(res).
		Error
	if gorm.IsRecordNotFoundError(err) {
		return 0, nil
	} else if err != nil {
		return 0, err
	}

	return res.BlockTimestamp.Unix(), nil
}
