package seqs

import (
	"slices"
	"testing"
)

func TestConcat(t *testing.T) {
	var (
		a    = slices.Values([]int{1, 2, 3})
		b    = slices.Values([]int{4, 5, 6})
		c    = Concat(a, b)
		got  = slices.Collect(c)
		want = []int{1, 2, 3, 4, 5, 6}
	)
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
