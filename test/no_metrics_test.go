package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Cache(t *testing.T) {
	t.Parallel()
	testCache := mu

	t.Run("correctly stored value", func(t *testing.T) {
		t.Parallel()
		key := "someKey"
		value := "someValue"
		err := testCache.Set(key, value)
		assert.NoError(t, err)
		storedValue, err := testCache.Get(key)
		assert.NoError(t, err)
		assert.Equal(t, value, storedValue)
	})

	t.Run("no data races", func(t *testing.T) {
		t.Parallel()
		parallelFactor := 100_100_0
		emulatedLoad(t, testCache, parallelFactor)
	})
}
