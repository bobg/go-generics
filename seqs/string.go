package seqs

import "iter"

// String produces an [iter.Seq2] over position-rune pairs in a string.
// The position of each rune is measured in bytes from the beginning of the string.
func String(inp string) iter.Seq2[int, rune] {
	return func(yield func(int, rune) bool) {
		for i, r := range inp {
			if !yield(i, r) {
				return
			}
		}
	}
}

// Bytes returns an iterator over the bytes in a string.
func Bytes(inp string) iter.Seq[byte] {
	return func(yield func(byte) bool) {
		for i := 0; i < len(inp); i++ {
			if !yield(inp[i]) {
				return
			}
		}
	}
}

// Runes returns an iterator over the runes in a string.
// This is the same as Right(String(inp)).
func Runes(inp string) iter.Seq[rune] {
	return func(yield func(rune) bool) {
		for _, r := range inp {
			if !yield(r) {
				return
			}
		}
	}
}
