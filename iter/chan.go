package iter

import "context"

type chanIter[T any] struct {
	ch    <-chan T
	done  <-chan struct{}
	latch T
}

var _ Of[any] = &chanIter[any]{}

func (ch *chanIter[T]) Next() bool {
	select {
	case <-ch.done:
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

func FromChan[T any](ch <-chan T) Of[T] {
	return &chanIter[T]{ch: ch}
}

func FromChanContext[T any](ctx context.Context, ch <-chan T) Of[T] {
	return &chanIter[T]{ch: ch, done: ctx.Done()}
}

func ToChan[T any](inp Of[T]) <-chan T {
	return toChan(inp, nil)
}

func ToChanContext[T any](ctx context.Context, inp Of[T]) <-chan T {
	return toChan(inp, ctx.Done())
}

func toChan[T any](inp Of[T], done <-chan struct{}) <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)

		for inp.Next() {
			select {
			case <-done:
				return
			case ch <- inp.Val():
			}
		}
	}()
	return ch
}
