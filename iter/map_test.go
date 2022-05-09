package iter

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	inp := FromSlice([]int{1, 2, 3, 4})
	m := Map(inp, func(x int) (int, error) { return x * x, nil })
	s, err := ToSlice(m)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(s, []int{1, 4, 9, 16}) {
		t.Errorf("got %v, want [1, 4, 9, 16]", s)
	}
}
