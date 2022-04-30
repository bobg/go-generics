package iter

func FirstN[T any](inp Of[T], n int) Of[T] {
	return Gen(func() (T, bool) {
		var zero T
		if n <= 0 {
			return zero, false
		}
		if !inp.Next() {
			return zero, false
		}
		n--
		return inp.Val(), true
	})
}

func LastN[T any](inp Of[T], n int) Of[T] {
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
	return Concat(FromSlice(buf[start:]), FromSlice(buf[:start]))
}

func SkipN[T any](inp Of[T], n int) Of[T] {
	first := true

	return Gen(func() (T, bool) {
		var zero T

		if first {
			for i := 0; i < n; i++ {
				if !inp.Next() {
					return zero, false
				}
			}
			first = false
		}

		if !inp.Next() {
			return zero, false
		}
		return inp.Val(), true
	})
}
