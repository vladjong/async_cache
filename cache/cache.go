package cache

type Cache interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}
