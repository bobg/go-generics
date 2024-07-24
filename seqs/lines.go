package seqs

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"iter"
)

// Lines produces an iterator over the lines of text in r.
// This uses a bufio.Scanner
// and is subject to its default line-length limit
// (see https://pkg.go.dev/bufio#pkg-constants).
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
// If not consuming all readers from the iterator,
// the caller should cancel the context to reclaim resources.
func LongLines(r io.Reader) (iter.Seq[io.Reader], *error) {
	br, ok := r.(io.ByteReader)
	if !ok {
		br = bufio.NewReader(r)
	}

	var err error

	f := func(yield func(io.Reader) bool) {
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

		newpipe := func() (bool, error) {
			pr, pw = io.Pipe()

			fmt.Printf("xxx calling yield\n")
			if !yield(pr) {
				fmt.Printf("xxx yield returned false\n")
				return false, nil
			}
			fmt.Printf("xxx yield returned true\n")
			return true, nil
		}

		newline := func() error {
			if pw == nil {
				if ok, err := newpipe(); err != nil || !ok {
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
				if ok, err := newpipe(); err != nil || !ok {
					return err
				}
			}
			_, err := pw.Write([]byte{b})
			return err
		}

		var ok bool
		if ok, err = newpipe(); err != nil || !ok {
			return
		}

		for {
			var b byte

			b, err = br.ReadByte()
			if errors.Is(err, io.EOF) {
				err = nil
				if pw != nil {
					err = pw.Close()
				}
				return
			}
			if err != nil {
				return
			}

			if b == '\n' {
				if err = newline(); err != nil {
					return
				}
				continue
			}

			if sawCR {
				if err = write('\r'); err != nil {
					return
				}
				sawCR = false
			}
			if b == '\r' {
				sawCR = true
				continue
			}

			if err = write(b); err != nil {
				return
			}
		}
	}

	return f, &err
}
