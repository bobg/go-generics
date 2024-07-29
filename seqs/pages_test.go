package seqs

import (
	"reflect"
	"slices"
	"testing"
)

func TestPage(t *testing.T) {
	var (
		ints    = Ints(1, 1)
		first10 = FirstN(ints, 10)
		pages   = Pages(first10, 3)
		got     = slices.Collect(pages)
		want    = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10}}
	)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
