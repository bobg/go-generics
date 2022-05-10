package iter

// Zip takes two iterators and produces a new iterator containing pairs of corresponding elements.
// If one input iterator ends before the other,
// Zip produces zero values of the appropriate type in constructing pairs.
func Zip[T, U any](t Of[T], u Of[U]) Of[Pair[T, U]] {
	return Gen(func() (Pair[T, U], bool, error) {
		var (
			x   T
			y   U
			err error
		)

		okx := t.Next()
		if !okx {
			err = t.Err()
		}

		oky := u.Next()
		if !oky && err == nil {
			err = u.Err()
		}

		if (!okx && !oky) || err != nil {
			return Pair[T, U]{}, false, err
		}
		if okx {
			x = t.Val()
		}
		if oky {
			y = u.Val()
		}
		return Pair[T, U]{X: x, Y: y}, true, nil
	})
}
