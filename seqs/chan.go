package seqs

import (
	"context"
	"iter"
)

func FromChan[T any](ch <-chan T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for val := range ch {
			if !yield(val) {
				break
			}
		}
	}
}

func FromChanContext[T any](ctx context.Context, ch <-chan T) (iter.Seq[T], *error) {
	var err error

	f := func(yield func(T) bool) {
		for {
			select {
			case <-ctx.Done():
				err = ctx.Err()
				return

			case val, ok := <-ch:
				if !ok {
					return
				}
				if !yield(val) {
					return
				}
			}
		}
	}

	return f, &err
}

func ToChan[T any](f iter.Seq[T]) <-chan T {
	ch := make(chan T)

	go func() {
		for val := range f {
			ch <- val
		}
		close(ch)
	}()

	return ch
}

// Go runs a function in a goroutine and returns an iterator over the values it produces.
// The function receives a channel for producing values.
// The channel closes when the function exits.
func Go[F ~func(chan<- T) error, T any](f F) (iter.Seq[T], *error) {
	var (
		ch  = make(chan T)
		err error
	)

	go func() {
		err = f(ch)
		close(ch)
	}()

	return FromChan(ch), &err
}
