// Package slices contains utility functions for working with slices.
// It encapsulates hard-to-remember idioms for inserting and removing elements;
// it adds the ability to index from the right end of a slice using negative integers
// (for example, Get(s, -1) is the same as s[len(s)-1]),
// and it includes Map, Filter,
// and a few other such functions
// for processing slice elements with callbacks.
//
// This package is a drop-in replacement
// for the slices package
// added to the Go stdlib
// in Go 1.21 (https://go.dev/doc/go1.21#slices).
// There is one difference:
// this version of slices
// allows the index value passed to `Insert`, `Delete`, and `Replace`
// to be negative for counting backward from the end of the slice.
package slices

import (
	"slices"
	"sort"
)

// Get gets the idx'th element of s.
//
// If idx < 0 it counts from the end of s.
func Get[T any, S ~[]T](s S, idx int) T {
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
func Put[T any, S ~[]T](s S, idx int, val T) {
	if idx < 0 {
		idx += len(s)
	}
	s[idx] = val
}

// Append is the same as Go's builtin append and is included for completeness.
func Append[T any, S ~[]T](s S, vals ...T) S {
	return append(s, vals...)
}

// Insert inserts the given values at the idx'th location in s and returns the result.
// After the insert, the first new value has position idx.
//
// If idx < 0, it counts from the end of s.
// (This is a change from the behavior of Go's standard slices.Insert.)
//
// The input slice is modified.
//
// Example: Insert([x, y, z], 1, a, b, c) -> [x, a, b, c, y, z]
func Insert[S ~[]E, E any](s S, idx int, vals ...E) S {
	if idx < 0 {
		idx += len(s)
	}
	return slices.Insert(s, idx, vals...)
}

// Delete removes the elements s[i:j] from s, returning the modified slice.
// Delete panics if s[i:j] is not a valid slice of s.
// Delete is O(len(s)-j), so if many items must be deleted, it is better to
// make a single call deleting them all together than to delete one at a time.
// Delete might not modify the elements s[len(s)-(j-i):len(s)]. If those
// elements contain pointers you might consider zeroing those elements so that
// objects they reference can be garbage collected.
//
// If i < 0 it counts from the end of s.
// If j <= 0 it counts from the end of s.
// (This is a change from the behavior of Go's standard slices.Delete.)
func Delete[S ~[]E, E any](s S, i, j int) S {
	return RemoveTo(s, i, j)
}

// Replace replaces the elements s[i:j] by the given v, and returns the
// modified slice. Replace panics if s[i:j] is not a valid slice of s.
//
// If i < 0 it counts from the end of s.
// If j <= 0 it counts from the end of s.
// (This is a change from the behavior of Go's standard slices.Replace.)
func Replace[S ~[]E, E any](s S, i, j int, v ...E) S {
	return ReplaceTo(s, i, j, v...)
}

// ReplaceN replaces the n values of s beginning at position idx with the given values.
// After the replace, the first new value has position idx.
//
// If idx < 0, it counts from the end of s.
//
// The input slice is modified.
func ReplaceN[T any, S ~[]T](s S, idx, n int, vals ...T) S {
	if idx < 0 {
		idx += len(s)
	}
	return slices.Replace(s, idx, idx+n, vals...)
}

// ReplaceTo replaces the values of s beginning at from and ending before to with the given values.
// After the replace, the first new value has position from.
//
// If from < 0 it counts from the end of s.
// If to <= 0 it counts from the end of s.
//
// The input slice is modified.
func ReplaceTo[T any, S ~[]T](s S, from, to int, vals ...T) S {
	if from < 0 {
		from += len(s)
	}
	if to < 0 {
		to += len(s)
	} else if to == 0 {
		to = len(s)
	}
	return slices.Replace(s, from, to, vals...)
}

// RemoveN removes n items from s beginning at position idx and returns the result.
//
// If idx < 0 it counts from the end of s.
//
// The input slice is modified.
//
// Example: RemoveN([a, b, c, d], 1, 2) -> [a, d]
func RemoveN[T any, S ~[]T](s S, idx, n int) S {
	if idx < 0 {
		idx += len(s)
	}
	return slices.Delete(s, idx, idx+n)
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
func RemoveTo[T any, S ~[]T](s S, from, to int) S {
	if from < 0 {
		from += len(s)
	}
	if to < 0 {
		to += len(s)
	} else if to == 0 {
		to = len(s)
	}
	return slices.Delete(s, from, to)
}

// Prefix returns s up to but not including position idx.
//
// If idx < 0 it counts from the end of s.
func Prefix[T any, S ~[]T](s S, idx int) S {
	if idx < 0 {
		idx += len(s)
	}
	return s[:idx]
}

// PrefixFunc returns the longest prefix of s whose elements all satisfy the given predicate.
func PrefixFunc[T any, S ~[]T](s S, f func(T) bool) S {
	idx := IndexFunc(s, invert(f))
	if idx < 0 {
		return s
	}
	return s[:idx]
}

// Suffix returns s excluding elements before position idx.
//
// If idx < 0 it counts from the end of s.
func Suffix[T any, S ~[]T](s S, idx int) S {
	if idx < 0 {
		idx += len(s)
	}
	return s[idx:]
}

// Rindex returns the index of the last occurrence of v in s, or -1 if not present.
func Rindex[T comparable, S ~[]T](s S, v T) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == v {
			return i
		}
	}
	return -1
}

