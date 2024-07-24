package seqs

import (
	"iter"
	"slices"
)

// From creates an iterator over the given items.
func From[T any](items ...T) iter.Seq[T] {
	return slices.Values(items)
}
