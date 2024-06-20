package mysql

import (
	"context"
	"errors"
	"fmt"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/log"
	"github.com/jinzhu/gorm"
	"time"
)

func InitializeDao(o *Config) (*Dao, error) {
	masterDB, err := gorm.Open("mysql", o.Master.Dsn)
	if err != nil {
		errStr := fmt.Sprintf("failed to open MySQL master db, error=%v", err)
		return &Dao{}, errors.New(errStr)
	}
	master = masterDB
	master.DB().SetMaxIdleConns(o.Master.MaxIdle)
	master.DB().SetMaxOpenConns(o.Master.MaxOpen)
	master.SetNowFuncOverride(func() time.Time {
		return time.Now().UTC()
	})
	master.SetLogger(MySQLLogger{})
	master.SingularTable(true)
	if o.Debug {
		master = master.Debug()
	}
	// master.SetLogger(gorm.Logger{})
	log.Info().Msgf("connected to MySQL 'master' successfully, %v", o.Master.Dsn)

	return &Dao{db: master}, nil
}

type Dao struct {
	db *gorm.DB
}

type Txn struct {
	db *gorm.DB
}

func NewTxn(d *Dao) *Txn {
	return &Txn{db: d.db}
}

func (d *Dao) Db(ctx context.Context) *gorm.DB {
	txn, ok := ctx.Value(txn).(*gorm.DB)
	if !ok {
		return d.db
	}
	return txn
}

func (d *Txn) Tx(ctx context.Context, fc func(ctx context.Context) error) error {
	if _, ok := ctx.Value(txn).(*gorm.DB); ok {
		return fc(ctx)
	}
	return d.db.Transaction(func(tx *gorm.DB) error {
		newCtx := context.WithValue(ctx, txn, tx)
		return fc(newCtx)
	})
}
