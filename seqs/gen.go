package seqs

import "iter"

// Gen produces an iterator of values obtained by repeatedly calling f.
// If f returns an error,
// iteration stops and the error is available via the iterator's Err method.
// Otherwise, each call to f should return a value and a true boolean.
// When f returns a false boolean, it signals the normal end of iteration.
func Gen[T any, F ~func() (T, bool, error)](f F) (iter.Seq[T], *error) {
	var err error

	g := func(yield func(T) bool) {
		for {
			var (
				val T
				ok  bool
			)

			val, ok, err = f()
			if err != nil || !ok {
				return
			}
			if !yield(val) {
				return
			}
		}
	}

	return g, &err
}

// Ints produces an infinite iterator over integers beginning at start,
// with each element increasing by delta.
func Ints(start, delta int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for n := start; ; n += delta {
			if !yield(n) {
				return
			}
		}
	}
}

// Repeat produces an infinite iterator repeatedly containing the given value.
func Repeat[T any](val T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for {
			if !yield(val) {
				return
			}
		}
	}
}
