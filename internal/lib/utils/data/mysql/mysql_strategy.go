package mysql

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"
)

// Strategy defines the strategy to access MySQL
type Strategy struct {
	master *gorm.DB
	slave  *gorm.DB
}

const txn = "ongoing-txn"

func NewTxnContext(tx *Strategy) context.Context {
	return context.WithValue(context.Background(), txn, tx)
}

func TxnFromContext(ctx context.Context) *Strategy {
	s, ok := ctx.Value(txn).(*Strategy)
	if !ok {
		return nil
	}
	return s
}

type TxnBlock func(tx *Strategy) error

func DoInTxn(txnBlock TxnBlock) (err error) {
	mtx := master.Begin()
	if mtx.Error != nil {
		return mtx.Error
	}
	s := &Strategy{master: mtx}

	defer func() {
		if r := recover(); r != nil {
			// TODO: log stack trace when panic occurs like Echo does
			rerr, ok := r.(error)
			if !ok {
				rerr = fmt.Errorf("recovered from panic: %v", r)
			}
			err = rerr
		}
		// Commit or Rollback.
		if err == nil {
			// If no errors and we weren't panicking, then we can commit.
			err = s.master.Commit().Error
		} else {
			if rerr := s.master.Rollback(); rerr != nil {
				err = fmt.Errorf("failed to rollback from err: %v", err)
			}
		}
	}()

	err = txnBlock(s)
	return
}

// Register registers the Entity object to this dao
func (o *Strategy) Register(mysqlConfig *Config) {
	if mysqlConfig == nil {
		panic("DAO setup needs a non-nil mysqlConfig")
	}
	o.master = master
	o.slave = slave
}

func (o *Strategy) GormMaster() *gorm.DB {
	return master
}

func (o *Strategy) GormSlave() *gorm.DB {
	return slave
}
