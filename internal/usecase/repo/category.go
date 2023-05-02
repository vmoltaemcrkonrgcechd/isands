package repo

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"isands/internal/entity"
	"isands/pkg/postgres"
	"log"
	"net/http"
)

type CategoryRepo struct {
	*postgres.Postgres
}

func NewCategoryRepo(pg *postgres.Postgres) *CategoryRepo {
	return &CategoryRepo{pg}
}

func (r *CategoryRepo) Create(dto entity.DictionaryEntryDTO) (id string, err error) {
	if err = r.Sq.Insert("category").
		Columns("name").
		Values(dto.Name).
		Suffix("returning category_id").
		QueryRow().Scan(&id); err != nil {
		log.Println(err)
		return "", fiber.NewError(http.StatusInternalServerError,
			"произошла ошибка при добавлении категории")
	}

	return id, nil
}

func (r *CategoryRepo) Read() (dictionary entity.Dictionary, err error) {
	var rows *sql.Rows
	if rows, err = r.Sq.Select("category_id", "name").
		From("category").
		Query(); err != nil {
		log.Println(err)
		return nil, fiber.NewError(http.StatusInternalServerError,
			"произошла ошибка при получении всех категорий")
	}

	var dictionaryEntry entity.DictionaryEntry

	for rows.Next() {
		if err = rows.Scan(&dictionaryEntry.ID, &dictionaryEntry.Name); err != nil {
			log.Println(err)
			return nil, fiber.NewError(http.StatusInternalServerError,
				"произошла ошибка при получении всех категорий")
		}

		dictionary = append(dictionary, dictionaryEntry)
	}

	return dictionary, nil
}

func (r *CategoryRepo) Update(dto entity.DictionaryEntryDTO, id string) (string, error) {
	if _, err := r.Sq.Update("category").
		Set("name", dto.Name).
		Where("category_id = ?", id).
		Exec(); err != nil {
		log.Println(err)
		return "", fiber.NewError(http.StatusInternalServerError,
			"произошла ошибка при редактировании категории")
	}

	return id, nil
}

func (r *CategoryRepo) Delete(id string) error {
	if _, err := r.Sq.Delete("category").
		Where("category_id = ?", id).
		Exec(); err != nil {
		log.Println(err)
		return fiber.NewError(http.StatusInternalServerError,
			"произошла ошибка при удалении категории")
	}

	return nil
}
