package seqs

import (
	"slices"
	"testing"
)

func TestString(t *testing.T) {
	const s = "こんにちは"

	t.Run("string", func(t *testing.T) {
		var (
			seq   = String(s)
			pairs = ToPairs(seq)
			got   = slices.Collect(pairs)
			want  = []Pair[int, rune]{{X: 0, Y: 'こ'}, {X: 3, Y: 'ん'}, {X: 6, Y: 'に'}, {X: 9, Y: 'ち'}, {X: 12, Y: 'は'}}
		)
		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("runes", func(t *testing.T) {
		var (
			seq  = Runes(s)
			got  = slices.Collect(seq)
			want = []rune{'こ', 'ん', 'に', 'ち', 'は'}
		)
		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("bytes", func(t *testing.T) {
		var (
			seq  = Bytes(s)
			got  = slices.Collect(seq)
			want = []byte{227, 129, 147, 227, 130, 147, 227, 129, 171, 227, 129, 161, 227, 129, 175}
		)
		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
