package repo

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"isands/internal/entity"
	"isands/pkg/postgres"
	"log"
	"net/http"
)

type SupplierRepo struct {
	*postgres.Postgres
}

func NewSupplierRepo(pg *postgres.Postgres) *SupplierRepo {
	return &SupplierRepo{pg}
}

func (r *SupplierRepo) Create(dto entity.DictionaryEntryDTO) (id string, err error) {
	if err = r.Sq.Insert("supplier").
		Columns("name").
		Values(dto.Name).
		Suffix("returning supplier_id").
		QueryRow().Scan(&id); err != nil {
		log.Println(err)
		return "", fiber.NewError(http.StatusInternalServerError,
			"произошла ошибка при добавлении поставщика")
	}

	return id, nil
}

func (r *SupplierRepo) Read() (dictionary entity.Dictionary, err error) {
	var rows *sql.Rows
	if rows, err = r.Sq.Select("supplier_id", "name").
		From("supplier").
		Query(); err != nil {
		log.Println(err)
		return nil, fiber.NewError(http.StatusInternalServerError,
			"произошла ошибка при получении всех поставщиков")
	}

	var dictionaryEntry entity.DictionaryEntry

	for rows.Next() {
		if err = rows.Scan(&dictionaryEntry.ID, &dictionaryEntry.Name); err != nil {
			log.Println(err)
			return nil, fiber.NewError(http.StatusInternalServerError,
				"произошла ошибка при получении всех поставщиков")
		}

		dictionary = append(dictionary, dictionaryEntry)
	}

	return dictionary, nil
}

func (r *SupplierRepo) Update(dto entity.DictionaryEntryDTO, id string) (string, error) {
	if _, err := r.Sq.Update("supplier").
		Set("name", dto.Name).
		Where("supplier_id = ?", id).
		Exec(); err != nil {
		log.Println(err)
		return "", fiber.NewError(http.StatusInternalServerError,
			"произошла ошибка при редактировании поставщика")
	}

	return id, nil
}

func (r *SupplierRepo) Delete(id string) error {
	if _, err := r.Sq.Delete("supplier").
		Where("supplier_id = ?", id).
		Exec(); err != nil {
		log.Println(err)
		return fiber.NewError(http.StatusInternalServerError,
			"произошла ошибка при удалении поставщика")
	}

	return nil
}
