package slices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want []int
	}{{
		in:   []int{1, 2, 3, 4, 5},
		want: []int{5, 4, 3, 2, 1},
	}, {
		in:   []int{1, 2, 3, 4},
		want: []int{4, 3, 2, 1},
	}, {
		in:   []int{1, 2, 3},
		want: []int{3, 2, 1},
	}, {
		in:   []int{1, 2},
		want: []int{2, 1},
	}, {
		in:   []int{1},
		want: []int{1},
	}, {
		in:   nil,
		want: nil,
	}}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case_%02d", i+1), func(t *testing.T) {
			Reverse(tc.in)
			if !reflect.DeepEqual(tc.in, tc.want) {
				t.Errorf("got %v, want %v", tc.in, tc.want)
			}
		})
	}
}
