package seqs

import "iter"

// Filter copies the input iterator to the output,
// including only those elements that cause f to return true.
func Filter[T any, F ~func(T) bool](inp iter.Seq[T], f F) iter.Seq[T] {
	return func(yield func(T) bool) {
		for val := range inp {
			if !f(val) {
				continue
			}
			if !yield(val) {
				return
			}
		}
	}
}

// SkipUntil copies the input iterator to the output,
// discarding the initial elements until the first one that causes f to return true.
// That element and the remaining elements of inp are included in the output,
// and f is not called again.
func SkipUntil[T any, F ~func(T) bool](inp iter.Seq[T], f F) iter.Seq[T] {
	skipping := true
	return Filter(inp, func(val T) bool {
		if !skipping {
			return true
		}
		if f(val) {
			skipping = false
		}
		return !skipping
	})
}
