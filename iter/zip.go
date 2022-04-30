package iter

func Zip[T, U any](t Of[T], u Of[U]) Of[Pair[T, U]] {
	return Gen(func() (Pair[T, U], bool) {
		var (
			x T
			y U
		)

		okx := t.Next()
		oky := u.Next()

		if !okx && !oky {
			return Pair[T, U]{}, false
		}
		if okx {
			x = t.Val()
		}
		if oky {
			y = u.Val()
		}
		return Pair[T, U]{X: x, Y: y}, true
	})
}
