package itest

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities"
	openapi "github.com/capybaralabs-xyz/sightai-services/internal/itest/client"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/config"
	sight_mysql "github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/data/mysql"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/log"
	"github.com/capybaralabs-xyz/sightai-services/internal/services"
	"github.com/capybaralabs-xyz/sightai-services/tools/insert_static_data/data"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	mysql_migrate "github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"strings"
	"time"
)

var TestContext *Context

type Context struct {
	webServer *echo.Echo
	client    *openapi.APIClient
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func NewContext() *Context {
	openapiConfig := openapi.NewConfiguration()
	openapiConfig.Host = "localhost:10101"
	openapiConfig.Scheme = "http"
	openapiConfig.Servers[0].URL = "http://localhost:10101/v1"

	client := openapi.NewAPIClient(openapiConfig)
	TestContext = &Context{
		client: client,
	}
	return TestContext
}

func (c *Context) Setup() {
	flag.Parse()

	testVersion := time.Now().UnixNano()
	config.Initialize()
	config.Cfg.AppName = fmt.Sprintf("Test%s%d", config.Cfg.AppName, testVersion)

	// init log
	log.Initialize(config.Cfg.AppName, false, zerolog.DebugLevel)

	// setup text database
	c.SetupDB(config.Cfg.AppName)

	// initialize mysql
	dao, err := sight_mysql.InitializeDao(config.Cfg.Mysql)
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to connect to MySQL")
	}

	// initialize dao
	entities.Initialize(dao)

	// insert static data
	c.InsertStaticData()

	var errCh chan error
	c.webServer, errCh, err = services.StartRest()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start service")
	}
	go func() {
		err = <-errCh
		log.Info().Err(err).Msg("service stopped")
	}()
}

func (c *Context) Teardown() {
	c.DropDB(config.Cfg.AppName)
	err := services.StopRest(c.webServer)
	if err != nil {
		log.Fatal(err)
	}
}

func (c *Context) SetupDB(name string) {
	dbInstanceName := strings.Split(config.Cfg.Mysql.Master.Dsn, "/")[0] + "/"
	db, err := sql.Open("mysql", dbInstanceName)
	check(err)
	defer db.Close()

	// create new schema and use it
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", name))
	check(err)

	// use new database
	migrateDsn := dbInstanceName + name + "?multiStatements=true"
	newDb, err := sql.Open("mysql", migrateDsn)
	check(err)
	defer newDb.Close()

	// migrate schema
	driver, err := mysql_migrate.WithInstance(newDb, &mysql_migrate.Config{})
	fsrc, err := (&file.File{}).Open("file://../migrations")
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

func (c *Context) DropDB(name string) {
	dbInstanceName := strings.Split(config.Cfg.Mysql.Master.Dsn, "/")[0] + "/"
	db, err := sql.Open("mysql", dbInstanceName)
	check(err)
	defer db.Close()

	// create new schema and use it
	_, err = db.Exec(fmt.Sprintf("DROP DATABASE %s", name))
	check(err)
}

func (c *Context) InsertStaticData() {
	err := data.InsertStaticData(context.Background())
	check(err)
}
