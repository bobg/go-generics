package iter

func Map[T, U any](inp Of[T], f func(T) U) Of[U] {
	return &mapIter[T, U]{inp: inp, f: f}
}

type mapIter[T, U any] struct {
	inp Of[T]
	f   func(T) U
}

func (m *mapIter[T, U]) Next() bool {
	return m.inp.Next()
}

func (m *mapIter[T, U]) Val() U {
	return m.f(m.inp.Val())
}
