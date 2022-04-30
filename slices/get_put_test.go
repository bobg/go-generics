package slices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	cases := []struct {
		inp  []int
		idx  int
		want int
	}{{
		inp:  []int{4, 5, 6},
		idx:  0,
		want: 4,
	}, {
		inp:  []int{4, 5, 6},
		idx:  -1,
		want: 6,
	}}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case_%02d", i+1), func(t *testing.T) {
			got := Get(tc.inp, tc.idx)
			if got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}

func TestPut(t *testing.T) {
	cases := []struct {
		inp      []int
		idx, val int
		want     []int
	}{{
		inp:  []int{4, 5, 6},
		idx:  0,
		val:  7,
		want: []int{7, 5, 6},
	}, {
		inp:  []int{4, 5, 6},
		idx:  -1,
		val:  7,
		want: []int{4, 5, 7},
	}}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case_%02d", i+1), func(t *testing.T) {
			Put(tc.inp, tc.idx, tc.val)
			if !reflect.DeepEqual(tc.inp, tc.want) {
				t.Errorf("got %v, want %v", tc.inp, tc.want)
			}
		})
	}
}
