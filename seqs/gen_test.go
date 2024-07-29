package seqs

import (
	"slices"
	"testing"
)

func TestInts(t *testing.T) {
	var (
		ints    = Ints(1, 2)
		first10 = FirstN(ints, 10)
		got     = slices.Collect(first10)
		want    = []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	)
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestRepeat(t *testing.T) {
	var (
		r       = Repeat("foo")
		first10 = FirstN(r, 10)
		got     = slices.Collect(first10)
		want    = []string{"foo", "foo", "foo", "foo", "foo", "foo", "foo", "foo", "foo", "foo"}
	)
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
