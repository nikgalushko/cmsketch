package cmsketch

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"math/rand"
)

func TestAccuracy(t *testing.T) {
	const N = 1_000_000

	s := NewWithEstimates[string](0.000001, 0.0001)
	m := make(map[string]uint64)

	for i := 0; i < N; i++ {
		v := uint64(rand.Intn(N))
		k := strconv.Itoa(i)

		s.Update(k, v)
		m[k] += v
	}

	var miss int
	for i := 0; i < N; i++ {
		k := strconv.Itoa(i)
		v := s.Estimate(k)
		if v != m[k] {
			miss++
		}
	}

	if miss > 10 {
		t.Fatalf("%d > 10", miss)
	}
}

func Benchmark_W_2000_D_10(b *testing.B) {
	dataset := make([]string, 10_000_000)
	for i := range dataset {
		dataset[i] = fmt.Sprintf("%d-%d", time.Now().UnixMilli(), i)
	}
	b.ResetTimer()

	b.Run("update", func(b *testing.B) {
		s := NewWithDimensions[string](10, 2_000)
		for i := range b.N {
			s.Update(dataset[i%len(dataset)], uint64(rand.Int63()))
		}
	})

	b.Run("estimate", func(b *testing.B) {
		s := NewWithDimensions[string](10, 2_000)
		for i := range b.N {
			s.Estimate(dataset[i%len(dataset)])
		}
	})
}

func Benchmark_W_2000000_D_14(b *testing.B) {
	dataset := make([]string, 50_000_000)
	for i := range dataset {
		dataset[i] = fmt.Sprintf("%d-%d", time.Now().UnixMilli(), i)
	}
	b.ResetTimer()

	b.Run("update", func(b *testing.B) {
		s := NewWithDimensions[string](14, 2_000_000)
		for i := range b.N {
			s.Update(dataset[i%len(dataset)], uint64(rand.Int63()))
		}
	})

	b.Run("estimate", func(b *testing.B) {
		s := NewWithDimensions[string](14, 2_000_000)
		for i := range b.N {
			s.Estimate(dataset[i%len(dataset)])
		}
	})
}
