package usecase

import (
	"github.com/gofiber/fiber/v2"
	"isands/internal/entity"
	"isands/internal/usecase/repo"
	"net/http"
)

type SupplierUseCase struct {
	repo *repo.SupplierRepo
}

func NewSupplierUseCase(r *repo.SupplierRepo) *SupplierUseCase {
	return &SupplierUseCase{repo: r}
}

func (r *SupplierUseCase) Create(dto entity.DictionaryEntryDTO) (string, error) {
	if len([]rune(dto.Name)) < 1 {
		return "", fiber.NewError(http.StatusBadRequest,
			"имя поставщика должно состоять более чем из одного символа")
	}

	return r.repo.Create(dto)
}

func (r *SupplierUseCase) Read() (entity.Dictionary, error) {
	dictionary, err := r.repo.Read()
	if err != nil {
		return nil, err
	}

	if dictionary == nil {
		dictionary = make([]entity.DictionaryEntry, 0, 0)
	}

	return dictionary, nil
}

func (r *SupplierUseCase) Update(dto entity.DictionaryEntryDTO, id string) (string, error) {
	if len([]rune(dto.Name)) < 1 {
		return "", fiber.NewError(http.StatusBadRequest,
			"имя поставщика должно состоять более чем из одного символа")
	}

	return r.repo.Update(dto, id)
}

func (r *SupplierUseCase) Delete(id string) error {
	return r.repo.Delete(id)
}
