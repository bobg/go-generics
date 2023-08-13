package iter

import "github.com/bobg/go-generics/v3/internal"

// FirstN produces an iterator containing the first n elements of the input
// (or all of the input, if there are fewer than n elements).
// Remaining elements of the input are not consumed.
// It is the caller's responsibility to release any associated resources.
func FirstN[T any](inp Of[T], n int) Of[T] {
	return Gen(func() (T, bool, error) {
		var zero T
		if n <= 0 {
			return zero, false, nil
		}
		if !inp.Next() {
			return zero, false, inp.Err()
		}
		n--
		return inp.Val(), true, nil
	})
}

// LastN produces a slice containing the last n elements of the input iterator
// (or all of the input, if there are fewer than n elements).
// There is no guarantee that any elements will ever be produced:
// the input iterator may be infinite!
func LastN[T any](inp Of[T], n int) ([]T, error) {
	var (
		buf   = make([]T, 0, n)
		start = 0
	)
	for inp.Next() {
		val := inp.Val()
		if len(buf) < n {
			buf = append(buf, val)
			continue
		}
		buf[start] = val
		start = (start + 1) % n
	}
	if err := inp.Err(); err != nil {
		return nil, err
	}
	internal.RotateSlice(buf, -start)
	return buf, nil
}

// SkipN copies the input iterator to the output,
// skipping the first N elements.
func SkipN[T any](inp Of[T], n int) Of[T] {
	first := true

	return Gen(func() (T, bool, error) {
		var zero T

		if first {
			for i := 0; i < n; i++ {
				if !inp.Next() {
					return zero, false, inp.Err()
				}
			}
			first = false
		}

		if !inp.Next() {
			return zero, false, inp.Err()
		}
		return inp.Val(), true, nil
	})
}
