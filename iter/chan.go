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

// FromChan copies a Go channel to an iterator.
func FromChan[T any](ch <-chan T) Of[T] {
	return &chanIter[T]{ch: ch}
}

// FromChanContext copies a Go channel to an iterator.
// Copying will end early if the context is canceled
// (and the iterator's Err() function will indicate that).
func FromChanContext[T any](ctx context.Context, ch <-chan T) Of[T] {
	return &chanIter[T]{ch: ch, ctx: ctx}
}

// ToChan creates a Go channel and copies the contents of an iterator to it.
// The second return value is a function that may be called
// after the channel is closed
// to inspect any error that occurred.
func ToChan[T any](inp Of[T]) (<-chan T, func() error) {
	return toChan(nil, inp)
}

// ToChanContext creates a Go channel and copies the contents of an iterator to it.
// The second return value is a function that may be called
// after the channel is closed
// to inspect any error that occurred.
// The channel will close early if the context is canceled
// (and the error-returning function will indicate that).
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

// Go runs a function in a goroutine and returns an iterator over the values it produces.
// The function receives a callback for producing values.
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
