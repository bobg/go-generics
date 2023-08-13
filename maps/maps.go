// Package maps contains functions for working with Go maps:
// duplicating, inverting, constructing, and iterating over them,
// as well as testing their equality.
//
// The maps package is a drop-in replacement
// for the maps package
// added to the Go stdlib
// in Go 1.21 (https://go.dev/doc/go1.21#maps).
package maps

import "github.com/bobg/go-generics/v2/iter"

// Each calls a function on each key-value pair in the given map.
func Each[M ~map[K]V, K comparable, V any](m M, f func(K, V)) {
	_ = Eachx(m, func(k K, v V) error {
		f(k, v)
		return nil
	})
}

// Eachx is the extended form of [Each].
// It calls a function on each key-value pair in the given map.
// If the function returns an error, Each exits early with that error.
func Eachx[M ~map[K]V, K comparable, V any](m M, f func(K, V) error) error {
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
func Invert[M ~map[K]V, K, V comparable](m M) map[V]K {
	result := make(map[V]K)
	for k, v := range m {
		result[v] = k
	}
	return result
}

// InvertMulti inverts the key-value pairs in the map.
// It is like Invert but handles duplicate values:
// the key slice contains all the keys that map to the same value.
func InvertMulti[M ~map[K]V, K, V comparable](m M) map[V][]K {
	result := make(map[V][]K)
	for k, v := range m {
		result[v] = append(result[v], k)
	}
	return result
}
