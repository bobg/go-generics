package seqs

import (
	"bufio"
	"errors"
	"io"
	"iter"
)

// Lines produces an iterator over the lines of text in r.
// This uses a bufio.Scanner
// and is subject to its default line-length limit
// (see https://pkg.go.dev/bufio#pkg-constants).
//
// The caller can dereference the returned error pointer to check for errors
// but only after iteration is done.
func Lines(r io.Reader) (iter.Seq[string], *error) {
	var err error

	f := func(yield func(string) bool) {
		sc := bufio.NewScanner(r)
		for sc.Scan() {
			if !yield(sc.Text()) {
				return
			}
		}
		err = sc.Err()
	}

	return f, &err
}

// LongLines produces an iterator of readers,
// each delivering a single line of text from r.
// Unlike [Lines],
// this does not use a [bufio.Scanner]
// and is not subject to its default line-length limit.
// Each reader must be fully consumed before the next one is available.
//
// The caller can dereference the returned error pointer to check for errors
// but only after iteration is done.
func LongLines(r io.Reader) (iter.Seq[io.Reader], *error) {
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

		newpipe := func() {
			pr, pw = io.Pipe()
			ch <- pr
		}

		newline := func() error {
			if pw == nil {
				newpipe()
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
				newpipe()
			}
			_, err := pw.Write([]byte{b})
			return err
		}

		newpipe()

		for {
			b, err := br.ReadByte()
			if errors.Is(err, io.EOF) {
				err = nil
				if pw != nil {
					err = pw.Close()
				}
				return err
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
