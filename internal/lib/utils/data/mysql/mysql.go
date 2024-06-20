package mysql

import (
	"errors"
	"fmt"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/log"
	"time"

	"github.com/jinzhu/gorm"
	// import mysql dialect
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	master *gorm.DB
	slave  *gorm.DB

	// ErrNoMasterDB when no master db not found
	ErrNoMasterDB = errors.New("no master db")
	// ErrNoMasterOrSlaveDB when no master or slave db not found
	ErrNoMasterOrSlaveDB = errors.New("no master or slave db")
	// ErrNoRecordUpdated when no record is updated
	ErrNoRecordUpdated = errors.New("no record updated")
	// ErrNoRecordFound when no record is found
	ErrNoRecordFound = errors.New("record not found")
	// ErrDuplicatedKey  when unique key violation
	ErrDuplicatedKey = errors.New("duplicated key")
	// ErrBeginTx when no record is found
	ErrBeginTx = errors.New("failed to start transaction")
)

// Initialize initializes connections to MySQL servers
func Initialize(o *Config) error {
	if o == nil {
		return errors.New("failed to initialize mysql because of empty mysql config")
	}
	// connect to master db if any
	if o.Master.Dsn != "" {
		masterDB, err := gorm.Open("mysql", o.Master.Dsn)
		if err != nil {
			errStr := fmt.Sprintf("failed to open MySQL master db, error=%v", err)
			log.Warn().Err(err).Msg(errStr)
			return errors.New(errStr)
		}
		master = masterDB
		master.DB().SetMaxIdleConns(o.Master.MaxIdle)
		master.DB().SetMaxOpenConns(o.Master.MaxOpen)
		master.DB().SetConnMaxLifetime(time.Duration(o.Master.MaxLifeTime) * time.Second)
		master.SetLogger(MySQLLogger{})
		master.SingularTable(true)
		//master.SetLogger(gorm.Logger{})
		log.Info().Msgf("connected to MySQL 'master' successfully, %v", o.Master.Dsn)
	}

	// connect to slave db if any
	if o.Slave.Dsn != "" {
		slaveDB, err := gorm.Open("mysql", o.Slave.Dsn)
		if err != nil {
			errStr := "failed to open MySQL slave db, dsn=%v"
			log.Warn().Msg(errStr)
			return errors.New(errStr)
		}
		slave = slaveDB
		slave.DB().SetMaxIdleConns(o.Slave.MaxIdle)
		slave.DB().SetMaxOpenConns(o.Slave.MaxOpen)
		slave.SetLogger(MySQLLogger{})
		slave.SingularTable(true)
		log.Info().Msgf("connected to MySQL 'slave' successfully, %v", o.Slave.Dsn)
	}

	if master == nil && slave == nil {
		errStr := "failed to connect any MySQL db"
		log.Error().Msg(errStr)
		return errors.New(errStr)
	}
	return nil
}

// Close close the connections to MySQL dbs
func Close() {
	if master != nil {
		log.Info().Msg("disconnected from MySQL master successfully")
		master.Close()
	}
	if slave != nil {
		log.Info().Msg("disconnected from MySQL slave successfully")
		slave.Close()
	}
}
