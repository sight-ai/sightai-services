package entities

import (
	"context"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/data/mysql"
	"github.com/jinzhu/gorm"
)

const gatewayTableName = "gateway"

// GatewayDao is the exported data access object for queries, e.g., CRUD
var GatewayDao gatewayDao

// gatewayDao is the interface lists all functions of the entity
type gatewayDao interface {
	Upsert(ctx context.Context, e *Gateway) error
	GetsAll(ctx context.Context) (gateways []*Gateway, err error)
	GetById(ctx context.Context, id uint) (*Gateway, error)
	Delete(ctx context.Context, id uint) error
}

type gatewayDaoImpl struct {
	*mysql.Dao
}

// Gateway is a struct
type Gateway struct {
	gorm.Model
	Address  string
	Endpoint string
	Name     string
}

// TableName get/set the table name of the entity
func (e *Gateway) TableName() string {
	return gatewayTableName
}

func (i *gatewayDaoImpl) Upsert(ctx context.Context, e *Gateway) error {
	db := i.Db(ctx)
	if e.ID != 0 {
		p := &Gateway{}
		db = i.Db(ctx).Unscoped().
			Where("id = ?", e.ID).First(p)
		if db.RowsAffected != 0 {
			// update or bring back soft deleted
			db = i.Db(ctx).Unscoped().Save(e)
		} else {
			db = i.Db(ctx).Create(e)
		}
	} else {
		db = i.Db(ctx).Create(e)
	}

	return db.Error

}

func (i *gatewayDaoImpl) GetsAll(ctx context.Context) (gateways []*Gateway, err error) {
	err = i.Db(ctx).Unscoped().Find(&gateways).Error

	return gateways, err

}

func (i *gatewayDaoImpl) GetById(ctx context.Context, id uint) (*Gateway, error) {
	res := &Gateway{}
	err := i.Db(ctx).First(res, id).Error

	return res, err
}

func (i *gatewayDaoImpl) Delete(ctx context.Context, id uint) error {
	err := i.Db(ctx).Delete(&Gateway{}, id).Error
	return err
}
