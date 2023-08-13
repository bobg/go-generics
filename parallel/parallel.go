// Package parallel contains generic typesafe functions to manage concurrent logic of various kinds.
package parallel

import (
	"context"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"

	"github.com/bobg/go-generics/v3/iter"
)

// Error is an error type for wrapping errors returned from worker goroutines.
// It contains the worker number of the goroutine that produced the error.
type Error struct {
	N   int
	Err error
}

func (e Error) Error() string {
	return fmt.Sprintf("in goroutine %d: %s", e.N, e.Err)
}

func (e Error) Unwrap() error {
	return e.Err
}

// Values produces a slice of n values using n parallel workers each running the function f.
//
// Each worker receives its worker number (in the range 0 through n-1).
//
// An error from any worker cancels them all.
// The first error is returned to the caller.
//
// The resulting slice has length n.
// The value at position i comes from worker i.
func Values[F ~func(context.Context, int) (T, error), T any](ctx context.Context, n int, f F) ([]T, error) {
	g, ctx := errgroup.WithContext(ctx)
	result := make([]T, n)

	for i := 0; i < n; i++ {
		i := i // Go loop var pitfall
		g.Go(func() error {
			val, err := f(ctx, i)
			result[i] = val
			if err != nil {
				return Error{N: i, Err: err}
			}
			return nil
		})
	}

	err := g.Wait()
	return result, err
}

// Producers launches n parallel workers each running the function f.
//
// Each worker receives its worker number
// (in the range 0 through n-1)
// and a callback to use for producing a value.
// If the callback returns an error,
// the worker should exit with that error.
//
// The callback that the worker uses to produce a value may block
// until the caller is able to consume the value.
//
// An error from any worker cancels them all.
//
// The caller gets an iterator over the values produced.
func Producers[F ~func(context.Context, int, func(T) error) error, T any](ctx context.Context, n int, f F) iter.Of[T] {
	ch := make(chan T)
	g, innerCtx := errgroup.WithContext(ctx)

	for i := 0; i < n; i++ {
		i := i
		g.Go(func() error {
			return f(innerCtx, i, func(val T) error {
				select {
				case <-innerCtx.Done():
					return Error{N: i, Err: innerCtx.Err()}
				case ch <- val:
					return nil
				}
			})
		})
	}

	var err error

	go func() {
		err = g.Wait()
		close(ch)
	}()

	return iter.FromChan(ch, iter.WithContext(ctx), iter.WithError(func() error { return err }))
}

// Consumers launches n parallel workers each consuming values supplied by the caller.
//
// When a value is available,
// an available worker calls the function f to consume it.
// This callback receives the worker's number
// (in the range 0 through n-1)
// and the value.
//
// The caller receives two callbacks:
// one for sending a value to the workers via an internal channel,
// and one for closing that channel,
// signaling the end of input and causing the workers to exit normally.
//
// The value-sending callback may block until a worker is available to consume the value.
//
// An error from any worker cancels them all.
// This error is returned from the close-channel callback.
// After any error, the value-sending callback will return an error.
// (Not the original error, however.
// For that, the caller should still invoke the close callback.)
func Consumers[F ~func(context.Context, int, T) error, T any](ctx context.Context, n int, f F) (func(T) error, func() error) {
	ch := make(chan T, n)

	g, ctx := errgroup.WithContext(ctx)

	for i := 0; i < n; i++ {
		i := i
		g.Go(func() error {
			for {
				select {
				case <-ctx.Done():
					return Error{N: i, Err: ctx.Err()}
				case val, ok := <-ch:
					if !ok {
						return nil
					}
					err := f(ctx, i, val)
					if err != nil {
						return Error{N: i, Err: err}
					}
				}
			}
		})
	}

	sendfn := func(val T) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case ch <- val:
			return nil
		}
	}

	closefn := func() error {
		close(ch)
		return g.Wait()
	}

	return sendfn, closefn
}

// Pool permits up to n concurrent calls to a function f.
// The caller receives a callback for requesting a worker from this pool.
// When no worker is available,
// the callback blocks until one becomes available.
// Then it invokes f and returns the result.
//
// Each call of the callback is synchronous.
// Any desired concurrency is the responsibility of the caller.
func Pool[F ~func(T) (U, error), T, U any](n int, f F) func(T) (U, error) {
	var (
		running int
		mu      sync.Mutex
		cond    = sync.NewCond(&mu)
	)
	return func(val T) (U, error) {
		mu.Lock()
		for running >= n {
			cond.Wait()
		}
		running++
		mu.Unlock()

		result, err := f(val)

		mu.Lock()
		running--
		cond.Signal()
		mu.Unlock()

		return result, err
	}
}
