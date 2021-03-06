package iter

// Map produces an iterator of values transformed from an input iterator by a mapping function.
// If the mapping function returns an error,
// iteration stops and the error is available via the output iterator's Err method.
func Map[T, U any](inp Of[T], f func(T) (U, error)) Of[U] {
	return &mapIter[T, U]{inp: inp, f: f}
}

type mapIter[T, U any] struct {
	inp Of[T]
	f   func(T) (U, error)
	val U
	err error
}

func (m *mapIter[T, U]) Next() bool {
	if !m.inp.Next() {
		m.err = m.inp.Err()
		return false
	}
	m.val, m.err = m.f(m.inp.Val())
	return m.err == nil
}

func (m *mapIter[T, U]) Val() U {
	return m.val
}

func (m *mapIter[T, U]) Err() error {
	return m.err
}
