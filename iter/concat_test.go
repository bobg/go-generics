package iter

import (
	"reflect"
	"testing"
)

func TestConcat(t *testing.T) {
	c := Concat(
		FromSlice([]int{1, 2, 3}),
		FromSlice([]int{4, 5, 6}),
	)
	got := ToSlice(c)
	if !reflect.DeepEqual(got, []int{1, 2, 3, 4, 5, 6}) {
		t.Errorf("got %v, want [1 2 3 4 5 6]", got)
	}
}
