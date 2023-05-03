package entity

import (
	sq "github.com/Masterminds/squirrel"
	"strings"
)

type ProductDTO struct {
	Name            string      `json:"name"`
	Price           float64     `json:"price"`
	Online          bool        `json:"online"`
	InstallmentPlan bool        `json:"installmentPlan"`
	InStock         bool        `json:"inStock"`
	CountryID       string      `json:"countryID"`
	SupplierID      string      `json:"supplierID"`
	CategoryID      string      `json:"categoryID"`
	ColorID         string      `json:"colorID"`
	Details         []DetailDTO `json:"details"`
}

type Product struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	Price           float64  `json:"price"`
	CountryName     *string  `json:"countryName"`
	SupplierName    *string  `json:"supplierName"`
	CategoryName    *string  `json:"categoryName"`
	ColorName       *string  `json:"colorName"`
	Online          bool     `json:"online"`
	InstallmentPlan bool     `json:"installmentPlan"`
	InStock         bool     `json:"inStock"`
	Details         []Detail `json:"details"`
}

type DetailDTO struct {
	DetailID string `json:"detailID"`
	Value    string `json:"value"`
}

type Detail struct {
	DetailName *string `json:"detailName"`
	Value      *string `json:"value"`
}

type ProductQP struct {
	Name    string   `json:"name"`
	OrderBy []string `query:"orderBy"`
}

func (qp ProductQP) Use(q sq.SelectBuilder) sq.SelectBuilder {
	if len(qp.OrderBy) != 0 {
		q = q.OrderBy("product." + strings.Join(qp.OrderBy, " "))
	}

	if len(qp.Name) != 0 {
		q = q.Where("lower(product.name) like lower('%' || ? || '%')", qp.Name)
	}

	return q
}
