package internal

import (
	"fmt"
	"slices"
	"testing"
)

func TestRotateSlice(t *testing.T) {
	cases := []struct {
		slice []int
		rot   int
		want  []int
	}{{
		slice: []int{1, 2, 3},
		rot:   0,
		want:  []int{1, 2, 3},
	}, {
		slice: []int{1, 2, 3},
		rot:   1,
		want:  []int{3, 1, 2},
	}, {
		slice: []int{1, 2, 3},
		rot:   2,
		want:  []int{2, 3, 1},
	}, {
		slice: []int{1, 2, 3},
		rot:   3,
		want:  []int{1, 2, 3},
	}, {
		slice: []int{1, 2, 3},
		rot:   4,
		want:  []int{3, 1, 2},
	}, {
		slice: []int{1, 2, 3},
		rot:   -1,
		want:  []int{2, 3, 1},
	}, {
		slice: []int{1, 2, 3},
		rot:   -2,
		want:  []int{3, 1, 2},
	}, {
		slice: []int{1, 2, 3},
		rot:   -3,
		want:  []int{1, 2, 3},
	}, {
		slice: []int{1, 2, 3},
		rot:   -4,
		want:  []int{2, 3, 1},
	}}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case_%02d", i+1), func(t *testing.T) {
			slice := append([]int(nil), tc.slice...)
			RotateSlice(slice, tc.rot)
			if !slices.Equal(slice, tc.want) {
				t.Fatalf("got %v, want %v", slice, tc.want)
			}
		})
	}
}
