package config

import (
	"github.com/spf13/viper"
)

type (
	Config struct {
		App      App      `yaml:"app"`
		Database Database `yaml:"database"`
	}

	App struct {
		Port int `yaml:"port"`
	}
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	}
)

const defaultConfigPath = "internal/config"

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(defaultConfigPath)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	return &config
}
