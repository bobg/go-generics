package slices

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	inp := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := []int{2, 3, 5, 7}
	got, err := Filter(inp, func(n int) (bool, error) {
		for i := 2; i < n; i++ {
			if n%i == 0 {
				return false, nil
			}
		}
		return true, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
