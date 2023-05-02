package usecase

import (
	"github.com/gofiber/fiber/v2"
	"isands/internal/entity"
	"isands/internal/usecase/repo"
	"net/http"
)

type ColorUseCase struct {
	repo *repo.ColorRepo
}

func NewColorUseCase(r *repo.ColorRepo) *ColorUseCase {
	return &ColorUseCase{repo: r}
}

func (r *ColorUseCase) Create(dto entity.DictionaryEntryDTO) (string, error) {
	if len([]rune(dto.Name)) < 1 {
		return "", fiber.NewError(http.StatusBadRequest,
			"название цвета должно состоять из более чем одного символа")
	}

	return r.repo.Create(dto)
}

func (r *ColorUseCase) Read() (entity.Dictionary, error) {
	dictionary, err := r.repo.Read()
	if err != nil {
		return nil, err
	}

	if dictionary == nil {
		dictionary = make([]entity.DictionaryEntry, 0, 0)
	}

	return dictionary, nil
}

func (r *ColorUseCase) Update(dto entity.DictionaryEntryDTO, id string) (string, error) {
	if len([]rune(dto.Name)) < 1 {
		return "", fiber.NewError(http.StatusBadRequest,
			"название цвета должно состоять из более чем одного символа")
	}

	return r.repo.Update(dto, id)
}

func (r *ColorUseCase) Delete(id string) error {
	return r.repo.Delete(id)
}
