package mysql

import (
	"errors"
	"fmt"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/log"
	"time"

	"github.com/jinzhu/gorm"
)

func InitializeRawGorm(o *Config) (*gorm.DB, error) {
	masterDB, err := gorm.Open("mysql", o.Master.Dsn)
	if err != nil {
		errStr := fmt.Sprintf("failed to open MySQL master db, error=%v", err)
		return nil, errors.New(errStr)
	}
	master = masterDB
	master.DB().SetMaxIdleConns(o.Master.MaxIdle)
	master.DB().SetMaxOpenConns(o.Master.MaxOpen)
	master.SetNowFuncOverride(func() time.Time {
		return time.Now().UTC()
	})
	master.SetLogger(MySQLLogger{})
	master.SingularTable(true)
	// master.SetLogger(gorm.Logger{})
	log.Info().Msgf("connected to MySQL 'master' successfully, %v", o.Master.Dsn)
	return master, nil
}
