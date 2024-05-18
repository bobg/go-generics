package slices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNonNil(t *testing.T) {
	cases := []struct {
		inp  []int
		want []int
	}{{
		inp:  nil,
		want: []int{},
	}, {
		inp:  []int{},
		want: []int{},
	}, {
		inp:  []int{1, 2, 3},
		want: []int{1, 2, 3},
	}}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case_%d", i+1), func(t *testing.T) {
			got := NonNil(c.inp)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("got %v, want %v", got, c.want)
			}
		})
	}
}
