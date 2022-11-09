package mutextcache

import (
	"sync"

	"github.com/vladjong/async_cache/helpers"
)

type mutexCache struct {
	mx      sync.RWMutex
	storage map[string]string
}

func NewCache() *mutexCache {
	return &mutexCache{
		storage: make(map[string]string),
	}
}

func (c *mutexCache) Set(key, value string) error {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.storage[key] = value
	return nil
}

func (c *mutexCache) Get(key string) (string, error) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	value, ok := c.storage[key]
	if !ok {
		return "", helpers.ErrNotFound
	}
	return value, nil
}

func (c *mutexCache) Delete(key string) error {
	c.mx.Lock()
	defer c.mx.Unlock()
	delete(c.storage, key)
	return nil
}
