package slices

import (
	"reflect"
	"testing"
)

func TestGroup(t *testing.T) {
	inp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := map[string][]int{
		"evens": {2, 4, 6, 8, 10},
		"odds":  {1, 3, 5, 7, 9},
	}
	got := Group(inp, func(n int) string {
		if n%2 == 0 {
			return "evens"
		}
		return "odds"
	})
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
