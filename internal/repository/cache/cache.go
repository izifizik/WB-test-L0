package cache

import (
	"WB-test-L0/internal/domain/model"
	"sync"
)

type cache struct {
	cache map[string]model.Entity
}

var cacheMap map[string]model.Entity
var m sync.Mutex

func NewCache() Cache {
	cacheMap = make(map[string]model.Entity)
	return &cache{cache: cacheMap}
}

func (c *cache) Set(uuid string, value model.Entity) {
	m.Lock()
	cacheMap[uuid] = value
	m.Unlock()
}

func (c *cache) Get(uuid string) model.Entity {
	return cacheMap[uuid]
}
