package seqs

import "iter"

type Pair[T, U any] struct {
	X T
	Y U
}

// Seq1 changes an [iter.Seq2] to an [iter.Seq] of [Pair]s.
func Seq1[T, U any](inp iter.Seq2[T, U]) iter.Seq[Pair[T, U]] {
	return func(yield func(Pair[T, U]) bool) {
		for x, y := range inp {
			if !yield(Pair[T, U]{X: x, Y: y}) {
				return
			}
		}
	}
}

// Enumerate changes an [iter.Seq] to an [iter.Seq2] of (index, val) pairs.
func Enumerate[T any](inp iter.Seq[T]) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		i := 0
		for x := range inp {
			if !yield(i, x) {
				return
			}
			i++
		}
	}
}

// Seq2 changes an [iter.Seq] of [Pair]s to an [iter.Seq2].
func Seq2[T, U any](inp iter.Seq[Pair[T, U]]) iter.Seq2[T, U] {
	return func(yield func(T, U) bool) {
		for val := range inp {
			if !yield(val.X, val.Y) {
				return
			}
		}
	}
}

func ChanSeq[T any](inp <-chan T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for x := range inp {
			if !yield(x) {
				return
			}
		}
	}
}

func StringSeq(inp string) iter.Seq2[int, rune] {
	return func(yield func(int, rune) bool) {
		for i, r := range inp {
			if !yield(i, r) {
				return
			}
		}
	}
}

func Empty[T any](func(T) bool)        {}
func Empty2[T, U any](func(T, U) bool) {}
