package iter

func Concat[T any](inps ...Of[T]) Of[T] {
	return &concatIter[T]{inps: inps}
}

type concatIter[T any] struct {
	inps []Of[T]
	err  error
}

func (c *concatIter[T]) Next() bool {
	for len(c.inps) > 0 {
		if c.inps[0].Next() {
			return true
		}
		if err := c.inps[0].Err(); err != nil {
			c.err = err
			return false
		}
		c.inps = c.inps[1:]
	}
	return false
}

func (c *concatIter[T]) Val() T {
	return c.inps[0].Val()
}

func (c *concatIter[T]) Err() error {
	return c.err
}
