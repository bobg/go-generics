package maps

import "github.com/bobg/go-generics/iter"

// Dup makes a (shallow) duplicate of the given map.
func Dup[K comparable, V any](m map[K]V) map[K]V {
	result := make(map[K]V, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}

// Each calls a function on each key-value pair in the given map.
// If the function returns an error, Each exits early with that error.
func Each[K comparable, V any](m map[K]V, f func(K, V) error) error {
	for k, v := range m {
		err := f(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

// FromPairs produces a new map from an iterator of key-value pairs.
func FromPairs[K comparable, V any](pairs iter.Of[iter.Pair[K, V]]) (map[K]V, error) {
	result := make(map[K]V)
	for pairs.Next() {
		p := pairs.Val()
		result[p.X] = p.Y
	}
	return result, pairs.Err()
}

// Invert inverts the key-value pairs in the given map,
// producing a new map with the values as keys and the keys as values.
// If any of the values in the input are duplicates,
// an indeterminate one will survive with its key while the others will be silently dropped.
func Invert[K, V comparable](m map[K]V) map[V]K {
	result := make(map[V]K)
	for k, v := range m {
		result[v] = k
	}
	return result
}

// Equal tests whether two maps are equal to each other.
// The values in the maps must be "comparable."
func Equal[K, V comparable](a, b map[K]V) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		v2, ok := b[k]
		if !ok {
			return false
		}
		if v != v2 {
			return false
		}
	}
	return true
}

// Keys returns a slice of the keys in m.
func Keys[K comparable, V any](m map[K]V) []K {
	result := make([]K, 0, len(m))
	for k := range m {
		result = append(result, k)
	}
	return result
}