// RindexFunc returns the index of the last element in s that satisfies the given predicate,
// or -1 if no such element exists.
func RindexFunc[T any, S ~[]T](s S, f func(T) bool) int {
	for i := len(s) - 1; i >= 0; i-- {
		if f(s[i]) {
			return i
		}
	}
	return -1
}

// SuffixFunc returns the longest suffix of s whose elements all satisfy the given predicate.
func SuffixFunc[T any, S ~[]T](s S, f func(T) bool) S {
	idx := RindexFunc(s, invert(f))
	if idx < 0 {
		return s
	}
	return s[idx+1:]
}

func invert[T any](pred func(T) bool) func(T) bool {
	return func(val T) bool {
		return !pred(val)
	}
}

// SliceN returns n elements of s beginning at position idx.
//
// If idx < 0 it counts from the end of s.
func SliceN[T any, S ~[]T](s S, idx, n int) S {
	if idx < 0 {
		idx += len(s)
	}
	return s[idx : idx+n]
}

// SliceTo returns the elements of s beginning at position from and ending before position to.
//
// If from < 0 it counts from the end of s.
// If to <= 0 it counts from the end of s.
func SliceTo[T any, S ~[]T](s S, from, to int) S {
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

// Each runs a simple function on each item of a slice.
func Each[T any, S ~[]T](s S, f func(T)) {
	_ = Eachx(s, func(_ int, val T) error {
		f(val)
		return nil
	})
}

// Eachx is the extended form of [Each].
// It runs a function on each item of a slice,
// passing the index and the item to the function.
// If any call to the function returns an error,
// Eachx stops looping and exits with the error.
func Eachx[T any, S ~[]T](s S, f func(int, T) error) error {
	for i, val := range s {
		if err := f(i, val); err != nil {
			return err
		}
	}
	return nil
}

// Map runs a simple function on each item of a slice,
// accumulating results in a new slice.
func Map[T, U any, S ~[]T](s S, f func(T) U) []U {
	result, _ := Mapx(s, func(_ int, val T) (U, error) {
		return f(val), nil
	})
	return result
}

// Mapx is the extended form of [Map].
// It runs a function on each item of a slice,
// accumulating results in a new slice.
// If any call to the function returns an error,
// Mapx stops looping and exits with the error.
func Mapx[T, U any, S ~[]T](s S, f func(int, T) (U, error)) ([]U, error) {
	if len(s) == 0 {
		return nil, nil
	}
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

// Accum accumulates the result of repeatedly applying a simple function to the elements of a slice.
//
// If the slice has length 0, the result is the zero value of type T.
// If the slice has length 1, the result is s[0].
// Otherwise, the result is R[len(s)-1],
// where R[0] is s[0]
// and R[n+1] = f(R[n], s[n+1]).
func Accum[T any, S ~[]T](s S, f func(T, T) T) T {
	result, _ := Accumx(s, func(a, b T) (T, error) {
		return f(a, b), nil
	})
	return result
}

// Accumx is the extended form of [Accum].
// It accumulates the result of repeatedly applying a function to the elements of a slice.
//
// If the slice has length 0, the result is the zero value of type T.
// If the slice has length 1, the result is s[0].
// Otherwise, the result is R[len(s)-1],
// where R[0] is s[0]
// and R[n+1] = f(R[n], s[n+1]).
func Accumx[T any, S ~[]T](s S, f func(T, T) (T, error)) (T, error) {
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

// Filter calls a simple predicate for each element of a slice,
// returning a slice of those elements for which the predicate returned true.
func Filter[T any, S ~[]T](s S, f func(T) bool) S {
	result, _ := Filterx(s, func(val T) (bool, error) {
		return f(val), nil
	})
	return result
}

// Filterx is the extended form of [Filter].
// It calls a predicate for each element of a slice,
// returning a slice of those elements for which the predicate returned true.
func Filterx[T any, S ~[]T](s S, f func(T) (bool, error)) (S, error) {
	var result S
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
// It does this by calling a simple grouping function on each element,
// which produces a grouping key.
// The result is a map of group keys to slices of elements having that key.
func Group[T any, K comparable, S ~[]T](s S, f func(T) K) map[K]S {
	result, _ := Groupx(s, func(val T) (K, error) {
		return f(val), nil
	})
	return result
}

// Groupx is the extended form of [Group].
// It partitions the elements of a slice into groups.
// It does this by calling a grouping function on each element,
// which produces a grouping key.
// The result is a map of group keys to slices of elements having that key.
func Groupx[T any, K comparable, S ~[]T](s S, f func(T) (K, error)) (map[K]S, error) {
	result := make(map[K]S)
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
func Rotate[T any, S ~[]T](s S, n int) {
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

// KeyedSort sorts the given slice according to the ordering of the given keys,
// whose items must map 1:1 with the slice.
// It is an unchecked error if keys.Len() != len(slice).
//
// Both arguments end up sorted in place:
// keys according to its Less method,
// and slice by mirroring the reordering that happens in keys.
func KeyedSort[T any, S ~[]T](slice S, keys sort.Interface) {
	ks := keyedSorter[T, S]{
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
type keyedSorter[T any, S ~[]T] struct {
	keys  sort.Interface
	slice S
}

func (k keyedSorter[T, S]) Len() int           { return len(k.slice) }
func (k keyedSorter[T, S]) Less(i, j int) bool { return k.keys.Less(i, j) }
func (k keyedSorter[T, S]) Swap(i, j int) {
	k.keys.Swap(i, j)
	k.slice[i], k.slice[j] = k.slice[j], k.slice[i]
}

// NonNil converts a nil slice to a non-nil empty slice.
// It returns other slices unchanged.
//
// A nil slice is usually preferable,
// since it is equivalent to an empty slice in almost every way
// and does not have the overhead of an allocation.
// (See https://dave.cheney.net/2018/07/12/slices-from-the-ground-up.)
// However, there are some corner cases where the difference matters,
// notably when marshaling to JSON,
// where an empty slice marshals as the array []
// but a nil slice marshals as the non-array `null`.
func NonNil[T any, S ~[]T](s S) S {
	if s == nil {
		return S{}
	}
	return s
}
