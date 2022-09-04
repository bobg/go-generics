// Package slices contains utility functions for working with slices.
// It encapsulates hard-to-remember idioms for inserting and removing elements;
// it adds the ability to index from the right end of a slice using negative integers
// (for example, Get(s, -1) is the same as s[len(s)-1]),
// and it includes Map, Filter,
// and a few other such functions
// for processing slice elements with callbacks.
package slices

import "sort"

// Get gets the idx'th element of s.
//
// If idx < 0 it counts from the end of s.
func Get[T any](s []T, idx int) T {
	if idx < 0 {
		idx += len(s)
	}
	return s[idx]
}

// Put puts a given value into the idx'th location in s.
//
// If idx < 0 it counts from the end of s.
//
// The input slice is modified.
func Put[T any](s []T, idx int, val T) {
	if idx < 0 {
		idx += len(s)
	}
	s[idx] = val
}

// Append is the same as Go's builtin append and is included for completeness.
func Append[T any](s []T, vals ...T) []T {
	return append(s, vals...)
}

// Insert inserts the given values at the idx'th location in s and returns the result.
// After the insert, the first new value has position idx.
//
// If idx < 0, it counts from the end of s.
//
// The input slice is modified.
//
// Example: Insert([x, y, z], 1, a, b, c) -> [x, a, b, c, y, z]
func Insert[T any](s []T, idx int, vals ...T) []T {
	if idx < 0 {
		idx += len(s)
	}
	return insert(s, idx, vals...)
}

func insert[T any](s []T, idx int, vals ...T) []T {
	// Make s long enough.
	s = append(s, vals...)

	// Make space in s at the right position.
	copy(s[idx+len(vals):], s[idx:])

	// Put values in the right spot.
	copy(s[idx:], vals)

	return s
}

// ReplaceN replaces the n values of s beginning at position idx with the given values.
// After the replace, the first new value has position idx.
//
// If idx < 0, it counts from the end of s.
//
// The input slice is modified.
func ReplaceN[T any](s []T, idx, n int, vals ...T) []T {
	if idx < 0 {
		idx += len(s)
	}
	return replaceN(s, idx, n, vals...)
}

func replaceN[T any](s []T, idx, n int, vals ...T) []T {
	if n > len(vals) {
		// Removing more items than inserting.
		s = removeN(s, idx, n-len(vals))
	} else if n < len(vals) {
		// Inserting more items than removing.
		delta := len(vals) - n
		s = insert(s, idx, vals[:delta]...)
		idx += delta
		vals = vals[delta:]
	}
	copy(s[idx:], vals)

	return s
}

// ReplaceTo replaces the values of s beginning at from and ending before to with the given values.
// After the replace, the first new value has position from.
//
// If from < 0 it counts from the end of s.
// If to <= 0 it counts from the end of s.
//
// The input slice is modified.
func ReplaceTo[T any](s []T, from, to int, vals ...T) []T {
	if from < 0 {
		from += len(s)
	}
	if to < 0 {
		to += len(s)
	} else if to == 0 {
		to = len(s)
	}
	return replaceN(s, from, to-from, vals...)
}

// RemoveN removes n items from s beginning at position idx and returns the result.
//
// If idx < 0 it counts from the end of s.
//
// The input slice is modified.
//
// Example: RemoveN([a, b, c, d], 1, 2) -> [a, d]
func RemoveN[T any](s []T, idx, n int) []T {
	if idx < 0 {
		idx += len(s)
	}
	return removeN(s, idx, n)
}

func removeN[T any](s []T, idx, n int) []T {
	copy(s[idx:], s[idx+n:])
	newlen := len(s) - n
	return s[:newlen]
}

// RemoveTo removes items from s beginning at position from and ending before position to.
// It returns the result.
//
// If from < 0 it counts from the end of s.
// If to <= 0 it counts from the end of s.
//
// The input slice is modified.
//
// Example: RemoveTo([a, b, c, d], 1, 3) -> [a, d]
func RemoveTo[T any](s []T, from, to int) []T {
	if from < 0 {
		from += len(s)
	}
	if to < 0 {
		to += len(s)
	} else if to == 0 {
		to = len(s)
	}
	return removeN(s, from, to-from)
}

// Prefix returns s up to but not including position idx.
//
// If idx < 0 it counts from the end of s.
func Prefix[T any](s []T, idx int) []T {
	if idx < 0 {
		idx += len(s)
	}
	return s[:idx]
}

// Suffix returns s excluding elements before position idx.
//
// If idx < 0 it counts from the end of s.
func Suffix[T any](s []T, idx int) []T {
	if idx < 0 {
		idx += len(s)
	}
	return s[idx:]
}

