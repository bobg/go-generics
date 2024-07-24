package seqs

import "iter"

// Map produces an iterator of values transformed from an input iterator by a simple mapping function.
func Map[T, U any, F ~func(T) U](inp iter.Seq[T], f F) iter.Seq[U] {
	m, _ := Mapx(inp, func(val T) (U, error) {
		return f(val), nil
	})
	return m
}

// Mapx is the extended form of [Map].
// It produces an iterator of values transformed from an input iterator by a mapping function.
// If the mapping function returns an error,
// iteration stops and the error is available via the output iterator's Err method.
func Mapx[T, U any, F ~func(T) (U, error)](inp iter.Seq[T], f F) (iter.Seq[U], *error) {
	var err error

	g := func(yield func(U) bool) {
		for val := range inp {
			var out U

			out, err = f(val)
			if err != nil {
				return
			}
			if !yield(out) {
				return
			}
		}
	}

	return g, &err
}
