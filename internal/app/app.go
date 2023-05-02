package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"isands/config"
	"isands/internal/routes"
	"isands/internal/usecase"
	"isands/internal/usecase/repo"
	"isands/pkg/postgres"
)

func Run(cfg *config.Config) error {
	app := fiber.New()

	app.Use(cors.New(cors.Config{AllowOrigins: "http://localhost:8080"}))

	pg, err := postgres.New(cfg)
	if err != nil {
		return err
	}

	routes.WithRouter(app,
		usecase.NewCountryUseCase(repo.NewCountryRepo(pg)),
		usecase.NewCategoryUseCase(repo.NewCategoryRepo(pg)),
		usecase.NewColorUseCase(repo.NewColorRepo(pg)),
	)

	return app.Listen(cfg.HTTPAddr)
}
