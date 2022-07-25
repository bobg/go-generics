package maps

import (
	"reflect"
	"testing"
)

func TestInvert(t *testing.T) {
	var (
		inp  = map[int]int{1: 2, 3: 6, 7: 14}
		want = map[int]int{2: 1, 6: 3, 14: 7}
	)
	got := Invert(inp)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
