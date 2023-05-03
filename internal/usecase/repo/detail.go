package repo

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"isands/internal/entity"
	"isands/pkg/postgres"
	"log"
	"net/http"
)

type DetailRepo struct {
	*postgres.Postgres
}

func NewDetailRepo(pg *postgres.Postgres) *DetailRepo {
	return &DetailRepo{pg}
}

func (r *DetailRepo) Create(dto entity.DictionaryEntryDTO) (id string, err error) {
	if err = r.Sq.Insert("detail").
		Columns("name").
		Values(dto.Name).
		Suffix("returning detail_id").
		QueryRow().Scan(&id); err != nil {
		log.Println(err)
		return "", fiber.NewError(http.StatusInternalServerError,
			"при добавлении детали произошла ошибка")
	}

	return id, nil
}

func (r *DetailRepo) Read() (dictionary entity.Dictionary, err error) {
	var rows *sql.Rows
	if rows, err = r.Sq.Select("detail_id", "name").
		From("detail").
		Query(); err != nil {
		log.Println(err)
		return nil, fiber.NewError(http.StatusInternalServerError,
			"произошла ошибка при получении всех деталей")
	}

	var dictionaryEntry entity.DictionaryEntry

	for rows.Next() {
		if err = rows.Scan(&dictionaryEntry.ID, &dictionaryEntry.Name); err != nil {
			log.Println(err)
			return nil, fiber.NewError(http.StatusInternalServerError,
				"произошла ошибка при получении всех деталей")
		}

		dictionary = append(dictionary, dictionaryEntry)
	}

	return dictionary, nil
}
