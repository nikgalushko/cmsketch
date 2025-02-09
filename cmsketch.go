package cmsketch

import (
	"math"
)

// Sketch is a Count-Min Sketch data structure that tracks frequency of elements.
type Sketch[T comparable] struct {
	w  uint64
	m  [][]uint64
	h1 func(T) uint64
	h2 func(T) uint64
}

// NewWithDimensions creates a new Sketch with specified dimensions (depth and width).
func NewWithDimensions[T comparable](d, w int) *Sketch[T] {
	matrix := make([][]uint64, d)
	for i := range matrix {
		matrix[i] = make([]uint64, w)
	}

	return &Sketch[T]{m: matrix, h1: hasher[T](), h2: hasher[T](), w: uint64(w)}
}

// NewWithEstimates creates a new Sketch based on error rate (eps) and confidence (delta).
func NewWithEstimates[T comparable](eps, delta float64) *Sketch[T] {
	w := int(math.Ceil(math.E / eps))
	d := int(math.Ceil(math.Log(1.0 / delta)))

	return NewWithDimensions[T](d, w)
}

// Update increments the value associated with the given key by v.
func (s *Sketch[T]) Update(key T, v uint64) {
	h1 := s.h1(key)
	h2 := s.h2(key)
	for i := range s.m {
		h := h1 + (uint64(i) * h2)
		s.m[i][int(h%s.w)] += v
	}
}

// Inc increments the value associated with the given key by 1.
func (s *Sketch[T]) Inc(key T) {
	s.Update(key, 1)
}

// Estimate returns the estimated count of occurrences for the given key.
func (s *Sketch[T]) Estimate(key T) uint64 {
	h1 := s.h1(key)
	h2 := s.h2(key)
	ret := s.m[0][h1%s.w]
	for i := 1; i < len(s.m); i++ {
		h := h1 + (uint64(i) * h2)
		ret = min(ret, s.m[i][int(h%s.w)])
	}

	return ret
}
