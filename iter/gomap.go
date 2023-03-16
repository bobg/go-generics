package iter

type goMapIter[K comparable, V any] struct {
	m map[K]V

	// Note: keysIter is the result of FromSlice and so its Err() is always nil.
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

func (*goMapIter[K, V]) Err() error {
	return nil
}

// FromMap produces an iterator of key-value pairs from a Go map.
func FromMap[K comparable, V any](m map[K]V) Of[Pair[K, V]] {
	return &goMapIter[K, V]{
		m:        m,
		keysIter: FromMapKeys(m),
	}
}

// FromMapKeys produces an iterator over the keys of a Go map.
func FromMapKeys[K comparable, V any](m map[K]V) Of[K] {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return FromSlice(keys)
}

// ToMap consumes an iterator of key-value pairs and produces a Go map of the values.
// All but the last of any pairs with duplicate keys are discarded.
// Be careful! The input may be very long or even infinite.
// Consider using FirstN to ensure the input has a reasonable size.
func ToMap[K comparable, V any](inp Of[Pair[K, V]]) (map[K]V, error) {
	res := make(map[K]V)
	for inp.Next() {
		val := inp.Val()
		res[val.X] = val.Y
	}
	return res, inp.Err()
}
