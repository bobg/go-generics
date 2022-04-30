package iter

func Concat[T any](inps ...Of[T]) Of[T] {
	return &concatIter[T]{inps: inps}
}

type concatIter[T any] struct {
	inps []Of[T]
}

func (c *concatIter[T]) Next() bool {
	for len(c.inps) > 0 {
		if c.inps[0].Next() {
			return true
		}
		c.inps = c.inps[1:]
	}
	return false
}

func (c *concatIter[T]) Val() T {
	return c.inps[0].Val()
}
