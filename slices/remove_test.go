package slices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRemoveN(t *testing.T) {
	cases := []struct {
		inp    []int
		idx, n int
		want   []int
	}{{
		inp:  nil,
		idx:  0,
		n:    0,
		want: nil,
	}, {
		inp:  []int{4, 5, 6, 7},
		idx:  1,
		n:    1,
		want: []int{4, 6, 7},
	}, {
		inp:  []int{4, 5, 6, 7},
		idx:  1,
		n:    2,
		want: []int{4, 7},
	}, {
		inp:  []int{4, 5, 6, 7},
		idx:  -2,
		n:    1,
		want: []int{4, 5, 7},
	}, {
		inp:  []int{4, 5, 6, 7},
		idx:  -2,
		n:    2,
		want: []int{4, 5},
	}}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case_%02d", i+1), func(t *testing.T) {
			got := RemoveN(tc.inp, tc.idx, tc.n)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestRemoveTo(t *testing.T) {
	cases := []struct {
		inp      []int
		from, to int
		want     []int
	}{{
		inp:  nil,
		from: 0,
		to:   0,
		want: nil,
	}, {
		inp:  []int{4, 5, 6, 7},
		from: 1,
		to:   2,
		want: []int{4, 6, 7},
	}, {
		inp:  []int{4, 5, 6, 7},
		from: 1,
		to:   3,
		want: []int{4, 7},
	}, {
		inp:  []int{4, 5, 6, 7},
		from: -2,
		to:   -1,
		want: []int{4, 5, 7},
	}, {
		inp:  []int{4, 5, 6, 7},
		from: -2,
		to:   0,
		want: []int{4, 5},
	}}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case_%02d", i+1), func(t *testing.T) {
			got := RemoveTo(tc.inp, tc.from, tc.to)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
