package entities

import (
	"context"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/data/mysql"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/log"
	"github.com/jinzhu/gorm"
	"time"
)

const withdrawEventTableName = "withdraw_event"

// WithdrawEventDao is the exported data access object for queries, e.g., CRUD
var WithdrawEventDao withdrawEventDao

// withdrawEventDao is the interface lists all functions of the entity
type withdrawEventDao interface {
	Create(context.Context, *WithdrawEvent) error
	BatchCreate(context.Context, []*WithdrawEvent) error
	GetLatestTimestamp(ctx context.Context) (int64, error)
}

type withdrawEventDaoImpl struct {
	*mysql.Dao
}

// WithdrawEvent is a struct
type WithdrawEvent struct {
	gorm.Model
	TxnHash        string
	BlockTimestamp time.Time
	ToAddr         string
	Amount         string
	Nonce          uint
}

// TableName get/set the table name of the entity
func (e *WithdrawEvent) TableName() string {
	return withdrawEventTableName
}

func (i *withdrawEventDaoImpl) Create(ctx context.Context, e *WithdrawEvent) error {
	r := i.Db(ctx).Where("txn_hash = ?", e.TxnHash).
		Find(&WithdrawEvent{})
	if r.RowsAffected > 0 {
		return mysql.ErrDuplicatedKey
	}
	return i.Db(ctx).Create(e).Error
}

func (i *withdrawEventDaoImpl) BatchCreate(ctx context.Context, l []*WithdrawEvent) error {
	for _, e := range l {
		err := i.Create(ctx, e)
		if mysql.GetGormErrorCode(err) == mysql.GormErrorCodeDuplicateEntry || err == mysql.ErrDuplicatedKey {
			log.Error(ctx).Err(err).Msgf("failed create withdrawEvent %+v", e)
			continue
		} else if err != nil {
			return err
		}
	}

	return nil
}

func (i *withdrawEventDaoImpl) GetLatestTimestamp(ctx context.Context) (int64, error) {
	res := &WithdrawEvent{}
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
