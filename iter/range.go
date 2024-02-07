//go:build go1.23 || goexperiment.rangefunc

package iter

import "iter"

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

// Seq is an iterator over sequences of individual values.
// When called as seq(yield), seq calls yield(v) for each value v in the sequence,
// stopping early if yield returns false.
//
// This type is defined in the same way as in the standard library,
// but is not identical,
// because Go type aliases cannot (yet?) be used with generic types.
type Seq[V any] func(yield func(V) bool)

// Seq2 is an iterator over sequences of pairs of values, most commonly key-value pairs.
// When called as seq(yield), seq calls yield(k, v) for each pair (k, v) in the sequence,
// stopping early if yield returns false.
//
// This type is defined in the same way as in the standard library,
// but is not identical,
// because Go type aliases cannot (yet?) be used with generic types.
type Seq2[K, V any] func(yield func(K, V) bool)

// Pull converts the “push-style” iterator sequence seq
// into a “pull-style” iterator accessed by the two functions
// next and stop.
//
// Next returns the next value in the sequence
// and a boolean indicating whether the value is valid.
// When the sequence is over, next returns the zero V and false.
// It is valid to call next after reaching the end of the sequence
// or after calling stop. These calls will continue
// to return the zero V and false.
//
// Stop ends the iteration. It must be called when the caller is
// no longer interested in next values and next has not yet
// signaled that the sequence is over (with a false boolean return).
// It is valid to call stop multiple times and when next has
// already returned false.
//
// It is an error to call next or stop from multiple goroutines
// simultaneously.
func Pull[V any](seq Seq[V]) (next func() (V, bool), stop func()) {
	return iter.Pull(iter.Seq[V](seq))
}

// Pull2 converts the “push-style” iterator sequence seq
// into a “pull-style” iterator accessed by the two functions
// next and stop.
//
// Next returns the next pair in the sequence
// and a boolean indicating whether the pair is valid.
// When the sequence is over, next returns a pair of zero values and false.
// It is valid to call next after reaching the end of the sequence
// or after calling stop. These calls will continue
// to return a pair of zero values and false.
//
// Stop ends the iteration. It must be called when the caller is
// no longer interested in next values and next has not yet
// signaled that the sequence is over (with a false boolean return).
// It is valid to call stop multiple times and when next has
// already returned false.
//
// It is an error to call next or stop from multiple goroutines
// simultaneously.
func Pull2[K, V any](seq Seq2[K, V]) (next func() (K, V, bool), stop func()) {
	return iter.Pull2(iter.Seq2[K, V](seq))
}
