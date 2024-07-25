package seqs

import "iter"

// Ints produces an infinite iterator over integers beginning at start,
// with each element increasing by delta.
func Ints(start, delta int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for n := start; ; n += delta {
			if !yield(n) {
				return
			}
		}
	}
}

// Repeat produces an infinite iterator repeatedly containing the given value.
func Repeat[T any](val T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for {
			if !yield(val) {
				return
			}
		}
	}
}
