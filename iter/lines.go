package iter

import (
	"bufio"
	"context"
	"errors"
	"io"
)

// Lines produces an iterator over the lines of text in r.
// This uses a bufio.Scanner
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

		newpipe := func() error {
			pr, pw = io.Pipe()
			select {
			case <-ctx.Done():
				return ctx.Err()
			case ch <- pr:
			}
			return nil
		}

		newline := func() error {
			if pw == nil {
				if err := newpipe(); err != nil {
					return err
				}
			}

			if err := pw.Close(); err != nil {
				return err
			}

			pw = nil
			sawCR = false

			return nil
		}

		write := func(b byte) error {
			if pw == nil {
				if err := newpipe(); err != nil {
					return err
				}
			}
			_, err := pw.Write([]byte{b})
			return err
		}

		if err := newpipe(); err != nil {
			return err
		}

		for {
			b, err := br.ReadByte()
			if errors.Is(err, io.EOF) {
				if pw != nil {
					return pw.Close()
				}
				return nil
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
				if err := write('\r'); err != nil {
					return err
				}
				sawCR = false
			}
			if b == '\r' {
				sawCR = true
				continue
			}

			if err := write(b); err != nil {
				return err
			}
		}
	})
}
