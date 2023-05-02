package postgres

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
	"isands/config"
)

type Postgres struct {
	DB *sql.DB
	Sq sq.StatementBuilderType
}

func New(cfg *config.Config) (*Postgres, error) {
	db, err := sql.Open("postgres", cfg.PostgresURL)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &Postgres{
		DB: db,
		Sq: sq.StatementBuilder.RunWith(db).PlaceholderFormat(
			sq.Dollar),
	}, nil
}
