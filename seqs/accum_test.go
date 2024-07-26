package seqs

import (
	"fmt"
	"slices"
	"testing"
)

func TestAccum(t *testing.T) {
	cases := []struct {
		inp  []int
		want []int
	}{{}, {
		inp:  []int{1},
		want: []int{1},
	}, {
		inp:  []int{1, 2},
		want: []int{1, 3},
	}, {
		inp:  []int{1, 2, 3},
		want: []int{1, 3, 6},
	}}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case_%02d", i+1), func(t *testing.T) {
			var (
				inp = slices.Values(tc.inp)
				a   = Accum(inp, func(a, b int) int { return a + b })
				got = slices.Collect(a)
			)
			if !slices.Equal(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
