package seqs

import (
	"context"
	"iter"
)

// FromChan produces an [iter.Seq] over the contents of a channel.
func FromChan[T any](inp <-chan T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for x := range inp {
			if !yield(x) {
				return
			}
		}
	}
}

// FromChanContext produces an [iter.Seq] over the contents of a channel.
// It stops at the end of the channel or when the given context is canceled.
//
// The caller can dereference the returned error pointer to check for errors
// (such as context cancellation),
// but only after iteration is done.
func FromChanContext[T any](ctx context.Context, inp <-chan T) (iter.Seq[T], *error) {
	var err error

	f := func(yield func(T) bool) {
		for {
			select {
			case val, ok := <-inp:
				if !ok {
					return
				}
				if !yield(val) {
					return
				}

			case <-ctx.Done():
				err = ctx.Err()
				return
			}
		}
	}

	return f, &err
}

// ToChan launches a goroutine that consumes an iterator and sends it values to a channel.
func ToChan[T any](inp iter.Seq[T]) <-chan T {
	ch := make(chan T)

	go func() {
		for val := range inp {
			ch <- val
		}
		close(ch)
	}()

	return ch
}

// ToChanContext launches a goroutine that consumes an iterator and sends it values to a channel.
// It stops early when the context is canceled.
//
// The caller can dereference the returned error pointer to check for errors
// (such as context cancellation),
// but only after reaching the end of the channel.
func ToChanContext[T any](ctx context.Context, f iter.Seq[T]) (<-chan T, *error) {
	var (
		ch  = make(chan T)
		err error
	)

	go func() {
		defer close(ch)

		for val := range f {
			// This extra check helps to ensure that context cancellation "wins" when both cases in the select can proceed.
			if err = ctx.Err(); err != nil {
				return
			}

			select {
			case ch <- val:
				// OK, do nothing.

			case <-ctx.Done():
				err = ctx.Err()
				return
			}
		}
	}()

	return ch, &err
}

// Go runs a function in a goroutine and returns an iterator over the values it produces.
// The function receives a channel for producing values.
// The channel closes when the function exits.
func Go[T any, F ~func(chan<- T) error](f F) (iter.Seq[T], *error) {
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
