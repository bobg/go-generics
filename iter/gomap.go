package iter

type goMapIter[K comparable, V any] struct {
	m        map[K]V
	keysIter Of[K]
}

func (m *goMapIter[K, V]) Next() bool {
	return m.keysIter.Next()
}

func (m *goMapIter[K, V]) Val() Pair[K, V] {
	k := m.keysIter.Val()
	return Pair[K, V]{
		X: k,
		Y: m.m[k],
	}
}

func FromMap[K comparable, V any](m map[K]V) Of[Pair[K, V]] {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return &goMapIter[K, V]{
		m:        m,
		keysIter: FromSlice(keys),
	}
}

func ToMap[K comparable, V any](inp Of[Pair[K, V]]) map[K]V {
	res := make(map[K]V)
	for inp.Next() {
		val := inp.Val()
		res[val.X] = val.Y
	}
	return res
}
