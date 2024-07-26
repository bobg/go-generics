package seqs

import (
	"cmp"
	"iter"
	"slices"
)

// From creates an iterator over the given items.
func From[T any](items ...T) iter.Seq[T] {
	return slices.Values(items)
}

// Pair is a simple, generic pair of values.
type Pair[T, U any] struct {
	X T
	Y U
}

// ToPairs changes an [iter.Seq2] to an [iter.Seq] of [Pair]s.
func ToPairs[T, U any](inp iter.Seq2[T, U]) iter.Seq[Pair[T, U]] {
	return func(yield func(Pair[T, U]) bool) {
		for x, y := range inp {
			if !yield(Pair[T, U]{X: x, Y: y}) {
				return
			}
		}
	}
}

// Left changes an [iter.Seq2] to an [iter.Seq] by dropping the second value.
func Left[T, U any](inp iter.Seq2[T, U]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for x := range inp {
			if !yield(x) {
				return
			}
		}
	}
}

// Right changes an [iter.Seq2] to an [iter.Seq] by dropping the first value.
func Right[T, U any](inp iter.Seq2[T, U]) iter.Seq[U] {
	return func(yield func(U) bool) {
		for _, y := range inp {
			if !yield(y) {
				return
			}
		}
	}
}

// Enumerate changes an [iter.Seq] to an [iter.Seq2] of (index, val) pairs.
func Enumerate[T any](inp iter.Seq[T]) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		i := 0
		for x := range inp {
			if !yield(i, x) {
				return
			}
			i++
		}
	}
}

// FromPairs changes an [iter.Seq] of [Pair]s to an [iter.Seq2].
func FromPairs[T, U any](inp iter.Seq[Pair[T, U]]) iter.Seq2[T, U] {
	return func(yield func(T, U) bool) {
		for val := range inp {
			if !yield(val.X, val.Y) {
				return
			}
		}
	}
}

// Compare performs an elementwise comparison of two sequences.
// It returns the result of [cmp.Compare] on the first pair of unequal elements.
// If a ends before b, Compare returns -1.
// If b ends before a, Compare returns 1.
// If the sequences are equal, Compare returns 0.
func Compare[T cmp.Ordered](a, b iter.Seq[T]) int {
	return CompareFunc(a, b, cmp.Compare)
}

// CompareFunc performs an elementwise comparison of two sequences, using a custom comparison function.
// The function should return a negative number if the first argument is less than the second,
// a positive number if the first argument is greater than the second,
// and zero if the arguments are equal.
//
// CompareFunc returns the result of f on the first pair of unequal elements.
// If a ends before b, CompareFunc returns -1.
// If b ends before a, CompareFunc returns 1.
// If the sequences are equal, CompareFunc returns 0.
func CompareFunc[T any](a, b iter.Seq[T], f func(T, T) int) int {
	anext, astop := iter.Pull(a)
	defer astop()

	bnext, bstop := iter.Pull(b)
	defer bstop()

	aOK, bOK := true, true

	var aVal, bVal T

	for {
		if aOK {
			aVal, aOK = anext()
		}
		if bOK {
			bVal, bOK = bnext()
		}
		if !aOK && !bOK {
			return 0
		}
		if !aOK {
			return -1
		}
		if !bOK {
			return 1
		}
		if cmp := f(aVal, bVal); cmp != 0 {
			return cmp
		}
	}
}

// Empty is an empty sequence that can be used where an [iter.Seq] is expected.
// Usage note: you generally don't want to call this function,
// just refer to it as Empty[typename].
func Empty[T any](func(T) bool) {}

// Empty2 is an empty sequence that can be used where an [iter.Seq2] is expected.
// Usage note: you generally don't want to call this function,
// just refer to it as Empty2[typename1, typename2].
func Empty2[T, U any](func(T, U) bool) {}
