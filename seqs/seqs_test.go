package seqs

import (
	"maps"
	"slices"
	"sort"
	"testing"
)

func TestEmpty(t *testing.T) {
	got := slices.Collect(Empty[int])
	if len(got) != 0 {
		t.Errorf("got %v, want []", got)
	}

	got2 := slices.Collect(ToPairs(Empty2[int, int]))
	if len(got2) != 0 {
		t.Errorf("got %v, want []", got2)
	}
}

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

func TestToPairs(t *testing.T) {
	var (
		m    = map[int]int{1: 2, 3: 4}
		seq2 = maps.All(m)
		seq1 = ToPairs(seq2)
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
		enum1 = ToPairs(enum)
		got   = slices.Collect(enum1)
		want  = []Pair[int, string]{{X: 0, Y: "alice"}, {X: 1, Y: "bob"}, {X: 2, Y: "charlie"}}
	)
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFromPairs(t *testing.T) {
	var (
		m     = map[int]int{1: 2, 3: 4}
		seq2  = maps.All(m)
		seq1  = ToPairs(seq2)
		seq2a = FromPairs(seq1)
		got   = maps.Collect(seq2a)
		want  = map[int]int{1: 2, 3: 4}
	)
	if !maps.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestLeftRight(t *testing.T) {
	var (
		left   = []int{1, 2, 3, 4, 5}
		right  = []int{6, 7, 8, 9, 10}
		zipped = Zip(slices.Values(left), slices.Values(right))
	)
	t.Run("left", func(t *testing.T) {
		got := slices.Collect(Left(zipped))
		if !slices.Equal(got, left) {
			t.Errorf("got %v, want %v", got, left)
		}
	})
	t.Run("right", func(t *testing.T) {
		got := slices.Collect(Right(zipped))
		if !slices.Equal(got, right) {
			t.Errorf("got %v, want %v", got, right)
		}
	})
}

func TestCompare(t *testing.T) {
	var (
		a = []int{1, 2, 3}
		b = []int{1, 2, 4}
		c = []int{1, 2}
		d = []int{1, 2, 3, 4}
	)
	t.Run("equal", func(t *testing.T) {
		if got := Compare(slices.Values(a), slices.Values(a)); got != 0 {
			t.Errorf("got %d, want 0", got)
		}
	})
	t.Run("shorter", func(t *testing.T) {
		if got := Compare(slices.Values(a), slices.Values(d)); got >= 0 {
			t.Errorf("got %d, want < 0", got)
		}
	})
	t.Run("longer", func(t *testing.T) {
		if got := Compare(slices.Values(a), slices.Values(c)); got <= 0 {
			t.Errorf("got %d, want > 0", got)
		}
	})
	t.Run("less", func(t *testing.T) {
		if got := Compare(slices.Values(a), slices.Values(b)); got >= 0 {
			t.Errorf("got %d, want < 0", got)
		}
	})
	t.Run("greater", func(t *testing.T) {
		if got := Compare(slices.Values(b), slices.Values(a)); got <= 0 {
			t.Errorf("got %d, want > 0", got)
		}
	})
}
