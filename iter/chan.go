package iter

import "context"

type chanIter[T any] struct {
	ch    <-chan T
	ctx   context.Context
	latch T
	err   error
}

var _ Of[any] = &chanIter[any]{}

func (ch *chanIter[T]) Next() bool {
	var done <-chan struct{}
	if ch.ctx != nil {
		done = ch.ctx.Done()
	}

	select {
	case <-done:
		ch.err = ch.ctx.Err()
		return false
	case val, ok := <-ch.ch:
		if !ok {
			return false
		}
		ch.latch = val
		return true
	}
}

func (ch *chanIter[T]) Val() T {
	return ch.latch
}

func (ch *chanIter[T]) Err() error {
	return ch.err
}

func FromChan[T any](ch <-chan T) Of[T] {
	return &chanIter[T]{ch: ch}
}

func FromChanContext[T any](ctx context.Context, ch <-chan T) Of[T] {
	return &chanIter[T]{ch: ch, ctx: ctx}
}

func ToChan[T any](inp Of[T]) (<-chan T, func() error) {
	return toChan(nil, inp)
}

func ToChanContext[T any](ctx context.Context, inp Of[T]) (<-chan T, func() error) {
	return toChan(ctx, inp)
}

// ctx can be nil
func toChan[T any](ctx context.Context, inp Of[T]) (<-chan T, func() error) {
	var (
		ch    = make(chan T)
		err   error
		errfn = func() error { return err }
		done  <-chan struct{}
	)

	if ctx != nil {
		done = ctx.Done()
	}

	go func() {
		defer close(ch)

		for inp.Next() {
			select {
			case <-done:
				err = ctx.Err()
				return
			case ch <- inp.Val():
			}
		}
		err = inp.Err()
	}()

	return ch, errfn
}

func Go[T any](ctx context.Context, f func(send func(T) error) error) Of[T] {
	var (
		ch  = make(chan T)
		res = &chanIter[T]{ch: ch, ctx: ctx}
	)
	go func() {
		send := func(val T) error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case ch <- val:
			}
			return nil
		}
		res.err = f(send)
		close(ch)
	}()
	return res
}
