package config

import (
	"go-jwt/helper"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	PORT             string `mapstructure:"PORT"`
	MYSQL_CONNECTION string `mapstructure:"MYSQL_CONNECTION"`
}

var AppConfig *Config

func LoadAppConfig() {
	log.Println("Loading Server Configurations...")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()

	helper.PanicIfError(err)

	err = viper.Unmarshal(&AppConfig)
	helper.PanicIfError(err)
	log.Println("Configuration Success...")
}
