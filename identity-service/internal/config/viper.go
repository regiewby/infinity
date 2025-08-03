package config

import (
	"fmt"

	"github.com/infinity/identity-service/server/config"
	"github.com/spf13/viper"
)

func LoadAppConfig() (*config.AppConfig, error) {
	v := viper.New()

	v.SetConfigFile("../../config_files/service-conf.json")
	v.SetConfigType("json")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var appConfig config.AppConfig
	if err := v.Unmarshal(&appConfig); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &appConfig, nil
}
