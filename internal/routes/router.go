package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func WithRouter(app *fiber.App) {
	app.Get("/swagger-ui/*", swagger.New(swagger.ConfigDefault))
}
