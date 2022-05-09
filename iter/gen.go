package iter

func Gen[T any](f func() (T, bool, error)) Of[T] {
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

func Ints(start, delta int) Of[int] {
	n := start
	return Gen(func() (int, bool, error) {
		res := n
		n += delta
		return res, true, nil
	})
}

func Repeat[T any](val T) Of[T] {
	return Gen(func() (T, bool, error) { return val, true, nil })
}
