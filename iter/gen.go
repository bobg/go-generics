package iter

// Gen produces an iterator of values obtained by repeatedly calling f.
// If f returns an error,
// iteration stops and the error is available via the iterator's Err method.
// Otherwise, each call to f should return a value and a true boolean.
// When f returns a false boolean, it signals the normal end of iteration.
func Gen[F ~func() (T, bool, error), T any](f F) Of[T] {
	return &genIter[T]{f: f}
}

type genIter[T any] struct {
	f   func() (T, bool, error)
	val T
	end bool
	err error
}

func (g *genIter[T]) Next() bool {
	if g.end {
		return false
	}
	val, ok, err := g.f()
	if err != nil || !ok {
		g.end = true
		g.err = err
		return false
	}
	g.val = val
	return true
}

func (g *genIter[T]) Val() T {
	return g.val
}

func (g *genIter[T]) Err() error {
	return g.err
}

// Ints produces an infinite iterator over integers beginning at start,
// with each element increasing by delta.
func Ints(start, delta int) Of[int] {
	n := start
	return Gen(func() (int, bool, error) {
		res := n
		n += delta
		return res, true, nil
	})
}

// Repeat produces an infinite iterator repeatedly containing the given value.
func Repeat[T any](val T) Of[T] {
	return Gen(func() (T, bool, error) { return val, true, nil })
}
