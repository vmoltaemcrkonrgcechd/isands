package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "isands/docs"
	"isands/internal/usecase"
)

func WithRouter(app *fiber.App, countryUseCase *usecase.CountryUseCase) {
	app.Get("/swagger-ui/*", swagger.New(swagger.ConfigDefault))

	withCountryRoutes(app, countryUseCase)
}
