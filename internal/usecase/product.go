package usecase

import (
	"isands/internal/entity"
	"isands/internal/usecase/repo"
)

type ProductUseCase struct {
	repo *repo.ProductRepo
}

func NewProductUseCase(r *repo.ProductRepo) *ProductUseCase {
	return &ProductUseCase{repo: r}
}

func (uc *ProductUseCase) Create(dto entity.ProductDTO) (string, error) {
	id, err := uc.repo.CreateProduct(dto)
	if err != nil {
		return "", err
	}

	if len(dto.Details) != 0 {
		if err = uc.repo.SaveDetails(dto.Details, id); err != nil {
			return "", err
		}
	}

	return id, nil
}

func (uc *ProductUseCase) Read(qp entity.ProductQP) ([]*entity.Product, error) {
	products, err := uc.repo.Read(qp)
	if err != nil {
		return nil, err
	}

	if products == nil {
		products = make([]*entity.Product, 0, 0)
	}

	return products, nil
}
