package database

import (
	"WB-test-L0/internal/errors"
	"database/sql"
	_ "github.com/lib/pq"
)

type repo struct {
	database *sql.DB
}

func NewRepository(database *sql.DB) Repository {
	return &repo{database: database}
}

func (r *repo) CreateEntity(id string, data []byte) error {
	insertStmt := "insert into entity(id, data) values($1, $2::jsonb)"
	_, err := r.database.Exec(insertStmt, id, data)
	if err != nil {
		return errors.DbError
	}

	return nil
}

func (r *repo) FindAllEntity() (*sql.Rows, error) {
	rows, err := r.database.Query(`SELECT * FROM entity`)
	if err != nil {
		return nil, errors.DbError
	}

	return rows, nil
}
