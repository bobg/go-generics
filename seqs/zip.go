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

		tOK, uOK := true, true

		for {
			var (
				tval T
				uval U
			)

			if tOK {
				tval, tOK = tnext()
			}
			if uOK {
				uval, uOK = unext()
			}
			if !tOK && !uOK {
				return
			}

			if !yield(tval, uval) {
				return
			}
		}
	}
}
