package usecase

import (
	"github.com/gofiber/fiber/v2"
	"isands/internal/entity"
	"isands/internal/usecase/repo"
	"net/http"
)

type CountryUseCase struct {
	repo *repo.CountryRepo
}

func NewCountryUseCase(r *repo.CountryRepo) *CountryUseCase {
	return &CountryUseCase{repo: r}
}

func (r *CountryUseCase) Create(dto entity.DictionaryEntryDTO) (string, error) {
	if len([]rune(dto.Name)) < 1 {
		return "", fiber.NewError(http.StatusBadRequest,
			"название страны должно состоять более чем из одного символа")
	}

	return r.repo.Create(dto)
}

func (r *CountryUseCase) Read() (entity.Dictionary, error) {
	dictionary, err := r.repo.Read()
	if err != nil {
		return nil, err
	}

	if dictionary == nil {
		dictionary = make([]entity.DictionaryEntry, 0, 0)
	}

	return dictionary, nil
}

func (r *CountryUseCase) Update(dto entity.DictionaryEntryDTO, id string) (string, error) {
	if len([]rune(dto.Name)) < 1 {
		return "", fiber.NewError(http.StatusBadRequest,
			"название страны должно состоять более чем из одного символа")
	}

	return r.repo.Update(dto, id)
}

func (r *CountryUseCase) Delete(id string) error {
	return r.repo.Delete(id)
}
