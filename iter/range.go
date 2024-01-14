//go:build go1.23 || rangefunc

package iter

import goiter "iter"

// All makes a Go 1.23 iterator from an Of[T],
// suitable for use in a one-variable for-range loop.
//
// The caller should still check the iterator's Err method after the loop terminates.
func All[T any](inp Of[T]) Seq[T] {
	return func(yield func(T) bool) {
		for inp.Next() {
			if !yield(inp.Val()) {
				return
			}
		}
	}
}

// AllCount makes a counting iterator from an Of[T],
// suitable for use in a two-variable for-range loop.
//
// The caller should still check the iterator's Err method after the loop terminates.
func AllCount[T any](inp Of[T]) Seq2[int, T] {
	return func(yield func(int, T) bool) {
		var i int
		for inp.Next() {
			if !yield(i, inp.Val()) {
				return
			}
			i++
		}
	}
}

// AllPairs makes a pair iterator from an Of[Pair[T, U]],
// suitable for use in a two-variable for-range loop.
//
// The caller should still check the iterator's Err method after the loop terminates.
func AllPairs[T, U any](inp Of[Pair[T, U]]) Seq2[T, U] {
	return func(yield func(T, U) bool) {
		for inp.Next() {
			p := inp.Val()
			if !yield(p.X, p.Y) {
				return
			}
		}
	}
}

type Seq[V any] func(yield func(V) bool)

type Seq2[K, V any] func(yield func(K, V) bool)

func Pull[V any](seq goiter.Seq[V]) (next func() (V, bool), stop func()) {
	return goiter.Pull(seq)
}

func Pull2[K, V any](seq goiter.Seq2[K, V]) (next func() (K, V, bool), stop func()) {
	return goiter.Pull2(seq)
}
