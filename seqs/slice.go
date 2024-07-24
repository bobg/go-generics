package seqs

import (
	"iter"
	"slices"
)

// FromSlice creates an iterator over the elements of a slice.
func FromSlice[S ~[]T, T any](s S) iter.Seq[T] {
	return slices.Values(s)
}

// From creates an iterator over the given items.
func From[T any](items ...T) iter.Seq[T] {
	return FromSlice(items)
}

// ToSlice consumes the elements of an iterator and returns them as a slice.
// Be careful! The input may be very long or even infinite.
// Consider using FirstN to ensure the input has a reasonable size.
func ToSlice[T any](seq iter.Seq[T]) []T {
	return slices.Collect(seq)
}
