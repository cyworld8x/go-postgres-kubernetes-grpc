package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DbDriver          string `mapstructure:"DB_DRIVER"`
	DbSource          string `mapstructure:"DB_SOURCE"`
	HttpServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
	PushGatewayURL    string `mapstructure:"PUSH_GATEWAY_URL"`
}

func Load() (config Config, err error) {
	viper.AddConfigPath("./cmd/crawler/config")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
