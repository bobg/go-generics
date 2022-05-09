package iter

import "github.com/bobg/go-generics/slices"

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
	slices.Rotate(buf, -start)
	return buf, nil
}

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
