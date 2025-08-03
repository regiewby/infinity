package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/infinity/identity-service/server/config"
)

func NewFiber(appConfig *config.AppConfig) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      appConfig.ServiceName,
		ErrorHandler: NewErrorHandler(),
	})

	return app
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		return ctx.Status(code).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
}
