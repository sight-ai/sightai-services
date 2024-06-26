package main

import (
	"context"
	"flag"
	"github.com/capybaralabs-xyz/sightai-services/internal/entities"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/config"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/data/mysql"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/jwt_auth"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/log"
	"github.com/rs/zerolog"
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

	// initialize mysql
	dao, err := mysql.InitializeDao(config.Cfg.Mysql)
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to connect to MySQL")
	}

	// initialize dao
	entities.Initialize(dao)
	// 1-admin 2-gateway 3-user 4-gateway
	user, err := entities.AccountDao.Get(context.Background(), 3)
	check(err)

	jwtToken, err := jwt_auth.GenerateJwtFromAccount(user, "")
	check(err)

	log.Info().Msgf("jwt: %s", jwtToken)

	id, err := jwt_auth.GetAccountID(jwtToken)
	check(err)

	log.Info().Msgf("account id: %d", id)

}
