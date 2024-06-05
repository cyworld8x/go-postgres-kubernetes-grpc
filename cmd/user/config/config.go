package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DbDriver          string `mapstructure:"DB_DRIVER"`
	DbSource          string `mapstructure:"DB_SOURCE"`
	HttpServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress string `mapstructure:"GRPC_SERVER_ADDRESS"`
}

func Load() (config Config, err error) {
	viper.AddConfigPath(".")
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
