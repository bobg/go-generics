package slices

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	cases := []struct {
		inp  []int
		want []string
	}{{
		inp:  []int{2, 3, 5},
		want: []string{"2", "3", "5"},
	}, {
		inp:  []int{},
		want: nil,
	}}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case_%02d", i+1), func(t *testing.T) {
			got := Map(tc.inp, func(val int) string { return strconv.Itoa(val) })
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
