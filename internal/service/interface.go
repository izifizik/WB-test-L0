package service

import (
	"WB-test-L0/internal/domain/model"
)

type Service interface {
	FindByUUID(uuid string) (model.Entity, error)

	SetEntity(message []byte) error
	GetAllEntity() error
}
