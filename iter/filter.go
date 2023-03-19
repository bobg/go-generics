package iter

// Filter copies the input iterator to the output,
// including only those elements that cause f to return true.
func Filter[F ~func(T) bool, T any](inp Of[T], f F) Of[T] {
	return &filterIter[T]{inp: inp, f: f}
}

type filterIter[T any] struct {
	inp Of[T]
	f   func(T) bool
	val T
}

func (f *filterIter[T]) Next() bool {
	for f.inp.Next() {
		f.val = f.inp.Val()
		if f.f(f.val) {
			return true
		}
	}
	return false
}

func (f *filterIter[T]) Val() T {
	return f.val
}

func (f *filterIter[T]) Err() error {
	return f.inp.Err()
}

// SkipUntil copies the input iterator to the output,
// discarding the initial elements until the first one that causes f to return true.
// That element and the remaining elements of inp are included in the output,
// and f is not called again.
func SkipUntil[F ~func(T) bool, T any](inp Of[T], f F) Of[T] {
	skipping := true
	return Filter(inp, func(val T) bool {
		if !skipping {
			return true
		}
		if f(val) {
			skipping = false
		}
		return !skipping
	})
}
