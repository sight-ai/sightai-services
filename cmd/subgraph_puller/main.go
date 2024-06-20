package main

import (
	"context"
	"flag"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/config"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/data/mysql"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/lib_subgraph"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/log"
	"github.com/capybaralabs-xyz/sightai-services/internal/services"
	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog"
	"time"
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

	// initialize subgraph
	lib_subgraph.Initialize(config.Cfg.Subgraph)

	// initialize product engine
	s := gocron.NewScheduler(time.UTC)
	s.TagsUnique()

	job, _ := s.Every(10).Second().Do(func() {
		services.PullSubgraph(context.Background())
	})
	job.SingletonMode()

	s.StartBlocking()
}
