package seqs

import "iter"

// Accum accumulates the result of repeatedly applying a simple function to the elements of an iterator.
// If inp[i] is the ith element of the input
// and out[i] is the ith element of the output,
// then:
//
//	out[0] == inp[0]
//
// and
//
//	out[i+1] == f(out[i], inp[i+1])
func Accum[T any, F ~func(T, T) T](inp iter.Seq[T], f F) iter.Seq[T] {
	res, _ := Accumx(inp, func(a, b T) (T, error) {
		return f(a, b), nil
	})
	return res
}

// Accumx is the extended form of [Accum].
// It accumulates the result of repeatedly applying a function to the elements of an iterator.
// If inp[i] is the ith element of the input
// and out[i] is the ith element of the output,
// then:
//
//	out[0] == inp[0]
//
// and
//
//	out[i+1] == f(out[i], inp[i+1])
//
// The caller gets the resulting iterator and a non-nil pointer to an error.
// After the iterator is fully consumed,
// the caller may dereference the error pointer to check for any error that occurred during iteration.
func Accumx[T any, F ~func(T, T) (T, error)](inp iter.Seq[T], f F) (iter.Seq[T], *error) {
	var (
		val T
		err error
	)

	seq := func(yield func(T) bool) {
		for x := range inp {
			val, err = f(val, x)
			if err != nil {
				return
			}
			if !yield(val) {
				return
			}
		}
	}

	return seq, &err
}
