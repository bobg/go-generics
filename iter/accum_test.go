package iter

import (
	"reflect"
	"testing"
)

func TestAccum(t *testing.T) {
	inp := FromSlice([]int{1, 2, 3, 4})
	a := Accum(inp, func(a, b int) int { return a + b })
	got := ToSlice(a)
	if !reflect.DeepEqual(got, []int{1, 3, 6, 10}) {
		t.Errorf("got %v, want [1 3 6 10]", got)
	}
}
