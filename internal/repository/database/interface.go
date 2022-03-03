package database

import (
	"database/sql"
)

type Repository interface {
	CreateEntity(id string, input []byte) error
	FindAllEntities() (*sql.Rows, error)
}
