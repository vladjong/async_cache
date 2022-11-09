package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	asyncache "github.com/vladjong/async_cache/asyn_cache"
	"github.com/vladjong/async_cache/helpers"
)

func TestAsync(t *testing.T) {
	t.Parallel()
	ac := asyncache.NewCache()
	t.Run("correctly stored value", func(t *testing.T) {
		to := time.Microsecond / 10
		key := "k"
		val := "v"
		ctxBase := context.Background()
		ctx, cancel := context.WithTimeout(ctxBase, to)
		defer cancel()
		err := ac.Set(ctx, key, val)
		fmt.Println(err)
		if err != helpers.ErrTimeout {
			t.Error("Expected timeout")
		}
		to = time.Millisecond * 2
		ctx, cancel = context.WithTimeout(ctxBase, to)
		defer cancel()
		err = ac.Set(ctx, key, val)
		if err != nil {
			t.Errorf("Expected Set %v", err)
		}
		storedValue, err := ac.Get(ctx, key)
		assert.NoError(t, err)
		assert.Equal(t, val, storedValue)
		err = ac.Delete(ctx, key)
		assert.NoError(t, err)
	})

	t.Run("no data races", func(t *testing.T) {
		t.Parallel()
		parallelFactor := 100_100_0
		emulatedLoadAsync(t, ac, parallelFactor)
	})
}
