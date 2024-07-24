package seqs

import (
	"slices"
	"testing"
)

func TestAccum(t *testing.T) {
	var (
		inp  = slices.Values([]int{1, 2, 3, 4})
		a    = Accum(inp, func(a, b int) int { return a + b })
		got  = slices.Collect(a)
		want = []int{1, 3, 6, 10}
	)
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want [1 3 6 10]", got)
	}
}
