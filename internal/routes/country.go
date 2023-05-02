package routes

import (
	"github.com/gofiber/fiber/v2"
	"isands/internal/entity"
	"isands/internal/usecase"
	"net/http"
)

type countryRoutes struct {
	uc *usecase.CountryUseCase
}

func withCountryRoutes(app *fiber.App, uc *usecase.CountryUseCase) {
	r := countryRoutes{uc: uc}
	app.Post("/country", r.create)
	app.Get("/country", r.read)
	app.Patch("/country/:id", r.update)
	app.Delete("/country/:id", r.delete)
}

// @tags	страна
// @param	dto	body	entity.DictionaryEntryDTO	true	"страна"
// @router	/country [post]
func (r countryRoutes) create(ctx *fiber.Ctx) error {
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

// @tags		страна
// @success	200	{object}	entity.Dictionary
// @router		/country [get]
func (r countryRoutes) read(ctx *fiber.Ctx) error {
	dictionary, err := r.uc.Read()
	if err != nil {
		return err
	}

	return ctx.JSON(dictionary)
}

// @tags	страна
// @param	dto	body	entity.DictionaryEntryDTO	true	"страна"
// @param	id	path	string						true	"идентификатор"
// @router	/country/{id} [patch]
func (r countryRoutes) update(ctx *fiber.Ctx) error {
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

// @tags	страна
// @param	id	path	string	true	"идентификатор"
// @router	/country/{id} [delete]
func (r countryRoutes) delete(ctx *fiber.Ctx) error {
	return r.uc.Delete(ctx.Params("id"))
}
