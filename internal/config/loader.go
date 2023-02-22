package config

import "github.com/spf13/viper"

// Load try to load application configuration.
func Load() (Config, error) {
	cfg := Default()

	viper.AddConfigPath(".")
	viper.AddConfigPath("~/.config/sd-gallery")
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return cfg, err
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
