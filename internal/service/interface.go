package service

import (
	"WB-test-L0/internal/domain/model"
	"context"
)

type Service interface {
	FindByUUID(ctx context.Context, uuid string) (model.JsonInput, error)
	Publish(ctx context.Context, input model.JsonInput) error
}

type Repository interface {
	Create(ctx context.Context, input model.JsonInput) error
	FindByUUID(ctx context.Context, uuid string) (model.JsonInput, error)
}
