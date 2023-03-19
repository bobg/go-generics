package parallel

import (
	"sync"
)

// Protect offers safe concurrent access to a protected value.
// It is a "share memory by communicating" alternative
// to protecting the value with [sync.RWMutex].
//
// The caller gets back three functions:
// a reader for getting the protected value,
// a writer for updating it,
// and a closer for releasing resources
// when no further reads or writes are needed.
//
// Any number of calls to the reader may run concurrently.
// If T is a "reference type" (see below)
// then the caller should not make any changes
// to the value it receives from the reader.
//
// A call to the writer prevents other reader and writer calls from running until it is done.
// It waits for pending calls to finish before it executes.
// After a call to the writer,
// future reader calls will receive the updated value.
//
// The closer should be called to release resources
// when no more reader or writer calls are needed.
// Calling any of the functions (reader, writer, or closer)
// after a call to the closer may cause a panic.
//
// The term "reference type" here means a type
// (such as pointer, slice, map, channel, function, and interface)
// that allows a caller C
// to make changes that will be visible to other callers
// outside of C's scope.
// In other words,
// if the type is int and caller A does this:
//
//	val := reader()
//	val++
//
// it will not affect the value that caller B sees when it does its own call to reader().
// But if the type is *int and caller A does this:
//
//	val := reader()
//	*val++
//
// then the change in the pointed-to value _will_ be seen by caller B.
//
// For more on the fuzzy concept of "reference types" in Go,
// see https://github.com/go101/go101/wiki/About-the-terminology-%22reference-type%22-in-Go
func Protect[T any](val T) (reader func() T, writer func(T), closer func()) {
	ch := make(chan rwRequest[T])

	go func() {
		var wg sync.WaitGroup

		for req := range ch {
			req := req // Go loop var pitfall
			if req.r != nil {
				wg.Add(1)
				go func() {
					req.r <- val
					close(req.r)
					wg.Done()
				}()
				continue
			}

			wg.Wait()
			val = <-req.w
		}
	}()

	reader = func() T {
		valch := make(chan T, 1)
		ch <- rwRequest[T]{r: valch}
		return <-valch
	}
	writer = func(val T) {
		valch := make(chan T, 1)
		valch <- val
		ch <- rwRequest[T]{w: valch}
		close(valch)
	}
	closer = func() {
		close(ch)
	}

	return reader, writer, closer
}

type rwRequest[T any] struct {
	r chan<- T
	w <-chan T
}
