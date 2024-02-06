package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type databaseCfg struct {
	Url string
}

type serverCfg struct {
	Port int32
}

type Cfg struct {
	Database databaseCfg
	Server serverCfg
}

var Config Cfg

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic( fmt.Errorf("erro fatal no arquivo de configuração: %w", err) )
	}

	Config.Database.Url = viper.GetString("database.url")
	Config.Server.Port = viper.GetInt32("server.port")
}