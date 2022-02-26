package cache

import (
	"WB-test-L0/internal/domain/model"
)

type Cache interface {
	Set(uuid string, value model.Entity)
	Get(uuid string) model.Entity
}
