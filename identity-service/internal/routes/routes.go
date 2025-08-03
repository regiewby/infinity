package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/infinity/identity-service/internal/handlers"
)

type RouteConfig struct {
	App         *fiber.App
	UserHandler *handlers.UserHandler
}

func (c *RouteConfig) Setup() {
	c.App.Post("/identity/v1/user/create", c.UserHandler.Create)
}
