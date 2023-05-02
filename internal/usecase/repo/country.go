package repo

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"isands/internal/entity"
	"isands/pkg/postgres"
	"log"
	"net/http"
)

type CountryRepo struct {
	*postgres.Postgres
}

func NewCountryRepo(pg *postgres.Postgres) *CountryRepo {
	return &CountryRepo{pg}
}

func (r *CountryRepo) Create(dto entity.DictionaryEntryDTO) (id string, err error) {
	if err = r.Sq.Insert("country").
		Columns("name").
		Values(dto.Name).
		Suffix("returning country_id").
		QueryRow().Scan(&id); err != nil {
		log.Println(err)
		return "", fiber.NewError(http.StatusInternalServerError,
			"при добавлении страны произошла ошибка")
	}

	return id, nil
}

func (r *CountryRepo) Read() (dictionary entity.Dictionary, err error) {
	var rows *sql.Rows
	if rows, err = r.Sq.Select("country_id", "name").
		From("country").
		Query(); err != nil {
		log.Println(err)
		return nil, fiber.NewError(http.StatusInternalServerError,
			"произошла ошибка при получении всех стран")
	}

	var dictionaryEntry entity.DictionaryEntry

	for rows.Next() {
		if err = rows.Scan(&dictionaryEntry.ID, &dictionaryEntry.Name); err != nil {
			log.Println(err)
			return nil, fiber.NewError(http.StatusInternalServerError,
				"произошла ошибка при получении всех стран")
		}

		dictionary = append(dictionary, dictionaryEntry)
	}

	return dictionary, nil
}

func (r *CountryRepo) Update(dto entity.DictionaryEntryDTO, id string) (string, error) {
	if _, err := r.Sq.Update("country").
		Set("name", dto.Name).
		Where("country_id = ?", id).
		Exec(); err != nil {
		log.Println(err)
		return "", fiber.NewError(http.StatusInternalServerError,
			"произошла ошибка при редактировании страны")
	}

	return id, nil
}

func (r *CountryRepo) Delete(id string) error {
	if _, err := r.Sq.Delete("country").
		Where("country_id = ?", id).
		Exec(); err != nil {
		log.Println(err)
		return fiber.NewError(http.StatusInternalServerError,
			"при удалении страны произошла ошибка")
	}

	return nil
}
