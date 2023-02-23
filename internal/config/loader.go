package config

import (
	"errors"
	"log"
	"strings"

	"github.com/spf13/viper"
)

// Load try to load application configuration.
func Load() (Config, error) {
	cfg := Default()

	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))

	if err := viper.ReadInConfig(); err != nil {
		switch {
		case errors.As(err, &viper.ConfigFileNotFoundError{}):
			log.Println("[WARN] Configuration file not found")
		default:
			return cfg, err
		}
	}

	err := viper.Unmarshal(&cfg)
	return cfg, err
}

// MustLoad load app configuration or throw panic.
func MustLoad() Config {
	if cfg, err := Load(); err != nil {
		panic(err)
	} else {
		return cfg
	}
}
