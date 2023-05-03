package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "isands/docs"
	"isands/internal/usecase"
)

func WithRouter(app *fiber.App,
	countryUseCase *usecase.CountryUseCase,
	categoryUseCase *usecase.CategoryUseCase,
	colorUseCase *usecase.ColorUseCase,
	supplierUseCase *usecase.SupplierUseCase,
	productUseCase *usecase.ProductUseCase) {

	app.Get("/swagger-ui/*", swagger.New(swagger.ConfigDefault))

	withCountryRoutes(app, countryUseCase)
	withCategoryRoutes(app, categoryUseCase)
	withColorRoutes(app, colorUseCase)
	withSupplierRoutes(app, supplierUseCase)
	withProductRoutes(app, productUseCase)
}
