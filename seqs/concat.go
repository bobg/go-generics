package seqs

import "iter"

// Concat concatenates the members of the input iterators.
func Concat[T any](inps ...iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, inp := range inps {
			if inp == nil {
				continue
			}
			for val := range inp {
				if !yield(val) {
					return
				}
			}
		}
	}
}
