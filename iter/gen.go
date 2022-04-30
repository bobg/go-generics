package iter

func Gen[T any](f func() (T, bool)) Of[T] {
	return &genIter[T]{f: f}
}

type genIter[T any] struct {
	f   func() (T, bool)
	val T
	end bool
}

func (g *genIter[T]) Next() bool {
	if g.end {
		return false
	}
	val, ok := g.f()
	if !ok {
		g.end = true
		return false
	}
	g.val = val
	return true
}

func (g *genIter[T]) Val() T {
	return g.val
}

func Ints(start, delta int) Of[int] {
	n := start
	return Gen(func() (int, bool) {
		res := n
		n += delta
		return res, true
	})
}

func Repeat[T any](val T) Of[T] {
	return Gen(func() (T, bool) { return val, true })
}
