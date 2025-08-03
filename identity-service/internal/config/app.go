package config

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/infinity/identity-service/internal/handlers"
	"github.com/infinity/identity-service/internal/logic"
	"github.com/infinity/identity-service/internal/repository"
	"github.com/infinity/identity-service/internal/routes"
	"github.com/infinity/identity-service/server/config"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB        *gorm.DB
	App       *fiber.App
	Logger    *logrus.Logger
	Validator *validator.Validate
	Config    *config.AppConfig
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	userRepository := repository.NewUserRepository(config.Logger)

	// setup logic
	userLogic := logic.NewUserLogic(config.Logger, config.DB, userRepository)

	// setup handler
	userHandler := handlers.NewUserHandler(config.Logger, config.Validator, userLogic)

	routeConfig := routes.RouteConfig{
		App:         config.App,
		UserHandler: userHandler,
	}
	routeConfig.Setup()
}
