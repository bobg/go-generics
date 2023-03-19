package iter

import (
	"context"
)

type chanIterConf struct {
	ctx   context.Context
	errfn func() error
}

type chanIter[T any] struct {
	chanIterConf

	ch    <-chan T
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
			if ch.errfn != nil {
				ch.err = ch.errfn()
			}
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

// Option is the type of options that can be passed to FromChan.
type Option func(*chanIterConf)

// WithContext associates a context option with a channel iterator.
func WithContext(ctx context.Context) Option {
	return func(conf *chanIterConf) {
		conf.ctx = ctx
	}
}

// WithError tells a channel iterator how to compute its Err value
// after its channel closes.
func WithError(f func() error) Option {
	return func(conf *chanIterConf) {
		conf.errfn = f
	}
}

// FromChan copies a Go channel to an iterator.
// If the WithContext option is given,
// copying will end early if the given context is canceled
// (and the iterator's Err function will indicate that).
// If the WithError option is given,
// it is called after the channel closes
// to determine the value of the iterator's Err function.
func FromChan[T any](ch <-chan T, opts ...Option) Of[T] {
	res := &chanIter[T]{ch: ch}
	for _, opt := range opts {
		opt(&res.chanIterConf)
	}
	return res
}

// ToChan creates a Go channel and copies the contents of an iterator to it.
// The second return value is a function that may be called
// after the channel is closed
// to inspect any error that occurred.
func ToChan[T any](inp Of[T]) (<-chan T, func() error) {
	//lint:ignore SA1012 nil context is OK here because it is not part of the public API
	return toChan(nil, inp) //nolint:all // nil context is OK here because it is not part of the public API
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

// ctx can be nil!
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
// The function receives a channel for producing values.
// The channel closes when the function exits.
// Any error produced by the function is the value of the iterator's Err method.
func Go[T any](f func(ch chan<- T) error) Of[T] {
	var (
		ch  = make(chan T)
		res = &chanIter[T]{ch: ch}
	)
	go func() {
		res.err = f(ch)
		close(ch)
	}()
	return res
}
