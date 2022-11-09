package simplecache

import "github.com/vladjong/async_cache/helpers"

type simpleCache struct {
	storage map[string]string
}

func NewCache() *simpleCache {
	return &simpleCache{
		storage: make(map[string]string),
	}
}

func (s *simpleCache) Set(key, value string) error {
	s.storage[key] = value
	return nil
}

func (s *simpleCache) Get(key string) (string, error) {
	value, ok := s.storage[key]
	if !ok {
		return "", helpers.ErrNotFound
	}
	return value, nil
}

func (s *simpleCache) Delete(key string) error {
	delete(s.storage, key)
	return nil
}
