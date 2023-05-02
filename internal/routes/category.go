package routes

import (
	"github.com/gofiber/fiber/v2"
	"isands/internal/entity"
	"isands/internal/usecase"
	"net/http"
)

type categoryRoutes struct {
	uc *usecase.CategoryUseCase
}

func withCategoryRoutes(app *fiber.App, uc *usecase.CategoryUseCase) {
	r := categoryRoutes{uc: uc}
	app.Post("/category", r.create)
	app.Get("/category", r.read)
	app.Patch("/category/:id", r.update)
	app.Delete("/category/:id", r.delete)
}

// @tags	категория
// @param	dto	body	entity.DictionaryEntryDTO	true	"категория"
// @router	/category [post]
func (r categoryRoutes) create(ctx *fiber.Ctx) error {
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

// @tags		категория
// @success	200	{object}	entity.Dictionary
// @router		/category [get]
func (r categoryRoutes) read(ctx *fiber.Ctx) error {
	dictionary, err := r.uc.Read()
	if err != nil {
		return err
	}

	return ctx.JSON(dictionary)
}

// @tags	категория
// @param	dto	body	entity.DictionaryEntryDTO	true	"категория"
// @param	id	path	string						true	"идентификатор"
// @router	/category/{id} [patch]
func (r categoryRoutes) update(ctx *fiber.Ctx) error {
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

// @tags	категория
// @param	id	path	string	true	"идентификатор"
// @router	/category/{id} [delete]
func (r categoryRoutes) delete(ctx *fiber.Ctx) error {
	return r.uc.Delete(ctx.Params("id"))
}
