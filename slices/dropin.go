package slices

// TODO: import slices from stdlib after Go 1.21.
import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

// This file contains entrypoints for each of the functions in in golang.org/x/exp/slices
// (except those, like Insert, extended by other functions in this package).

// Equal reports whether two slices are equal: the same length and all
// elements equal. If the lengths are different, Equal returns false.
// Otherwise, the elements are compared in increasing index order, and the
// comparison stops at the first unequal pair.
// Floating point NaNs are not considered equal.
func Equal[E comparable](s1, s2 []E) bool {
	return slices.Equal(s1, s2)
}

// EqualFunc reports whether two slices are equal using a comparison
// function on each pair of elements. If the lengths are different,
// EqualFunc returns false. Otherwise, the elements are compared in
// increasing index order, and the comparison stops at the first index
// for which eq returns false.
func EqualFunc[E1, E2 any](s1 []E1, s2 []E2, eq func(E1, E2) bool) bool {
	return slices.EqualFunc(s1, s2, eq)
}

// Compare compares the elements of s1 and s2.
// The elements are compared sequentially, starting at index 0,
// until one element is not equal to the other.
// The result of comparing the first non-matching elements is returned.
// If both slices are equal until one of them ends, the shorter slice is
// considered less than the longer one.
// The result is 0 if s1 == s2, -1 if s1 < s2, and +1 if s1 > s2.
// Comparisons involving floating point NaNs are ignored.
func Compare[E constraints.Ordered](s1, s2 []E) int {
	return slices.Compare(s1, s2)
}

// CompareFunc is like Compare but uses a comparison function
// on each pair of elements. The elements are compared in increasing
// index order, and the comparisons stop after the first time cmp
// returns non-zero.
// The result is the first non-zero result of cmp; if cmp always
// returns 0 the result is 0 if len(s1) == len(s2), -1 if len(s1) < len(s2),
// and +1 if len(s1) > len(s2).
func CompareFunc[E1, E2 any](s1 []E1, s2 []E2, cmp func(E1, E2) int) int {
	return slices.CompareFunc(s1, s2, cmp)
}

// Index returns the index of the first occurrence of v in s,
// or -1 if not present.
func Index[E comparable](s []E, v E) int {
	return slices.Index(s, v)
}

// IndexFunc returns the first index i satisfying f(s[i]),
// or -1 if none do.
func IndexFunc[E any](s []E, f func(E) bool) int {
	return slices.IndexFunc(s, f)
}

// Contains reports whether v is present in s.
func Contains[E comparable](s []E, v E) bool {
	return slices.Contains(s, v)
}

// ContainsFunc reports whether at least one
// element e of s satisfies f(e).
func ContainsFunc[E any](s []E, f func(E) bool) bool {
	return slices.ContainsFunc(s, f)
}

// Delete removes the elements s[i:j] from s, returning the modified slice.
// Delete panics if s[i:j] is not a valid slice of s.
// Delete modifies the contents of the slice s; it does not create a new slice.
// Delete is O(len(s)-j), so if many items must be deleted, it is better to
// make a single call deleting them all together than to delete one at a time.
// Delete might not modify the elements s[len(s)-(j-i):len(s)]. If those
// elements contain pointers you might consider zeroing those elements so that
// objects they reference can be garbage collected.
//
// If i < 0 it counts from the end of s.
// If j <= 0 it counts from the end of s.
// (This is a change from the behavior of "golang.org/x/exp/slices".Delete.)
func Delete[S ~[]E, E any](s S, i, j int) S {
	return RemoveTo(s, i, j)
}

// Replace replaces the elements s[i:j] by the given v, and returns the
// modified slice. Replace panics if s[i:j] is not a valid slice of s.
//
// If i < 0 it counts from the end of s.
// If j <= 0 it counts from the end of s.
// (This is a change from the behavior of "golang.org/x/exp/slices".Replace.)
func Replace[S ~[]E, E any](s S, i, j int, v ...E) S {
	return ReplaceTo(s, i, j, v...)
}

// Clone returns a copy of the slice.
// The elements are copied using assignment, so this is a shallow clone.
func Clone[S ~[]E, E any](s S) S {
	return slices.Clone(s)
}

// Compact replaces consecutive runs of equal elements with a single copy.
// This is like the uniq command found on Unix.
// Compact modifies the contents of the slice s; it does not create a new slice.
// When Compact discards m elements in total, it might not modify the elements
// s[len(s)-m:len(s)]. If those elements contain pointers you might consider
// zeroing those elements so that objects they reference can be garbage collected.
func Compact[S ~[]E, E comparable](s S) S {
	return slices.Compact(s)
}

// CompactFunc is like Compact but uses a comparison function.
func CompactFunc[S ~[]E, E any](s S, eq func(E, E) bool) S {
	return slices.CompactFunc(s, eq)
}

// Grow increases the slice's capacity, if necessary, to guarantee space for
// another n elements. After Grow(n), at least n elements can be appended
// to the slice without another allocation. If n is negative or too large to
// allocate the memory, Grow panics.
func Grow[S ~[]E, E any](s S, n int) S {
	return slices.Grow(s, n)
}

// Clip removes unused capacity from the slice, returning s[:len(s):len(s)].
func Clip[S ~[]E, E any](s S) S {
	return slices.Clip(s)
}

// Sort sorts a slice of any ordered type in ascending order.
// Sort may fail to sort correctly when sorting slices of floating-point
// numbers containing Not-a-number (NaN) values.
// Use slices.SortFunc(x, func(a, b float64) bool {return a < b || (math.IsNaN(a) && !math.IsNaN(b))})
// instead if the input may contain NaNs.
func Sort[E constraints.Ordered](x []E) {
	slices.Sort(x)
}

// SortFunc sorts the slice x in ascending order as determined by the less function.
// This sort is not guaranteed to be stable.
//
// SortFunc requires that less is a strict weak ordering.
// See https://en.wikipedia.org/wiki/Weak_ordering#Strict_weak_orderings.
func SortFunc[E any](x []E, less func(a, b E) bool) {
	slices.SortFunc(x, less)
}

// SortStableFunc sorts the slice x while keeping the original order of equal
// elements, using less to compare elements.
func SortStableFunc[E any](x []E, less func(a, b E) bool) {
	slices.SortStableFunc(x, less)
}

// IsSorted reports whether x is sorted in ascending order.
func IsSorted[E constraints.Ordered](x []E) bool {
	return slices.IsSorted(x)
}

// IsSortedFunc reports whether x is sorted in ascending order, with less as the
// comparison function.
func IsSortedFunc[E any](x []E, less func(a, b E) bool) bool {
	return slices.IsSortedFunc(x, less)
}

// BinarySearch searches for target in a sorted slice and returns the position
// where target is found, or the position where target would appear in the
// sort order; it also returns a bool saying whether the target is really found
// in the slice. The slice must be sorted in increasing order.
func BinarySearch[E constraints.Ordered](x []E, target E) (int, bool) {
	return slices.BinarySearch(x, target)
}

// BinarySearchFunc works like BinarySearch, but uses a custom comparison
// function. The slice must be sorted in increasing order, where "increasing" is
// defined by cmp. cmp(a, b) is expected to return an integer comparing the two
// parameters: 0 if a == b, a negative number if a < b and a positive number if
// a > b.
func BinarySearchFunc[E, T any](x []E, target T, cmp func(E, T) int) (int, bool) {
	return slices.BinarySearchFunc(x, target, cmp)
}
