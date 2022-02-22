package cache

import (
	"WB-test-L0/internal/domain/model"
	"WB-test-L0/internal/service"
)

type cache struct {
	cache map[string]model.JsonInput
}

func NewCache() service.Cache {
	c := make(map[string]model.JsonInput)
	return &cache{cache: c}
}

func (c *cache) Set(uuid string, value model.JsonInput) bool {
	return true
}

func (c *cache) Get(uuid string) (model.JsonInput, bool) {
	return model.JsonInput{}, true
}
