package config

import (
	"github.com/infinity/identity-service/server/config"
	"github.com/sirupsen/logrus"
)

func NewLogger(appConfig *config.AppConfig) *logrus.Logger {
	log := logrus.New()
	log.SetLevel(logrus.Level(appConfig.Logger.LogLevel))
	switch appConfig.Logger.LogFormat {
	case "json":
		log.SetFormatter(&logrus.JSONFormatter{})
	case "text":
		log.SetFormatter(&logrus.TextFormatter{})
	default:
		log.SetFormatter(&logrus.TextFormatter{})
	}

	return log
}
