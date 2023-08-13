// Package set contains generic typesafe set operations.
package set

import (
	"github.com/bobg/go-generics/v3/iter"
	"github.com/bobg/go-generics/v3/maps"
)

// Of is a set of elements of type T.
// It is called "Of" so that when qualified with this package name
// and instantiated with a member type,
// it reads naturally: e.g., set.Of[int].
//
// The zero value of Of is not safe for use.
// Create one with New instead.
type Of[T comparable] map[T]struct{}

// New produces a new set containing the given values.
func New[T comparable](vals ...T) Of[T] {
	s := Of[T](make(map[T]struct{}))
	for _, val := range vals {
		s.Add(val)
	}
	return Of[T](s)
}

// Add adds the given values to the set.
// Items already present in the set are silently ignored.
func (s Of[T]) Add(vals ...T) {
	for _, val := range vals {
		s[val] = struct{}{}
	}
}

// Has tells whether the given value is in the set.
// The set may be nil.
func (s Of[T]) Has(val T) bool {
	_, ok := s[val]
	return ok
}

// Del removes the given items from the set.
// Items already absent from the set are silently ignored.
func (s Of[T]) Del(vals ...T) {
	for _, val := range vals {
		delete(s, val)
	}
}

// Len tells the number of distinct values in the set.
// The set may be nil.
func (s Of[T]) Len() int {
	return len(s)
}

// Equal tests whether the set has the same membership as another.
// Either set may be nil.
func (s Of[T]) Equal(other Of[T]) bool {
	if len(s) != len(other) {
		return false
	}
	for val := range s {
		if !other.Has(val) {
			return false
		}
	}
	return true
}

func (s Of[T]) Each(f func(T)) {
	_ = s.Eachx(func(val T) error {
		f(val)
		return nil
	})
}

// Eachx calls a function on each element of the set in an indeterminate order.
// It is safe to add and remove items during a call to Each,
// but that can affect the sequence of values seen later during the same Eachx call.
// The set may be nil.
func (s Of[T]) Eachx(f func(T) error) error {
	for val := range s {
		err := f(val)
		if err != nil {
			return err
		}
	}
	return nil
}

// Iter produces an iterator over the members of the set,
// in an indeterminate order.
// The set may be nil.
func (s Of[T]) Iter() iter.Of[T] {
	return iter.FromMapKeys(s)
}

// Slice produces a new slice of the elements in the set.
// The slice is in an indeterminate order.
func (s Of[T]) Slice() []T {
	if s.Len() == 0 {
		return nil
	}
	return maps.Keys(s)
}

// Intersect produces a new set containing only items that appear in all the given sets.
// The input may include nils,
// representing empty sets
// and therefore producing an empty (but non-nil) intersection.
func Intersect[T comparable](sets ...Of[T]) Of[T] {
	result := New[T]()
	if len(sets) == 0 {
		return result
	}
	for _, s := range sets {
		if s == nil {
			return result
		}
	}
	sets[0].Each(func(val T) {
		for _, s := range sets[1:] {
			if !s.Has(val) {
				return
			}
		}
		result.Add(val)
	})
	return result
}

// Union produces a new set containing all the items in all the given sets.
// The input may include nils,
// representing empty sets.
// The result is never nil (but may be empty).
func Union[T comparable](sets ...Of[T]) Of[T] {
	result := New[T]()
	for _, s := range sets {
		if s == nil {
			continue
		}
		s.Each(func(val T) { result.Add(val) })
	}
	return result
}

// Diff produces a new set containing the items in s1 that are not also in s2.
// Either set may be nil.
// The result is never nil (but may be empty).
func Diff[T comparable](s1, s2 Of[T]) Of[T] {
	s := New[T]()
	s1.Each(func(val T) {
		if !s2.Has(val) {
			s.Add(val)
		}
	})
	return s
}
