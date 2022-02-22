package database

import (
	"WB-test-L0/internal/domain/model"
	"WB-test-L0/internal/service"
	"context"
)

type repo struct {
}

func NewRepository() service.Repository {
	return &repo{}
}

func (r *repo) Create(ctx context.Context, input model.JsonInput) error {
	return nil
}

func (r *repo) FindByUUID(ctx context.Context, uuid string) (model.JsonInput, error) {
	return model.JsonInput{}, nil
}

