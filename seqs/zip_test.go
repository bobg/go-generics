package seqs

import (
	"slices"
	"testing"
)

func TestZip(t *testing.T) {
	var (
		inp1 = slices.Values([]int{1, 2, 3})
		inp2 = slices.Values([]string{"a", "b", "c", "d"})
		z    = Zip(inp1, inp2)
		z1   = ToPairs(z)
		got  = slices.Collect(z1)
		want = []Pair[int, string]{{X: 1, Y: "a"}, {X: 2, Y: "b"}, {X: 3, Y: "c"}, {Y: "d"}}
	)
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
