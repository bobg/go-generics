package seqs

import (
	"maps"
	"slices"
	"sort"
	"testing"
)

func TestFrom(t *testing.T) {
	var (
		slice = []int{1, 2, 3}
		seq   = From(slice...)
		got   = slices.Collect(seq)
	)
	if !slices.Equal(got, slice) {
		t.Errorf("got %v, want %v", got, slice)
	}
}

func TestSeq1(t *testing.T) {
	var (
		m    = map[int]int{1: 2, 3: 4}
		seq2 = maps.All(m)
		seq1 = Seq1(seq2)
		got  = slices.Collect(seq1)
		want = []Pair[int, int]{{X: 1, Y: 2}, {X: 3, Y: 4}}
	)
	sort.Slice(got, func(i, j int) bool { return got[i].X < got[j].X })
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestEnumerate(t *testing.T) {
	var (
		slice = []string{"alice", "bob", "charlie"}
		seq   = slices.Values(slice)
		enum  = Enumerate(seq)
		enum1 = Seq1(enum)
		got   = slices.Collect(enum1)
		want  = []Pair[int, string]{{X: 0, Y: "alice"}, {X: 1, Y: "bob"}, {X: 2, Y: "charlie"}}
	)
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSeq2(t *testing.T) {
	var (
		m     = map[int]int{1: 2, 3: 4}
		seq2  = maps.All(m)
		seq1  = Seq1(seq2)
		seq2a = Seq2(seq1)
		got   = maps.Collect(seq2a)
		want  = map[int]int{1: 2, 3: 4}
	)
	if !maps.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestString(t *testing.T) {
	var (
		seq  = String("abc")
		s1   = Seq1(seq)
		got  = slices.Collect(s1)
		want = []Pair[int, rune]{{X: 0, Y: 'a'}, {X: 1, Y: 'b'}, {X: 2, Y: 'c'}}
	)
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
