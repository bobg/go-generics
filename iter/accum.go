package iter

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
func Accum[T any](inp Of[T], f func(T, T) T) Of[T] {
	return Accumx(inp, func(a, b T) (T, error) {
		return f(a, b), nil
	})
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
func Accumx[T any](inp Of[T], f func(T, T) (T, error)) Of[T] {
	return &accumIter[T]{
		inp:   inp,
		f:     f,
		first: true,
	}
}

type accumIter[T any] struct {
	inp   Of[T]
	f     func(T, T) (T, error)
	first bool
	val   T
	err   error
}

func (a *accumIter[T]) Next() bool {
	if a.inp.Next() {
		if a.first {
			a.val = a.inp.Val()
			a.first = false
		} else {
			a.val, a.err = a.f(a.val, a.inp.Val())
		}
		return a.err == nil
	}
	a.err = a.inp.Err()
	return false
}

func (a *accumIter[T]) Val() T {
	return a.val
}

func (a *accumIter[T]) Err() error {
	return a.err
}
