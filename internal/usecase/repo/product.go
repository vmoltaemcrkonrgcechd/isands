package repo

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"isands/internal/entity"
	"isands/pkg/postgres"
	"log"
	"net/http"
)

type ProductRepo struct {
	*postgres.Postgres
}

func NewProductRepo(pg *postgres.Postgres) *ProductRepo {
	return &ProductRepo{pg}
}

func (r *ProductRepo) CreateProduct(dto entity.ProductDTO) (id string, err error) {
	if err = r.Sq.Insert("product").
		Columns("name", "price", "country_id", "supplier_id", "category_id",
			"color_id", "online", "installment_plan", "in_stock").
		Values(dto.Name, dto.Price, dto.CountryID, dto.SupplierID, dto.CategoryID,
			dto.ColorID, dto.Online, dto.InstallmentPlan, dto.InStock).
		Suffix("returning product_id").
		QueryRow().Scan(&id); err != nil {
		log.Println(err)
		return "", fiber.NewError(http.StatusInternalServerError, "")
	}

	return id, nil
}

func (r *ProductRepo) SaveDetails(details []entity.DetailDTO, id string) error {
	query := r.Sq.Insert("product_detail").
		Columns("product_id", "detail_id", "value")

	for _, detail := range details {
		query = query.Values(id, detail.DetailID, detail.Value)
	}

	if _, err := query.Exec(); err != nil {
		log.Println(err)
		return fiber.NewError(http.StatusInternalServerError,
			"")
	}

	return nil
}

func (r *ProductRepo) Read(qp entity.ProductQP) (products []*entity.Product, err error) {
	query := r.Sq.Select("product_id", "product.name",
		"price", "country.name", "supplier.name",
		"category.name", "color.name",
		"online", "installment_plan", "in_stock", "detail.name", "value").
		From("product").
		LeftJoin("country using (country_id)").
		LeftJoin("supplier using (supplier_id)").
		LeftJoin("category using (category_id)").
		LeftJoin("color using (color_id)").
		LeftJoin("product_detail using (product_id)").
		LeftJoin("detail using (detail_id)")

	query = qp.Use(query)

	var rows *sql.Rows
	if rows, err = query.Query(); err != nil {
		log.Println(err)
		return nil, fiber.NewError(http.StatusInternalServerError,
			"")
	}

	var (
		p          entity.Product
		productMap = make(map[string]*entity.Product)
		d          entity.Detail
	)

	for rows.Next() {
		if err = rows.Scan(&p.ID, &p.Name, &p.Price, &p.CountryName, &p.SupplierName,
			&p.CategoryName, &p.ColorName, &p.Online, &p.InstallmentPlan,
			&p.InStock, &d.DetailName, &d.Value); err != nil {
			log.Println(err)
			return nil, fiber.NewError(http.StatusInternalServerError,
				"")
		}

		if _, ok := productMap[p.ID]; !ok {
			copyP := p

			copyP.Details = make([]entity.Detail, 0)

			addr := &copyP

			productMap[p.ID] = addr

			products = append(products, addr)
		}

		if d.DetailName == nil || d.Value == nil {
			continue
		}

		productMap[p.ID].Details = append(productMap[p.ID].Details, d)
	}

	return products, nil
}
