package slices

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	inp := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := []int{2, 3, 5, 7}
	got := Filter(inp, func(n int) bool {
		for i := 2; i < n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	})
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
