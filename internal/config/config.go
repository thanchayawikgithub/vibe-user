package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type (
	Config struct {
		App      App      `yaml:"app"`
		Database Database `yaml:"database"`
		Oauth    Oauth    `yaml:"oauth"`
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

	Oauth struct {
		Google googleOauth `yaml:"google"`
	}

	googleOauth struct {
		ClientID     string   `yaml:"clientID"`
		ClientSecret string   `yaml:"clientSecret"`
		RedirectURL  string   `yaml:"redirectURL"`
		Scopes       []string `yaml:"scopes"`
		Endpoint     struct {
			AuthURL     string `yaml:"authURL"`
			TokenURL    string `yaml:"tokenURL"`
			UserinfoURL string `yaml:"userinfoURL"`
		} `yaml:"endpoint"`
	}
)

const (
	configPath = "internal/config"
	configName = "config"
	configType = "yaml"
)

func LoadConfig() *Config {
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Loaded config: %+v\n", config)
	return &config
}
