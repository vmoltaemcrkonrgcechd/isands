package routes

import (
	"github.com/gofiber/fiber/v2"
	"isands/internal/entity"
	"isands/internal/usecase"
	"net/http"
)

type productRoutes struct {
	uc *usecase.ProductUseCase
}

func withProductRoutes(app *fiber.App, uc *usecase.ProductUseCase) {
	r := productRoutes{uc: uc}
	app.Post("/product", r.create)
	app.Get("/product", r.read)
}

//	@tags	продукт
//	@param	dto	body	entity.ProductDTO	true	"продукт"
//	@router	/product [post]
func (r productRoutes) create(ctx *fiber.Ctx) error {
	var dto entity.ProductDTO

	err := ctx.BodyParser(&dto)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest,
			"неверный json")
	}

	var id string

	if id, err = r.uc.Create(dto); err != nil {
		return err
	}

	return ctx.Status(http.StatusCreated).SendString(id)
}

//	@tags		продукт
//	@success	200			{array}	entity.Product
//	@param		name		query	string	false	"название продукта"
//	@param		countryID	query	string	false	"идентификатор страны, разделенный запятыми"
//	@param		supplierID	query	string	false	"идентификатор поставщика, разделенный запятыми"
//	@param		categoryID	query	string	false	"идентификатор категории, разделенный запятыми"
//	@param		colorID		query	string	false	"идентификатор цвета, разделенный запятыми"
//	@param		orderBy		query	string	false	"сортировка: name,price: name,desc/name,asc"
//	@router		/product [get]
func (r productRoutes) read(ctx *fiber.Ctx) error {
	var qp entity.ProductQP

	err := ctx.QueryParser(&qp)
	if err != nil {
		return err
	}

	products, err := r.uc.Read(qp)
	if err != nil {
		return err
	}

	return ctx.JSON(products)
}
