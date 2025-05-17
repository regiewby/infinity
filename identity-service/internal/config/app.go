package config

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/infinity/identity-service/internal/delivery/http/route"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
	Producer *kafka.Producer
}

func Bootstrap(config *BootstrapConfig) {
	routeConfig := route.RouteConfig{
		App: config.App,
	}
	routeConfig.Setup()
}
