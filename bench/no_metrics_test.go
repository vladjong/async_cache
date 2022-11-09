package bench

import (
	"testing"

	mutextcache "github.com/vladjong/async_cache/mutext_cache"
	simplecache "github.com/vladjong/async_cache/simple_cache"
)

const parallelFactor = 100_000

func Benchmark_NoMutex(b *testing.B) {
	b.Skip("panic in NoMutex")
	c := simplecache.NewCache()
	for i := 0; i < b.N; i++ {
		emulatedLoad(c, parallelFactor)
	}
}

func Benchmark_RWMutexLoad(b *testing.B) {
	c := mutextcache.NewCache()
	for i := 0; i < b.N; i++ {
		emulatedLoad(c, parallelFactor)
	}
}

func Benchmark_RWMutexReadLoad(b *testing.B) {
	c := mutextcache.NewCache()
	for i := 0; i < b.N; i++ {
		emulatedReadLoad(c, parallelFactor)
	}
}
