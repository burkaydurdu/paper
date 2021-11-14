package config

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port   int
	APIURL string
}

type Config struct {
	AppName string
	IsDebug bool
	Server  ServerConfig
}

func New() (*Config, error) {
	config := &Config{}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("PAPER")

	const defaultPort = 8080
	viper.SetDefault("APPNAME", "paper")
	viper.SetDefault("ISDEBUG", true)
	viper.SetDefault("PORT", defaultPort)
	viper.SetDefault("APIURL", fmt.Sprintf("http://localhost:%d", defaultPort))

	config.AppName = viper.GetString("APPNAME")
	config.IsDebug = viper.GetBool("ISDEBUG")
	config.Server = ServerConfig{
		Port:   viper.GetInt("PORT"),
		APIURL: viper.GetString("APIURL"),
	}

	return config, nil
}

func (c Config) Print() {
	_, _ = pp.Println(c)
}
