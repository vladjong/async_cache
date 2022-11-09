package asyncache

import (
	"context"

	"github.com/vladjong/async_cache/helpers"
	mutextcache "github.com/vladjong/async_cache/mutext_cache"
	"github.com/vladjong/async_cache/storage"
)

type Cache struct {
	c storage.Cache
}

func NewCache() *Cache {
	return &Cache{
		c: mutextcache.NewCache(),
	}
}

func (c *Cache) Get(ctx context.Context, key string) (string, error) {
	ch := make(chan string)
	go func() {
		defer close(ch)
		v, err := c.c.Get(key)
		if err == nil {
			ch <- v
		}
	}()

	select {
	case <-ctx.Done():
		return "", helpers.ErrTimeout
	case x, ok := <-ch:
		if ok {
			return x, nil
		}
		return "", helpers.ErrNotFound
	}
}

func (c *Cache) Set(ctx context.Context, key, value string) error {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		c.c.Set(key, value)
	}()

	select {
	case <-ctx.Done():
		return helpers.ErrTimeout
	case <-ch:
		return nil
	}
}

func (c *Cache) Delete(ctx context.Context, key string) error {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		c.c.Delete(key)
	}()

	select {
	case <-ctx.Done():
		return helpers.ErrTimeout
	case <-ch:
		return nil
	}
}
