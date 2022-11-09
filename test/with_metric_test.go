package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	metriccache "github.com/vladjong/async_cache/metric_cache"
)

func Test_CacheWithMetric(t *testing.T) {
	t.Parallel()
	testCache := metriccache.NewCache()

	t.Run("correctly stored value", func(t *testing.T) {
		t.Parallel()
		key := "someKey"
		value := "someValue"
		err := testCache.Set(key, value)
		assert.NoError(t, err)
		storedValue, err := testCache.Get(key)
		assert.NoError(t, err)
		assert.Equal(t, value, storedValue)
		err = testCache.Delete(key)
		assert.NoError(t, err)
	})

	t.Run("no data races", func(t *testing.T) {
		t.Parallel()
		parallelFactor := 100_100_0
		emulatedLoadWithMetric(t, testCache, parallelFactor)
	})
}
