package main

import (
	"fmt"
	"log"

	"github.com/infinity/identity-service/internal/config"
)

func main() {
	viperConfig := config.NewViper()
	app := config.NewFiber(viperConfig)

	config.Bootstrap(&config.BootstrapConfig{
		App:    app,
		Config: viperConfig,
	})

	webPort := viperConfig.GetInt("web.port")
	err := app.Listen(fmt.Sprintf(":%d", webPort))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
