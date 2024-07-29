package seqs

import (
	"slices"
	"testing"
)

func TestLast(t *testing.T) {
	var (
		ints     = Ints(0, 1)
		first100 = FirstN(ints, 100)
		got      = LastN(first100, 10)
		want     = []int{90, 91, 92, 93, 94, 95, 96, 97, 98, 99}
	)
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSkip(t *testing.T) {
	var (
		ints      = Ints(0, 1)
		skipped20 = SkipN(ints, 20)
		twenties  = FirstN(skipped20, 10)
		got       = slices.Collect(twenties)
		want      = []int{20, 21, 22, 23, 24, 25, 26, 27, 28, 29}
	)
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
