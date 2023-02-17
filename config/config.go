package config

import (
	"fmt"
	"log"

	"github.com/samriddhadev/go-acme-flights/util"
	"github.com/spf13/viper"
)

const ENVIRONMENT string = "ENVIRONMENT"
const CONFIG_FILE_NAME string = "app-config"

type AppConfig struct {
	AppName             string               `mapstructure:"app-name"`
	Port                int                  `mapstructure:"port"`
	BasePath            string               `mapstructure:"basepath"`
	APIValidationSchema map[string]string    `mapstructure:"api-validation-schemas"`
	FlightDB            FlightDatabaseConfig `mapstructure:"flightdb"`
}

type FlightDatabaseConfig struct {
	Network     string `mapstructure:"network"`
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	User        string `mapstructure:"user"`
	PasswordKey string `mapstructure:"password-key"`
	Database    string `mapstructure:"database"`
	Timeout     int    `mapstructure:"timeout"`
}

type Config struct {
	AppConfig
}

func NewConfig() (*Config, error) {
	config, err := load()
	if err != nil {
		return &Config{}, err
	}
	return &config, nil
}

func load() (Config, error) {
	viper.SetConfigName(fmt.Sprintf("%s-%s", CONFIG_FILE_NAME, util.GetEnv(ENVIRONMENT, "local")))
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicf("error - reading config file - %s", err)
		return Config{}, err
	}
	viper.SetDefault("port", 8080)
	var appConfig AppConfig
	err = viper.Unmarshal(&appConfig)
	config := Config{}
	config.AppConfig = appConfig
	log.Println("configuration : ", config)
	return config, err
}
