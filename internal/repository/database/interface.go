package database

import (
	"database/sql"
)

type Repository interface {
	CreateEntity(id string, input []byte) error
	FindAllEntity() (*sql.Rows, error)
}
