package iter

// sliceIter is an iterator over a slice.
type sliceIter[S ~[]T, T any] struct {
	s   S
	idx int
}

var _ Of[any] = &sliceIter[[]any, any]{}

// Next implements Of.Next.
func (s *sliceIter[S, T]) Next() bool {
	if s.idx >= len(s.s) {
		return false
	}
	s.idx++
	return true
}

// Val implements Of.Val.
func (s *sliceIter[S, T]) Val() T {
	return s.s[s.idx-1]
}

func (*sliceIter[S, T]) Err() error {
	return nil
}

// FromSlice creates an iterator over the elements of a slice.
func FromSlice[S ~[]T, T any](s S) Of[T] {
	return &sliceIter[S, T]{s: s}
}

// From creates an iterator over the given items.
func From[T any](items ...T) Of[T] {
	return FromSlice(items)
}

// ToSlice consumes the elements of an iterator and returns them as a slice.
// Be careful! The input may be very long or even infinite.
// Consider using FirstN to ensure the input has a reasonable size.
func ToSlice[T any](iter Of[T]) ([]T, error) {
	var result []T
	for iter.Next() {
		result = append(result, iter.Val())
	}
	return result, iter.Err()
}
