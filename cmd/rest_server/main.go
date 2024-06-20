package main

import (
	"flag"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/config"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/data/mysql"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/log"
	"github.com/capybaralabs-xyz/sightai-services/internal/services"
	"github.com/rs/zerolog"
)

func init() {
	flag.Parse()
}

func main() {
	// load configurations
	config.Initialize()

	// init log
	if config.Cfg.Env == "prod" {
		log.InitializeJson(zerolog.InfoLevel, config.Cfg.AppName)
	} else {
		log.Initialize(config.Cfg.AppName, false, zerolog.DebugLevel)
	}
	log.Info().Str("env", config.Cfg.Env).Msg("log initiated")

	// initialize dao
	dao, err := mysql.InitializeDao(config.Cfg.Mysql)
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to connect to MySQL")
	}
	entities.Initialize(dao)

	// start the restful services
	_, errCh, err := services.StartRest()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start restful service")
	}
	err = <-errCh
	log.Fatal().Err(err).Msg("service stopped")
}
