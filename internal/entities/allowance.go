package entities

import (
	"context"
	"errors"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/data/mysql"
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

const allowanceTableName = "allowance"

// AllowanceDao is the exported data access object for queries, e.g., CRUD
var AllowanceDao allowanceDao

// allowanceDao is the interface lists all functions of the entity
type allowanceDao interface {
	GetsByFromAccountID(ctx context.Context, id uint) ([]*Allowance, error)
	GetsByToAccountID(ctx context.Context, id uint) ([]*Allowance, error)
	GetByFromToAccountID(ctx context.Context, fid, tid uint) (*Allowance, error)
	UpsertAllowance(ctx context.Context, fid, tid, version uint, allowance decimal.Decimal) (*Allowance, error)
}

type allowanceDaoImpl struct {
	*mysql.Dao
}

// Allowance is a struct
type Allowance struct {
	gorm.Model
	FromAccountID uint
	ToAccountID   uint
	Allowance     decimal.Decimal
	Version       uint
}

// TableName get/set the table name of the entity
func (e *Allowance) TableName() string {
	return allowanceTableName
}

func (i *allowanceDaoImpl) GetsByFromAccountID(ctx context.Context, id uint) ([]*Allowance, error) {
	res := []*Allowance{}
	if err := i.Db(ctx).Where("from_account_id = ?", id).Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func (i *allowanceDaoImpl) GetsByToAccountID(ctx context.Context, id uint) ([]*Allowance, error) {
	res := []*Allowance{}
	if err := i.Db(ctx).Where("to_account_id = ?", id).Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func (i *allowanceDaoImpl) GetByFromToAccountID(ctx context.Context, fid, tid uint) (*Allowance, error) {
	res := &Allowance{}
	if err := i.Db(ctx).Where("from_account_id = ? and to_account_id = ?", fid, tid).First(res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func (i *allowanceDaoImpl) UpsertAllowance(ctx context.Context, fid, tid, version uint, allowance decimal.Decimal) (*Allowance, error) {
	res := &Allowance{}
	err := i.Db(ctx).Where("from_account_id = ? and to_account_id = ?", fid, tid).First(res).Error
	if err == gorm.ErrRecordNotFound {
		// create a new record
		res = &Allowance{
			FromAccountID: fid,
			ToAccountID:   tid,
			Allowance:     allowance,
		}
		err = i.Db(ctx).Create(res).Error
		return res, err
	} else {
		// update existing record
		if res.Version >= version {
			return res, errors.New("invalid version")
		}
		res.Allowance = allowance
		res.Version = version
		return res, i.Db(ctx).Save(res).Error
	}
}
