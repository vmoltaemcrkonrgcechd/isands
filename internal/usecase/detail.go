package usecase

import (
	"github.com/gofiber/fiber/v2"
	"isands/internal/entity"
	"isands/internal/usecase/repo"
	"net/http"
)

type DetailUseCase struct {
	repo *repo.DetailRepo
}

func NewDetailUseCase(r *repo.DetailRepo) *DetailUseCase {
	return &DetailUseCase{repo: r}
}

func (r *DetailUseCase) Create(dto entity.DictionaryEntryDTO) (string, error) {
	if len([]rune(dto.Name)) < 1 {
		return "", fiber.NewError(http.StatusBadRequest,
			"название детали должно состоять из более чем одного символа")
	}

	return r.repo.Create(dto)
}

func (r *DetailUseCase) Read() (entity.Dictionary, error) {
	dictionary, err := r.repo.Read()
	if err != nil {
		return nil, err
	}

	if dictionary == nil {
		dictionary = make([]entity.DictionaryEntry, 0, 0)
	}

	return dictionary, nil
}
