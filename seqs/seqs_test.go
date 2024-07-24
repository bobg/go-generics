package seqs

import (
	"maps"
	"slices"
	"testing"
)

func TestSeq1(t *testing.T) {
	var (
		m    = map[int]int{1: 2, 3: 4}
		seq2 = maps.All(m)
		seq1 = Seq1(seq2)
		got  = slices.Collect(seq1)
		want = []Pair[int, int]{{X: 1, Y: 2}, {X: 3, Y: 4}}
	)
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

func TestChanSeq(t *testing.T) {
	ch := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch)
	}()

	var (
		seq  = ChanSeq(ch)
		got  = slices.Collect(seq)
		want = []int{0, 1, 2}
	)
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestStringSeq(t *testing.T) {
	var (
		seq  = StringSeq("abc")
		s1   = Seq1(seq)
		got  = slices.Collect(s1)
		want = []Pair[int, rune]{{X: 0, Y: 'a'}, {X: 1, Y: 'b'}, {X: 2, Y: 'c'}}
	)
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
