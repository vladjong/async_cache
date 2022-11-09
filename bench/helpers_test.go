package bench

import (
	"errors"
	"fmt"
	"sync"

	"github.com/vladjong/async_cache/helpers"
	"github.com/vladjong/async_cache/storage"
)

func emulatedLoad(c storage.Cache, parallelFactor int) {
	wg := sync.WaitGroup{}
	for i := 0; i < parallelFactor/10; i++ {
		key := fmt.Sprintf("%d-key", i)
		value := fmt.Sprintf("%d-value", i)
		wg.Add(1)
		go func(k string) {
			err := c.Set(k, value)
			if err != nil {
				panic(err)
			}
			wg.Done()
		}(key)
		wg.Add(1)
		go func(k string, v string) {
			_, err := c.Get(k)
			if err != nil && !errors.Is(err, helpers.ErrNotFound) {
				panic(err)
			}
			wg.Done()
		}(key, value)
		wg.Add(1)
		go func(k string) {
			err := c.Delete(k)
			if err != nil {
				panic(err)
			}
			wg.Done()
		}(key)
	}
	wg.Wait()
}

func emulatedReadLoad(c storage.Cache, parallelFactor int) {
	wg := sync.WaitGroup{}
	for i := 0; i < parallelFactor/10; i++ {
		key := fmt.Sprintf("%d-key", i)
		value := fmt.Sprintf("%d-value", i)
		wg.Add(1)
		go func(k string) {
			err := c.Set(k, value)
			if err != nil {
				panic(err)
			}
			wg.Done()
		}(key)
		wg.Add(1)
		go func(k string) {
			err := c.Delete(k)
			if err != nil {
				panic(err)
			}
			wg.Done()
		}(key)
	}
	for i := 0; i < parallelFactor; i++ {
		key := fmt.Sprintf("%d-key", i)
		value := fmt.Sprintf("%d-value", i)
		wg.Add(1)
		go func(k string, v string) {
			_, err := c.Get(k)
			if err != nil && !errors.Is(err, helpers.ErrNotFound) {
				panic(err)
			}
			wg.Done()
		}(key, value)
	}
	wg.Wait()
}

func emulatedLoadWithMetric(c storage.CacheWithMetrics, parallelFactor int) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		emulatedLoad(c, parallelFactor)
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
}
