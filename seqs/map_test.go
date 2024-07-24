package seqs

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	var (
		inp = FromSlice([]int{1, 2, 3, 4})
		m   = Map(inp, func(x int) int { return x * x })
		s   = ToSlice(m)
	)
	if !reflect.DeepEqual(s, []int{1, 4, 9, 16}) {
		t.Errorf("got %v, want [1, 4, 9, 16]", s)
	}
}
