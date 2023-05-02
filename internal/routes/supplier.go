package routes

import (
	"github.com/gofiber/fiber/v2"
	"isands/internal/entity"
	"isands/internal/usecase"
	"net/http"
)

type supplierRoutes struct {
	uc *usecase.SupplierUseCase
}

func withSupplierRoutes(app *fiber.App, uc *usecase.SupplierUseCase) {
	r := supplierRoutes{uc: uc}
	app.Post("/supplier", r.create)
	app.Get("/supplier", r.read)
	app.Patch("/supplier/:id", r.update)
	app.Delete("/supplier/:id", r.delete)
}

//	@tags	поставщик
//	@param	dto	body	entity.DictionaryEntryDTO	true	"поставщик"
//	@router	/supplier [post]
func (r supplierRoutes) create(ctx *fiber.Ctx) error {
	var dto entity.DictionaryEntryDTO

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

//	@tags		поставщик
//	@success	200	{object}	entity.Dictionary
//	@router		/supplier [get]
func (r supplierRoutes) read(ctx *fiber.Ctx) error {
	dictionary, err := r.uc.Read()
	if err != nil {
		return err
	}

	return ctx.JSON(dictionary)
}

//	@tags	поставщик
//	@param	dto	body	entity.DictionaryEntryDTO	true	"поставщик"
//	@param	id	path	string						true	"идентификатор"
//	@router	/supplier/{id} [patch]
func (r supplierRoutes) update(ctx *fiber.Ctx) error {
	var dto entity.DictionaryEntryDTO

	err := ctx.BodyParser(&dto)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest,
			"неверный json")
	}

	var id string

	if id, err = r.uc.Update(dto, ctx.Params("id")); err != nil {
		return err
	}

	return ctx.SendString(id)
}

//	@tags	поставщик
//	@param	id	path	string	true	"идентификатор"
//	@router	/supplier/{id} [delete]
func (r supplierRoutes) delete(ctx *fiber.Ctx) error {
	return r.uc.Delete(ctx.Params("id"))
}
