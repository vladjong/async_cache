package test

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vladjong/async_cache/helpers"
	"github.com/vladjong/async_cache/storage"
)

func emulatedLoad(t *testing.T, c storage.Cache, parallelFactor int) {
	wg := sync.WaitGroup{}
	for i := 0; i < parallelFactor; i++ {
		key := fmt.Sprintf("%d-key", i)
		value := fmt.Sprintf("%d-value", i)
		wg.Add(1)
		go func(k string) {
			err := c.Set(k, value)
			assert.NoError(t, err)
			wg.Done()
		}(key)
		wg.Add(1)
		go func(k string, v string) {
			storedValue, err := c.Get(k)
			if !errors.Is(err, helpers.ErrNotFound) {
				assert.Equal(t, v, storedValue)
			}
			wg.Done()
		}(key, value)
		wg.Add(1)
		go func(k string) {
			err := c.Delete(k)
			assert.NoError(t, err)
			wg.Done()
		}(key)
	}
	wg.Wait()
}

func emulatedLoadWithMetric(t *testing.T, c storage.CacheWithMetrics, parallelFactor int) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		emulatedLoad(t, c, parallelFactor)
		wg.Done()
	}()
	var min, max int64
	for i := 0; i < parallelFactor; i++ {
		wg.Add(1)
		go func() {
			total := c.TotalAmount()
			if total > max {
				max = total
			}
			if total < min {
				min = total
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(max, min)
}

func emulatedLoadAsync(t *testing.T, c storage.AsyncCache, parallelFactor int) {
	wg := sync.WaitGroup{}
	ctxBase := context.Background()
	ctx, cancel := context.WithTimeout(ctxBase, time.Second*10)
	defer cancel()
	for i := 0; i < parallelFactor; i++ {
		key := fmt.Sprintf("%d-key", i)
		value := fmt.Sprintf("%d-value", i)
		wg.Add(1)
		go func(k string) {
			err := c.Set(ctx, k, value)
			assert.NoError(t, err)
			wg.Done()
		}(key)
		wg.Add(1)
		go func(k string, v string) {
			storedValue, err := c.Get(ctx, k)
			if !errors.Is(err, helpers.ErrNotFound) {
				assert.Equal(t, v, storedValue)
			}
			wg.Done()
		}(key, value)
		wg.Add(1)
		go func(k string) {
			err := c.Delete(ctx, k)
			assert.NoError(t, err)
			wg.Done()
		}(key)
	}
	wg.Wait()
}
