package slices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReplaceN(t *testing.T) {
	cases := []struct {
		inp    []int
		idx, n int
		vals   []int
		want   []int
	}{{
		inp:  nil,
		idx:  0,
		n:    0,
		vals: []int{4},
		want: []int{4},
	}, {
		inp:  []int{1, 2, 3},
		idx:  0,
		n:    0,
		vals: []int{4},
		want: []int{4, 1, 2, 3},
	}, {
		inp:  []int{1, 2, 3},
		idx:  1,
		n:    0,
		vals: []int{4},
		want: []int{1, 4, 2, 3},
	}, {
		inp:  []int{1, 2, 3},
		idx:  1,
		n:    1,
		vals: []int{4},
		want: []int{1, 4, 3},
	}, {
		inp:  []int{1, 2, 3},
		idx:  -2,
		n:    1,
		vals: []int{4},
		want: []int{1, 4, 3},
	}, {
		inp:  []int{1, 2, 3},
		idx:  1,
		n:    2,
		vals: []int{4},
		want: []int{1, 4},
	}, {
		inp:  []int{1, 2, 3},
		idx:  0,
		n:    3,
		vals: []int{4, 5},
		want: []int{4, 5},
	}, {
		inp:  []int{1, 2, 3},
		idx:  -2,
		n:    2,
		vals: []int{4},
		want: []int{1, 4},
	}}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case_%02d", i+1), func(t *testing.T) {
			got := ReplaceN(tc.inp, tc.idx, tc.n, tc.vals...)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestReplaceTo(t *testing.T) {
	cases := []struct {
		inp      []int
		from, to int
		vals     []int
		want     []int
	}{{
		inp:  []int{1, 2, 3},
		from: 1,
		to:   1,
		vals: []int{4},
		want: []int{1, 4, 2, 3},
	}, {
		inp:  []int{1, 2, 3},
		from: 1,
		to:   2,
		vals: []int{4},
		want: []int{1, 4, 3},
	}, {
		inp:  []int{1, 2, 3},
		from: -2,
		to:   -1,
		vals: []int{4},
		want: []int{1, 4, 3},
	}, {
		inp:  []int{1, 2, 3},
		from: 1,
		to:   0,
		vals: []int{4},
		want: []int{1, 4},
	}, {
		inp:  []int{1, 2, 3},
		from: -2,
		to:   3,
		vals: []int{4},
		want: []int{1, 4},
	}}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case_%02d", i+1), func(t *testing.T) {
			got := ReplaceTo(tc.inp, tc.from, tc.to, tc.vals...)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
