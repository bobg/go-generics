// Package iter defines an iterator interface,
// a collection of concrete iterator types,
// and some functions for operating on iterators.
// It is also a drop-in replacement for the Go 1.23 standard library package iter
// (a preview of which is available in Go 1.22 when building with GOEXPERIMENT=rangefunc).
package iter

import "context"

// Of is the interface implemented by iterators.
// It is called "Of" so that when qualified with this package name
// and instantiated with a member type,
// it reads naturally: e.g., iter.Of[int].
type Of[T any] interface {
	// Next advances the iterator to its next value and tells whether one is available to read.
	// A true result is necessary before calling Val.
	// Once Next returns false, it must continue returning false.
	Next() bool

	// Val returns the current value of the iterator.
	// Callers must get a true result from Next before calling Val.
	// Repeated calls to Val
	// with no intervening call to Next
	// should return the same value.
	Val() T

	// Err returns the error that this iterator's source encountered during iteration, if any.
	// It may be called only after Next returns false.
	Err() error
}

// Seq is a Go 1.23 iterator over sequences of individual values.
// When called as seq(yield), seq calls yield(v) for each value v in the sequence,
// stopping early if yield returns false.
//
// This type is defined in the same way as in the standard library,
// but is not identical,
// because Go type aliases cannot (yet?) be used with generic types.
type Seq[V any] func(yield func(V) bool)

// Seq2 is a Go 1.23 iterator over sequences of pairs of values, most commonly key-value pairs.
// When called as seq(yield), seq calls yield(k, v) for each pair (k, v) in the sequence,
// stopping early if yield returns false.
//
// This type is defined in the same way as in the standard library,
// but is not identical,
// because Go type aliases cannot (yet?) be used with generic types.
type Seq2[K, V any] func(yield func(K, V) bool)

// All makes a Go 1.23 iterator from an Of[T],
// suitable for use in a one-variable for-range loop.
// To try this in Go 1.22,
// build with the environment variable GOEXPERIMENT set to rangefunc.
// See https://go.dev/wiki/RangefuncExperiment.
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

// AllCount makes a Go 1.23 counting iterator from an Of[T],
// suitable for use in a two-variable for-range loop.
// To try this in Go 1.22,
// build with the environment variable GOEXPERIMENT set to rangefunc.
// See https://go.dev/wiki/RangefuncExperiment.
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

// AllPairs makes a Go 1.23 pair iterator from an Of[Pair[T, U]],
// suitable for use in a two-variable for-range loop.
// To try this in Go 1.22,
// build with the environment variable GOEXPERIMENT set to rangefunc.
// See https://go.dev/wiki/RangefuncExperiment.
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

// FromSeq converts a Go 1.23 iterator into an Of[T].
func FromSeq[T any](seq Seq[T]) Of[T] {
	return Go(func(ch chan<- T) error {
		seq(func(v T) bool {
			ch <- v
			return true
		})
		return nil
	})
}

// FromSeqContext converts a Go 1.23 iterator into an Of[T].
func FromSeqContext[T any](ctx context.Context, seq Seq[T]) Of[T] {
	return Go(func(ch chan<- T) error {
		seq(func(v T) bool {
			select {
			case <-ctx.Done():
				return false
			case ch <- v:
				return true
			}
		})
		return ctx.Err()
	})
}

// FromSeq2 converts a Go 1.23 pair iterator into an Of[Pair[T, U]].
func FromSeq2[T, U any](seq Seq2[T, U]) Of[Pair[T, U]] {
	return Go(func(ch chan<- Pair[T, U]) error {
		seq(func(t T, u U) bool {
			ch <- Pair[T, U]{X: t, Y: u}
			return true
		})
		return nil
	})
}

// FromSeq2Context converts a Go 1.23 pair iterator into an Of[Pair[T, U]].
func FromSeq2Context[T, U any](ctx context.Context, seq Seq2[T, U]) Of[Pair[T, U]] {
	return Go(func(ch chan<- Pair[T, U]) error {
		seq(func(t T, u U) bool {
			select {
			case <-ctx.Done():
				return false
			case ch <- Pair[T, U]{X: t, Y: u}:
				return true
			}
		})
		return ctx.Err()
	})
}
