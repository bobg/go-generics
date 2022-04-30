// Package set contains generic typesafe set operations.
package set

import "github.com/bobg/go-generics/iter"

// Of is a set of elements of type T.
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
func (s Of[T]) Len() int {
	return len(s)
}

// Each calls a function on each element of the set in an indeterminate order.
// It is safe to add and remove items during a call to Each,
// but that can affect the sequence of values seen later during the same Each call.
func (s Of[T]) Each(f func(T) error) error {
	for val := range s {
		err := f(val)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s Of[T]) Iter() iter.Of[T] {
	return iter.Map(iter.FromMap(s), func(pair iter.Pair[T, struct{}]) T { return pair.X })
}

func (s Of[T]) Slice() []T {
	return iter.ToSlice(s.Iter())
}

// Intersect produces a new set containing only items that appear in all the given sets.
func Intersect[T comparable](sets ...Of[T]) Of[T] {
	s := New[T]()
	if len(sets) == 0 {
		return s
	}
	sets[0].Each(func(val T) error {
		for _, other := range sets[1:] {
			if !other.Has(val) {
				return nil
			}
		}
		s.Add(val)
		return nil
	})
	return s
}

// Union produces a new set containing all the items in all the given sets.
func Union[T comparable](sets ...Of[T]) Of[T] {
	result := New[T]()
	for _, s := range sets {
		s.Each(func(val T) error {
			result.Add(val)
			return nil
		})
	}
	return result
}

// Diff produces a new set containing the items in s1 that are not also in s2.
func Diff[T comparable](s1, s2 Of[T]) Of[T] {
	s := New[T]()
	s1.Each(func(val T) error {
		if !s2.Has(val) {
			s.Add(val)
		}
		return nil
	})
	return s
}
