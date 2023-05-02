package usecase

import (
	"github.com/gofiber/fiber/v2"
	"isands/internal/entity"
	"isands/internal/usecase/repo"
	"net/http"
)

type CategoryUseCase struct {
	repo *repo.CategoryRepo
}

func NewCategoryUseCase(r *repo.CategoryRepo) *CategoryUseCase {
	return &CategoryUseCase{repo: r}
}

func (r *CategoryUseCase) Create(dto entity.DictionaryEntryDTO) (string, error) {
	if len([]rune(dto.Name)) < 1 {
		return "", fiber.NewError(http.StatusBadRequest,
			"название категории должно состоять более чем из одного символа")
	}

	return r.repo.Create(dto)
}

func (r *CategoryUseCase) Read() (entity.Dictionary, error) {
	dictionary, err := r.repo.Read()
	if err != nil {
		return nil, err
	}

	if dictionary == nil {
		dictionary = make([]entity.DictionaryEntry, 0, 0)
	}

	return dictionary, nil
}

func (r *CategoryUseCase) Update(dto entity.DictionaryEntryDTO, id string) (string, error) {
	if len([]rune(dto.Name)) < 1 {
		return "", fiber.NewError(http.StatusBadRequest,
			"название категории должно состоять более чем из одного символа")
	}

	return r.repo.Update(dto, id)
}

func (r *CategoryUseCase) Delete(id string) error {
	return r.repo.Delete(id)
}
