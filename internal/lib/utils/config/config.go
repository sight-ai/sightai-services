package config

import (
	"flag"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/comm_utils"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/data/mysql"
	"github.com/capybaralabs-xyz/sightai-services/internal/lib/utils/lib_subgraph"
	"github.com/spf13/viper"
	"path"
	"strings"
)

var cfgFile = flag.String("config-path", "./build", "Configuration file path")
var cfgName = flag.String("config-name", "config", "Configuration file name")

// Cfg is the exported global constant for server configurations
var Cfg = &struct {
	AppName      string               `json:"appName"`
	Env          string               `json:"env"`
	HostAndPort  string               `json:"hostAndPort"`
	JwtSecret    string               `json:"jwtSecret"`
	JwtPrv       string               `json:"jwtPrv"`
	JwtPub       string               `json:"jwtPub"`
	PrivateKey   string               `json:"privateKey"`
	SightChainId string               `json:"sightChainId"`
	Mysql        *mysql.Config        `json:"mysql"`
	Subgraph     *lib_subgraph.Config `json:"subgraph"`
}{}

func Initialize() {
	// read env variables
	viper.SetEnvPrefix(Cfg.AppName)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.BindEnv()

	// redis vars
	viper.BindEnv("redis.default.host")

	// read config file
	viper.SetConfigName(*cfgName)

	/*
		Look for the config file in a list of paths in order of precedence.
	*/
	// Path given by flag.
	viper.AddConfigPath(*cfgFile)

	// For local development runs. Should be last in precedence.
	dir := path.Join(comm_utils.GetGoBaseDirectory(), "build")
	viper.AddConfigPath(dir)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// parse the config file
	if err := viper.Unmarshal(Cfg); err != nil {
		panic(err)
	}
}