// SliceN returns n elements of s beginning at position idx.
//
// If idx < 0 it counts from the end of s.
func SliceN[T any](s []T, idx, n int) []T {
	if idx < 0 {
		idx += len(s)
	}
	return s[idx : idx+n]
}

// SliceTo returns the elements of s beginning at position from and ending before position to.
//
// If from < 0 it counts from the end of s.
// If to <= 0 it counts from the end of s.
func SliceTo[T any](s []T, from, to int) []T {
	if from < 0 {
		from += len(s)
	}
	if to < 0 {
		to += len(s)
	} else if to == 0 {
		to = len(s)
	}
	return s[from:to]
}

// Each runs a function on each item of a slice,
// passing the index and the item to the function.
// If any call to the function returns an error,
// Each stops looping and exits with the error.
func Each[T any](s []T, f func(int, T) error) error {
	for i, val := range s {
		if err := f(i, val); err != nil {
			return err
		}
	}
	return nil
}

// Map runs a function on each item of a slice,
// accumulating results in a new slice.
// If any call to the function returns an error,
// Map stops looping and exits with the error.
func Map[T, U any](s []T, f func(int, T) (U, error)) ([]U, error) {
	result := make([]U, 0, len(s))
	for i, val := range s {
		u, err := f(i, val)
		if err != nil {
			return nil, err
		}
		result = append(result, u)
	}
	return result, nil
}

// Accum accumulates the result of repeatedly applying a function to the elements of a slice.
//
// If the slice has length 0, the result is the zero value of type T.
// If the slice has length 1, the result is s[0].
// Otherwise, the result is R[len(s)-1],
// where R[0] is s[0]
// and R[n+1] = f(R[n], s[n+1]).
func Accum[T any](s []T, f func(T, T) (T, error)) (T, error) {
	if len(s) == 0 {
		var zero T
		return zero, nil
	}
	result := s[0]
	for i := 1; i < len(s); i++ {
		var err error
		result, err = f(result, s[i])
		if err != nil {
			return result, err
		}
	}
	return result, nil
}

// Filter calls a predicate for each element of a slice,
// returning a slice of those elements for which the predicate returned true.
func Filter[T any](s []T, f func(T) (bool, error)) ([]T, error) {
	var result []T
	for _, val := range s {
		ok, err := f(val)
		if err != nil {
			return nil, err
		}
		if !ok {
			continue
		}
		result = append(result, val)
	}
	return result, nil
}

// Group partitions the elements of a slice into groups.
// It does this by calling a grouping function on each element,
// which produces a grouping key.
// The result is a map of group keys to slices of elements having that key.
func Group[T any, K comparable](s []T, f func(T) (K, error)) (map[K][]T, error) {
	result := make(map[K][]T)
	for _, val := range s {
		key, err := f(val)
		if err != nil {
			return nil, err
		}
		result[key] = append(result[key], val)
	}
	return result, nil
}

// Rotate rotates a slice in place by n places to the right.
// (With negative n, it's to the left.)
// Example: Rotate([D, E, A, B, C], 3) -> [A, B, C, D, E]
func Rotate[T any](s []T, n int) {
	if n < 0 {
		// Convert left-rotation to right-rotation.
		n = -n
		n %= len(s)
		n = len(s) - n
	} else {
		n %= len(s)
	}
	if n == 0 {
		return
	}
	tmp := make([]T, n)
	copy(tmp, s[len(s)-n:])
	copy(s[n:], s)
	copy(s, tmp)
}

// Dup makes a (shallow) duplicate of the given slice.
func Dup[T any](s []T) []T {
	if len(s) == 0 {
		return nil
	}
	result := make([]T, len(s))
	copy(result, s)
	return result
}

// KeyedSort sorts the given slice according to the ordering of the given keys,
// whose items must map 1:1 with the slice.
// It is an error if keys.Len() != len(slice).
//
// Both arguments end up sorted in place:
// keys according to its Less method,
// and slice by mirroring the reordering that happens in keys.
func KeyedSort[T any](slice []T, keys sort.Interface) {
	ks := keyedSorter[T]{
		keys:  keys,
		slice: slice,
	}
	sort.Sort(ks)
}

// KeyedSorter allows sorting a slice according to the order of a set of sort keys.
// It works by sorting a [sort.Interface] containing sort keys
// that must map 1:1 with the items of the slice you wish to sort.
// (It is an error for Keys.Len() to differ from len(Slice).)
// Any reordering applied to Keys is also applied to Slice.
type keyedSorter[T any] struct {
	keys  sort.Interface
	slice []T
}

func (k keyedSorter[T]) Len() int           { return len(k.slice) }
func (k keyedSorter[T]) Less(i, j int) bool { return k.keys.Less(i, j) }
func (k keyedSorter[T]) Swap(i, j int) {
	k.keys.Swap(i, j)
	k.slice[i], k.slice[j] = k.slice[j], k.slice[i]
}
