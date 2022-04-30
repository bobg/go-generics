package iter

// Slice is an iterator over a slice.
type sliceIter[T any] struct {
	s   []T
	idx int
}

var _ Of[any] = &sliceIter[any]{}

// Next implements Of.Next.
func (s *sliceIter[T]) Next() bool {
	if s.idx >= len(s.s) {
		return false
	}
	s.idx++
	return true
}

// Val implements Of.Val.
func (s *sliceIter[T]) Val() T {
	return s.s[s.idx-1]
}

// FromSlice creates an iterator over the elements of a slice.
func FromSlice[T any](s []T) Of[T] {
	return &sliceIter[T]{s: s}
}

// From creates an iterator over the given items.
func From[T any](items ...T) Of[T] {
	return FromSlice(items)
}

// ToSlice consumes the elements of an iterator and returns them as a slice.
// Be careful! The input may be very long or even infinite.
// Consider using FirstN to ensure the input has a reasonable size.
func ToSlice[T any](iter Of[T]) []T {
	var result []T
	for iter.Next() {
		result = append(result, iter.Val())
	}
	return result
}
