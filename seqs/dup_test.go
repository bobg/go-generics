package seqs

import (
	"slices"
	"testing"
)

func TestDup(t *testing.T) {
	var (
		slice = []int{1, 2, 3}
		it    = slices.Values(slice)
		dups  = Dup(it, 2)
		s1    = slices.Collect(dups[0])
	)
	if !slices.Equal(s1, slice) {
		t.Errorf("got %v, want %v", s1, slice)
	}
	s2 := slices.Collect(dups[1])
	if !slices.Equal(s2, slice) {
		t.Errorf("got %v, want %v", s2, slice)
	}
}
