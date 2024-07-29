package seqs

import (
	"iter"

	"github.com/bobg/go-generics/v4/internal"
)

// FirstN produces an iterator containing the first n elements of the input
// (or all of the input, if there are fewer than n elements).
// Remaining elements of the input are not consumed.
// It is the caller's responsibility to release any associated resources.
func FirstN[T any](inp iter.Seq[T], n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i, val := range Enumerate(inp) {
			if i >= n {
				return
			}
			if !yield(val) {
				return
			}
		}
	}
}

// LastN produces a slice containing the last n elements of the input iterator
// (or all of the input, if there are fewer than n elements).
// There is no guarantee that any elements will ever be produced:
// the input iterator may be infinite!
func LastN[T any, S ~func(func(T) bool)](inp S, n int) []T {
	var (
		buf   = make([]T, 0, n)
		start = 0
	)
	for val := range inp {
		if len(buf) < n {
			buf = append(buf, val)
			continue
		}
		buf[start] = val
		start = (start + 1) % n
	}
	internal.RotateSlice(buf, -start)
	return buf
}

// SkipN copies the input iterator to the output,
// skipping the first N elements.
func SkipN[T any](inp iter.Seq[T], n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i, val := range Enumerate(inp) {
			if i < n {
				continue
			}
			if !yield(val) {
				return
			}
		}
	}
}
