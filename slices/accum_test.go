package slices

import (
	"fmt"
	"testing"
)

func TestAccum(t *testing.T) {
	cases := []struct {
		inp  []int
		want int
	}{{
		inp: nil, want: 0,
	}, {
		inp: []int{1}, want: 1,
	}, {
		inp: []int{1, 2}, want: 3,
	}, {
		inp: []int{1, 2, 3, 4}, want: 10,
	}}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case_%02d", i+1), func(t *testing.T) {
			got, err := Accum(tc.inp, func(a, b int) (int, error) { return a + b, nil })
			if err != nil {
				t.Fatal(err)
			}
			if got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}
