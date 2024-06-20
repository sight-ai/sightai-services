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

const transactionTableName = "transaction"

// TransactionDao is the exported data access object for queries, e.g., CRUD
var TransactionDao transactionDao

// transactionDao is the interface lists all functions of the entity
type transactionDao interface {
	BatchCreate(ctx context.Context, l []*Transaction) error
	Update(ctx context.Context, e *Transaction) error
	GetsByFilters(ctx context.Context, filters *GetsTransactionsParams) ([]*Transaction, error)
}

type transactionDaoImpl struct {
	*mysql.Dao
}

// Transaction is a struct
type Transaction struct {
	gorm.Model
	AccountId uint
	Available decimal.Decimal `sql:"type:decimal(36,18);"`
	Hold      decimal.Decimal `sql:"type:decimal(36,18);"`
	Type      string
	Notes     null.String
}

// TableName get/set the table name of the entity
func (e *Transaction) TableName() string {
	return transactionTableName
}

func (i *transactionDaoImpl) BatchCreate(ctx context.Context, l []*Transaction) error {
	for _, transaction := range l {
		if err := i.Db(ctx).Create(transaction).Error; err != nil {
			return err
		}
	}

	return nil
}

func (i *transactionDaoImpl) Update(ctx context.Context, e *Transaction) error {
	return i.Db(ctx).Save(e).Error
}

type GetsTransactionsParams struct {
	Limit           uint // 0 means no limit
	Offset          uint // 0 means no limit
	TransactionType *types.TransactionType
	AccountId       *uint
	Before          *time.Time
	After           *time.Time
	OrderBy         *types.OrderBy
}

func (i *transactionDaoImpl) GetsByFilters(ctx context.Context, filters *GetsTransactionsParams) ([]*Transaction, error) {
	res := []*Transaction{}
	db := i.Db(ctx).Table(transactionTableName)

	if filters.AccountId != nil {
		db = db.Where("account_id = ?", filters.AccountId)
	}

	if filters.TransactionType != nil {
		db = db.Where("`type` = ?", filters.TransactionType)
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
