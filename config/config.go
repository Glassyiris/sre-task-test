package config

import (
	"github.com/spf13/viper"
	"task-test/logger"
)

type Config struct {
	Database struct {
		Dsn     string `yaml:"dsn"`
		SqlType string `yaml:"sqlType"`
	}
	Server struct {
		Port uint32 `yaml:"port"`
	}
	Jwt struct {
		Secret string `yaml:"secret"`
	}
	Redis struct {
		Addr string `yaml:"address"`
		Port uint32 `yaml:"port"`
	}
}

func ParseConfig(file string) *Config {
	var config Config
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(file)
	err := viper.ReadInConfig()

	if err != nil {
		logger.Error(err.Error())
	}

	viper.Unmarshal(&config)
	return &config
}
