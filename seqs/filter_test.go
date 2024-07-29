package seqs

import (
	"slices"
	"testing"
)

func TestFilter(t *testing.T) {
	var (
		ints   = Ints(1, 1)
		evens  = Filter(ints, func(n int) bool { return n%2 == 0 })
		first3 = FirstN(evens, 3)
		got    = slices.Collect(first3)
		want   = []int{2, 4, 6}
	)
	if !slices.Equal(got, want) {
		t.Errorf("got %v, %v", got, want)
	}
}

func TestSkipUntil(t *testing.T) {
	var (
		ints    = Ints(1, 1)
		first10 = FirstN(ints, 10)
		latter  = SkipUntil(first10, func(x int) bool { return x > 7 })
		got     = slices.Collect(latter)
		want    = []int{8, 9, 10}
	)
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
