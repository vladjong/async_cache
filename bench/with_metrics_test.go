package bench

import (
	"testing"

	metriccache "github.com/vladjong/async_cache/metric_cache"
)

func Benchmark_MutextWithMetricInt(b *testing.B) {
	c := metriccache.NewCache()
	for i := 0; i < b.N; i++ {
		emulatedLoadWithMetric(c, parallelFactor)
	}
}
