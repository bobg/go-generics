package parallel

import (
	"context"
	"sync"
)

type (
	// RWReader is the type of a callback for reading a value protected by RW.
	RWReader[T any] func(T)

	// RWWriter is the type of a callback for writing a value protected by RW.
	RWWriter[T any] func(T) T
)

type rwRequest[T any] struct {
	r    func(T)
	w    func(T) T
	done chan struct{}
}

// RW offers read and write access to a protected value.
// It is a "share memory by communicating" alternative
// to protecting the value with [sync.RWMutex].
//
// The caller gets back two functions:
// one to use for reading the protected value,
// and one for writing it.
// In each case, the caller supplies a callback that receives the protected value.
// The return value of the writer callback is used as the new protected value.
//
// Any number of read calls may run concurrently.
// If T is a pointer type,
// read calls should not make any changes in the pointed-to data.
//
// A write call prevents any other read and write calls from running until it is done.
// It allows pending calls to finish before executing.
//
// When no more read or write calls are needed,
// the caller should cancel the context to release resources.
//
// Calling the read or write function after the context has been canceled may result in a panic.
func RW[T any](ctx context.Context, protected T) (reader func(func(T)), writer func(func(T) T)) {
	ch := make(chan rwRequest[T])

	go func() {
		defer close(ch)

		var wg sync.WaitGroup

		for {
			select {
			case <-ctx.Done():
				wg.Wait()
				return

			case req := <-ch:
				if req.r != nil {
					wg.Add(1)
					go func() {
						req.r(protected)
						wg.Done()
						close(req.done)
					}()
					continue
				}

				wg.Wait()
				protected = req.w(protected)
				close(req.done)
			}
		}
	}()

	return rwReaderCallback(ctx, ch), rwWriterCallback(ctx, ch)
}

func rwReaderCallback[T any](ctx context.Context, ch chan<- rwRequest[T]) func(func(T)) {
	return func(r func(T)) {
		done := make(chan struct{})
		select {
		case <-ctx.Done():
			return
		case ch <- rwRequest[T]{r: r, done: done}:
			select {
			case <-ctx.Done():
			case <-done:
			}
			return
		}
	}
}

func rwWriterCallback[T any](ctx context.Context, ch chan<- rwRequest[T]) func(func(T) T) {
	return func(w func(T) T) {
		done := make(chan struct{})
		select {
		case <-ctx.Done():
			return
		case ch <- rwRequest[T]{w: w, done: done}:
			select {
			case <-ctx.Done():
			case <-done:
			}
			return
		}
	}
}
