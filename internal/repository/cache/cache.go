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

//NewCache - create new Cache interface reference
func NewCache() Cache {
	cacheMap = make(map[string]model.Entity)
	return &cache{cache: cacheMap}
}

//Set - set value to cache (map)
func (c *cache) Set(uuid string, value model.Entity) {
	m.Lock()
	cacheMap[uuid] = value
	m.Unlock()
}

//Get - get value from cache (map)
func (c *cache) Get(uuid string) model.Entity {
	return cacheMap[uuid]
}
