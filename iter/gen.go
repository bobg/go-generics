package iter

import (
	"bufio"
	"context"
	"errors"
	"io"
)

// Gen produces an iterator of values obtained by repeatedly calling f.
// If f returns an error,
// iteration stops and the error is available via the iterator's Err method.
// Otherwise, each call to f should return a value and a true boolean.
// When f returns a false boolean, it signals the normal end of iteration.
func Gen[F ~func() (T, bool, error), T any](f F) Of[T] {
	return &genIter[T]{f: f}
}

type genIter[T any] struct {
	f   func() (T, bool, error)
	val T
	end bool
	err error
}

func (g *genIter[T]) Next() bool {
	if g.end {
		return false
	}
	val, ok, err := g.f()
	if err != nil || !ok {
		g.end = true
		g.err = err
		return false
	}
	g.val = val
	return true
}

func (g *genIter[T]) Val() T {
	return g.val
}

func (g *genIter[T]) Err() error {
	return g.err
}

// Ints produces an infinite iterator over integers beginning at start,
// with each element increasing by delta.
func Ints(start, delta int) Of[int] {
	n := start
	return Gen(func() (int, bool, error) {
		res := n
		n += delta
		return res, true, nil
	})
}

// Repeat produces an infinite iterator repeatedly containing the given value.
func Repeat[T any](val T) Of[T] {
	return Gen(func() (T, bool, error) { return val, true, nil })
}

// Lines produces an iterator over the lines of text in r.
// This uses a [bufio.Scanner]
// and is subject to its default line-length limit
// (see https://pkg.go.dev/bufio#pkg-constants).
func Lines(r io.Reader) Of[string] {
	sc := bufio.NewScanner(r)
	return Gen(func() (string, bool, error) {
		if sc.Scan() {
			return sc.Text(), true, nil
		}
		return "", false, sc.Err()
	})
}

// LongLines produces an iterator of readers,
// each delivering a single line of text from r.
// Unlike [Lines],
// this does not use a [bufio.Scanner]
// and is not subject to its default line-length limit.
// Each reader must be fully consumed before the next one is available.
// If not consuming all readers from the iterator,
// the caller should cancel the context to reclaim resources.
func LongLines(ctx context.Context, r io.Reader) Of[io.Reader] {
	br, ok := r.(io.ByteReader)
	if !ok {
		br = bufio.NewReader(r)
	}

	return Go(func(ch chan<- io.Reader) error {
		var (
			pr    io.Reader
			pw    io.WriteCloser
			sawCR bool
		)

		// Defer closing only the final value of pw.
		defer func() {
			if pw != nil {
				pw.Close()
			}
		}()

		newline := func() error {
			if pw != nil {
				if err := pw.Close(); err != nil {
					return err
				}
			}

			pr, pw = io.Pipe()

			select {
			case <-ctx.Done():
				return ctx.Err()
			case ch <- pr:
			}

			sawCR = false

			return nil
		}

		if err := newline(); err != nil {
			return err
		}

		for {
			b, err := br.ReadByte()
			if errors.Is(err, io.EOF) {
				return pw.Close()
			}
			if err != nil {
				return err
			}

			if b == '\n' {
				if err := newline(); err != nil {
					return err
				}
				continue
			}

			if sawCR {
				if _, err := pw.Write([]byte{'\r'}); err != nil {
					return err
				}
				sawCR = false
			}
			if b == '\r' {
				sawCR = true
				continue
			}

			if _, err := pw.Write([]byte{b}); err != nil {
				return err
			}
		}
	})
}
