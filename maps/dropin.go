package maps

import "maps"

// This file contains entrypoints for each of the functions in the standard Go maps package.

// Equal reports whether two maps contain the same key/value pairs.
// Values are compared using ==.
func Equal[M1, M2 ~map[K]V, K, V comparable](m1 M1, m2 M2) bool {
	return maps.Equal(m1, m2)
}

// EqualFunc is like Equal, but compares values using eq.
// Keys are still compared with ==.
func EqualFunc[M1 ~map[K]V1, M2 ~map[K]V2, K comparable, V1, V2 any](m1 M1, m2 M2, eq func(V1, V2) bool) bool {
	return maps.EqualFunc(m1, m2, eq)
}

// Clone returns a copy of m.  This is a shallow clone:
// the new keys and values are set using ordinary assignment.
func Clone[M ~map[K]V, K comparable, V any](m M) M {
	return maps.Clone(m)
}

// Copy copies all key/value pairs in src adding them to dst.
// When a key in src is already present in dst,
// the value in dst will be overwritten by the value associated
// with the key in src.
func Copy[M1 ~map[K]V, M2 ~map[K]V, K comparable, V any](dst M1, src M2) {
	maps.Copy(dst, src)
}

// DeleteFunc deletes any key/value pairs from m for which del returns true.
func DeleteFunc[M ~map[K]V, K comparable, V any](m M, del func(K, V) bool) {
	maps.DeleteFunc(m, del)
}
