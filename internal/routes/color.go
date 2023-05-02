package routes

import (
	"github.com/gofiber/fiber/v2"
	"isands/internal/entity"
	"isands/internal/usecase"
	"net/http"
)

type colorRoutes struct {
	uc *usecase.ColorUseCase
}

func withColorRoutes(app *fiber.App, uc *usecase.ColorUseCase) {
	r := colorRoutes{uc: uc}
	app.Post("/color", r.create)
	app.Get("/color", r.read)
	app.Patch("/color/:id", r.update)
	app.Delete("/color/:id", r.delete)
}

//	@tags	цвет
//	@param	dto	body	entity.DictionaryEntryDTO	true	"цвет"
//	@router	/color [post]
func (r colorRoutes) create(ctx *fiber.Ctx) error {
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

//	@tags		цвет
//	@success	200	{object}	entity.Dictionary
//	@router		/color [get]
func (r colorRoutes) read(ctx *fiber.Ctx) error {
	dictionary, err := r.uc.Read()
	if err != nil {
		return err
	}

	return ctx.JSON(dictionary)
}

//	@tags	цвет
//	@param	dto	body	entity.DictionaryEntryDTO	true	"цвет"
//	@param	id	path	string						true	"идентификатор"
//	@router	/color/{id} [patch]
func (r colorRoutes) update(ctx *fiber.Ctx) error {
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

//	@tags	цвет
//	@param	id	path	string	true	"идентификатор"
//	@router	/color/{id} [delete]
func (r colorRoutes) delete(ctx *fiber.Ctx) error {
	return r.uc.Delete(ctx.Params("id"))
}
