package iter

import (
	"reflect"
	"testing"
)

func TestZip(t *testing.T) {
	inp1 := FromSlice([]int{1, 2, 3})
	inp2 := FromSlice([]string{"a", "b", "c", "d"})
	z := Zip(inp1, inp2)
	got := ToSlice(z)
	if !reflect.DeepEqual(got, []Pair[int, string]{{X: 1, Y: "a"}, {X: 2, Y: "b"}, {X: 3, Y: "c"}, {Y: "d"}}) {
		t.Errorf("got %v, want [[1 a] [2 b] [3 c] [0 d]]", got)
	}
}
