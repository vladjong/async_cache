package async

import (
	"github.com/vladjong/async_cache/cache"
	"github.com/vladjong/async_cache/storage"
)

type Cache struct {
	c *storage.Cache
}

func NewCache() *Cache {
	return &Cache{
		c: cache.NewCache()
	}
}
