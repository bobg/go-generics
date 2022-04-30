// Package iter defines an iterator interface,
// a collection of concrete iterator types,
// and some functions for operating on iterators.
package iter

// Of is the interface implemented by iterators.
// It is called "Of" so that when qualified with this package name
// and instantiated with a member type,
// it reads naturally: e.g., iter.Of[int].
type Of[T any] interface {
	// Next advances the iterator to its next value and tells whether one is available to read.
	// A true result is necessary before calling Val.
	// Once Next returns false, it must continue returning false.
	Next() bool

	// Val returns the current value of the iterator.
	// Callers must get a true result from Next before calling Val.
	// Repeated calls to Val
	// with no intervening call to Next
	// should return the same value.
	Val() T
}
