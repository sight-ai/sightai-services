package main

import (
	"context"
	"flag"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/config"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/data/mysql"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/log"
	"github.com/capybaralabs-xyz/sightai-services/tools/insert_static_data/data"
	"github.com/rs/zerolog"
)

// go run main.go
func main() {
	flag.Set("config-path", "../../../build")
	flag.Set("config-name", "config")
	flag.Parse()

	config.Initialize()
	log.Initialize(config.Cfg.AppName, false, zerolog.DebugLevel)

	// initialize mysql
	dao, err := mysql.InitializeDao(config.Cfg.Mysql)
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to connect to MySQL")
	}

	// initialize dao
	entities.Initialize(dao)

	err = data.InsertStaticData(context.Background())
	if err != nil {
		panic(err)
	}
}
