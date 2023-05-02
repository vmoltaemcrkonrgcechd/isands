package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"isands/config"
	"isands/internal/routes"
	"isands/pkg/postgres"
)

func Run(cfg *config.Config) error {
	app := fiber.New()

	app.Use(cors.New(cors.Config{AllowOrigins: "http://localhost:8080"}))

	_, err := postgres.New(cfg)
	if err != nil {
		return err
	}

	routes.WithRouter(app)

	return app.Listen(cfg.HTTPAddr)
}
