package repo

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"isands/internal/entity"
	"isands/pkg/postgres"
	"log"
	"net/http"
)

type ColorRepo struct {
	*postgres.Postgres
}

func NewColorRepo(pg *postgres.Postgres) *ColorRepo {
	return &ColorRepo{pg}
}

func (r *ColorRepo) Create(dto entity.DictionaryEntryDTO) (id string, err error) {
	if err = r.Sq.Insert("color").
		Columns("name").
		Values(dto.Name).
		Suffix("returning color_id").
		QueryRow().Scan(&id); err != nil {
		log.Println(err)
		return "", fiber.NewError(http.StatusInternalServerError,
			"произошла ошибка при добавлении цвета")
	}

	return id, nil
}

func (r *ColorRepo) Read() (dictionary entity.Dictionary, err error) {
	var rows *sql.Rows
	if rows, err = r.Sq.Select("color_id", "name").
		From("color").
		Query(); err != nil {
		log.Println(err)
		return nil, fiber.NewError(http.StatusInternalServerError,
			"произошла ошибка при получении всех цветов")
	}

	var dictionaryEntry entity.DictionaryEntry

	for rows.Next() {
		if err = rows.Scan(&dictionaryEntry.ID, &dictionaryEntry.Name); err != nil {
			log.Println(err)
			return nil, fiber.NewError(http.StatusInternalServerError,
				"произошла ошибка при получении всех цветов")
		}

		dictionary = append(dictionary, dictionaryEntry)
	}

	return dictionary, nil
}

func (r *ColorRepo) Update(dto entity.DictionaryEntryDTO, id string) (string, error) {
	if _, err := r.Sq.Update("color").
		Set("name", dto.Name).
		Where("color_id = ?", id).
		Exec(); err != nil {
		log.Println(err)
		return "", fiber.NewError(http.StatusInternalServerError,
			"произошла ошибка при редактировании цвета")
	}

	return id, nil
}

func (r *ColorRepo) Delete(id string) error {
	if _, err := r.Sq.Delete("color").
		Where("color_id = ?", id).
		Exec(); err != nil {
		log.Println(err)
		return fiber.NewError(http.StatusInternalServerError,
			"произошла ошибка при удалении цвета")
	}

	return nil
}
