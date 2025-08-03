package server

import (
	"fmt"
	"log"

	"github.com/infinity/identity-service/internal/config"
)

func Serve() {
	appConfig, err := config.LoadAppConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	logger := config.NewLogger(appConfig)
	app := config.NewFiber(appConfig)
	db := config.NewDatabase(appConfig, logger)
	validator := config.NewValidator()

	config.Bootstrap(&config.BootstrapConfig{
		DB:        db,
		App:       app,
		Config:    appConfig,
		Validator: validator,
		Logger:    logger,
	})

	err = app.Listen(fmt.Sprintf(":%d", appConfig.Port))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
