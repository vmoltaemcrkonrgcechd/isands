package routes

import (
	"github.com/gofiber/fiber/v2"
	"isands/internal/entity"
	"isands/internal/usecase"
	"net/http"
)

type detailRoutes struct {
	uc *usecase.DetailUseCase
}

func withDetailRoutes(app *fiber.App, uc *usecase.DetailUseCase) {
	r := detailRoutes{uc: uc}
	app.Post("/detail", r.create)
	app.Get("/detail", r.read)
}

//	@tags	деталь
//	@param	dto	body	entity.DictionaryEntryDTO	true	"деталь"
//	@router	/detail [post]
func (r detailRoutes) create(ctx *fiber.Ctx) error {
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

//	@tags		деталь
//	@success	200	{object}	entity.Dictionary
//	@router		/detail [get]
func (r detailRoutes) read(ctx *fiber.Ctx) error {
	dictionary, err := r.uc.Read()
	if err != nil {
		return err
	}

	return ctx.JSON(dictionary)
}
