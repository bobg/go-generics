package seqs

import "iter"

// Zip takes two iterators and produces a new iterator containing pairs of corresponding elements.
// If one input iterator ends before the other,
// Zip produces zero values of the appropriate type in constructing pairs.
func Zip[T, U any](t iter.Seq[T], u iter.Seq[U]) iter.Seq2[T, U] {
	return func(yield func(T, U) bool) {
		tnext, tstop := iter.Pull(t)
		defer tstop()

		unext, ustop := iter.Pull(u)
		defer ustop()

		var tdone, udone bool

		for {
			var (
				tval T
				uval U
			)

			if !tdone {
				tval, tdone = tnext()
			}
			if !udone {
				uval, udone = unext()
			}
			if tdone && udone {
				return
			}

			if !yield(tval, uval) {
				return
			}
		}
	}
}
