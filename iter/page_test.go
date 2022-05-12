package iter

import (
	"reflect"
	"testing"
)

func TestPage(t *testing.T) {
	var (
		ints    = Ints(1, 1)
		first10 = FirstN(ints, 10)
		got     [][]int
	)
	err := Page(first10, 3, func(page []int, final bool) error {
		dup := make([]int, len(page))
		copy(dup, page)
		got = append(got, dup)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	want := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
