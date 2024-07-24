package seqs

import (
	"reflect"
	"testing"
)

func TestZip(t *testing.T) {
	var (
		inp1 = FromSlice([]int{1, 2, 3})
		inp2 = FromSlice([]string{"a", "b", "c", "d"})
		z    = Zip(inp1, inp2)
		z1   = Seq1(z)
		got  = ToSlice(z1)
		want = []Pair[int, string]{{X: 1, Y: "a"}, {X: 2, Y: "b"}, {X: 3, Y: "c"}, {Y: "d"}}
	)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
