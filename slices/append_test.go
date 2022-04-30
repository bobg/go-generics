package slices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAppend(t *testing.T) {
	cases := []struct {
		inp, vals, want []int
	}{{
		inp:  nil,
		vals: []int{1},
		want: []int{1},
	}, {
		inp:  nil,
		vals: []int{1, 2},
		want: []int{1, 2},
	}, {
		inp:  []int{1, 2},
		vals: []int{3},
		want: []int{1, 2, 3},
	}, {
		inp:  []int{1, 2},
		vals: []int{3, 4},
		want: []int{1, 2, 3, 4},
	}}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case_%02d", i+1), func(t *testing.T) {
			got := Append(tc.inp, tc.vals...)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
