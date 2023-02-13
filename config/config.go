package config

import (
	"fmt"
	"log"

	"github.com/samriddhadev/go-acme-flights/utils"
	"github.com/spf13/viper"
)

const ENVIRONMENT string = "ENVIRONMENT"
const CONFIG_FILE_NAME string = "app-config"

type AppConfig struct {
	AppName  string               `yaml:"app-name"`
	Port     int                  `yaml:"port"`
	BasePath string               `yaml:"basepath"`
	FlightDB FlightDatabaseConfig `yaml:"flightdb"`
}

type FlightDatabaseConfig struct {
	Network     string `yaml:"network"`
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	User        string `yaml:"user"`
	PasswordKey string `yaml:"password-key"`
	Database    string `yaml:"database"`
	Timeout     int    `yaml:"timeout"`
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
	viper.SetConfigName(fmt.Sprintf("%s-%s", CONFIG_FILE_NAME, utils.GetEnv(ENVIRONMENT, "local")))
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
