package metriccache

import (
	"sync"

	"github.com/vladjong/async_cache/helpers"
)

type metricCacheInt struct {
	mx      sync.Mutex
	storage map[string]string
	total   int64
}

func NewCache() *metricCacheInt {
	return &metricCacheInt{
		storage: make(map[string]string),
		mx:      sync.Mutex{},
	}
}

func (c *metricCacheInt) Set(key, value string) error {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.storage[key] = value
	c.total++
	return nil
}

func (c *metricCacheInt) Get(key string) (string, error) {
	c.mx.Lock()
	defer c.mx.Unlock()
	value, ok := c.storage[key]
	if !ok {
		return "", helpers.ErrNotFound
	}
	return value, nil
}

func (c *metricCacheInt) Delete(key string) error {
	c.mx.Lock()
	defer c.mx.Unlock()
	delete(c.storage, key)
	c.total--
	return nil
}

func (c *metricCacheInt) TotalAmount() int64 {
	return c.total
}
