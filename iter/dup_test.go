package iter

import (
	"reflect"
	"testing"
)

func TestDup(t *testing.T) {
	inp := FromSlice([]int{1, 2, 3})
	dups := Dup(inp, 2)
	s1 := ToSlice(dups[0])
	if !reflect.DeepEqual(s1, []int{1, 2, 3}) {
		t.Errorf("got %v, want [1 2 3]", s1)
	}
	s2 := ToSlice(dups[1])
	if !reflect.DeepEqual(s2, []int{1, 2, 3}) {
		t.Errorf("got %v, want [1 2 3]", s2)
	}
}
