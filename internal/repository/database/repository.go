package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type repo struct {
	database *sql.DB
}

//NewRepository - create new Repository interface reference
func NewRepository(database *sql.DB) Repository {
	return &repo{database: database}
}

//CreateEntity - set to postgres entity
func (r *repo) CreateEntity(id string, data []byte) error {
	tx, err := r.database.Begin()

	insertStmt := "insert into entity(id, data) values($1, $2::jsonb)"
	_, err = tx.Exec(insertStmt, id, data)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

//FindAllEntities - get all entities from postgres
func (r *repo) FindAllEntities() (*sql.Rows, error) {
	rows, err := r.database.Query(`SELECT * FROM entity`)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
