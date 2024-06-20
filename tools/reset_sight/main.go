package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/config"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/data/mysql"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/log"
	"github.com/capybaralabs-xyz/sightai-services/tools/insert_static_data/data"
	"github.com/golang-migrate/migrate/v4"
	mysql_migrate "github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rs/zerolog"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	flag.Set("config-path", "build")
	flag.Set("config-name", "config")
	flag.Parse()

	config.Initialize()
	log.Initialize(config.Cfg.AppName, false, zerolog.DebugLevel)
	log.Info().Msgf("Starting in %s", config.Cfg.Env)

	clearMysqlData()
}

func clearMysqlData() {
	dropDB("sightai-services")
	setupDB("sightai-services")
	insertStaticData()
}

func setupDB(name string) {
	dbInstanceName := strings.Split(config.Cfg.Mysql.Master.Dsn, "/")[0] + "/"
	db, err := sql.Open("mysql", dbInstanceName)
	check(err)
	defer db.Close()

	// create new schema and use it
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE `%s`", name))
	check(err)

	// use new database
	migrateDsn := dbInstanceName + name + "?multiStatements=true"
	newDb, err := sql.Open("mysql", migrateDsn)
	check(err)
	defer newDb.Close()

	// migrate schema
	driver, err := mysql_migrate.WithInstance(newDb, &mysql_migrate.Config{})
	fsrc, err := (&file.File{}).Open("file://internal/migrations")
	check(err)
	m, err := migrate.NewWithInstance(
		"file",
		fsrc,
		"mysql",
		driver)
	m.Up()

	// overwrite config
	newDBDsn := dbInstanceName + name + "?parseTime=true"
	config.Cfg.Mysql.Master.Dsn = newDBDsn
	config.Cfg.Mysql.Slave.Dsn = newDBDsn
}

func dropDB(name string) {
	dbInstanceName := strings.Split(config.Cfg.Mysql.Master.Dsn, "/")[0] + "/"
	db, err := sql.Open("mysql", dbInstanceName)
	check(err)
	defer db.Close()

	// create new schema and use it
	_, err = db.Exec(fmt.Sprintf("DROP DATABASE `%s`", name))
}

func insertStaticData() {
	// initialize dao
	dao, err := mysql.InitializeDao(config.Cfg.Mysql)
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to connect to MySQL")
	}
	entities.Initialize(dao)

	err = data.InsertStaticData(context.Background())
	check(err)
}
