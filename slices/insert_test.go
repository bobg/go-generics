package slices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	cases := []struct {
		base, ins, want []int
		idx             int
	}{{
		base: nil,
		idx:  0,
		ins:  []int{1, 2},
		want: []int{1, 2},
	}, {
		base: []int{1, 2, 3},
		idx:  0,
		ins:  []int{4},
		want: []int{4, 1, 2, 3},
	}, {
		base: []int{1, 2, 3},
		idx:  0,
		ins:  []int{4, 5},
		want: []int{4, 5, 1, 2, 3},
	}, {
		base: []int{1, 2, 3},
		idx:  1,
		ins:  []int{4, 5},
		want: []int{1, 4, 5, 2, 3},
	}, {
		base: []int{1, 2, 3},
		idx:  -1,
		ins:  []int{4, 5},
		want: []int{1, 2, 4, 5, 3},
	}}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case_%02d", i+1), func(t *testing.T) {
			got := Insert(tc.base, tc.idx, tc.ins...)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
