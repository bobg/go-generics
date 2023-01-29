package iter

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestPermutations(t *testing.T) {
	cases := []struct {
		inp  []int
		want [][]int
	}{{
		inp:  nil,
		want: nil,
	}, {
		inp:  []int{1},
		want: [][]int{{1}},
	}, {
		inp:  []int{1, 2},
		want: [][]int{{1, 2}, {2, 1}},
	}, {
		inp:  []int{1, 2, 3},
		want: [][]int{{1, 2, 3}, {2, 1, 3}, {3, 1, 2}, {1, 3, 2}, {2, 3, 1}, {3, 2, 1}},
	}}

	ctx := context.Background()

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case_%02d", i+1), func(t *testing.T) {
			it := Permutations(ctx, tc.inp)
			got, err := ToSlice(it)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestCombinations(t *testing.T) {
	cases := []struct {
		inp  []int
		n    int
		want [][]int
	}{{
		inp:  nil,
		n:    0,
		want: nil,
	}, {
		inp:  []int{1},
		n:    1,
		want: [][]int{{1}},
	}, {
		inp:  []int{1, 2},
		n:    1,
		want: [][]int{{1}, {2}},
	}, {
		inp:  []int{1, 2},
		n:    2,
		want: [][]int{{1, 2}},
	}, {
		inp:  []int{1, 2, 3},
		n:    1,
		want: [][]int{{1}, {2}, {3}},
	}, {
		inp:  []int{1, 2, 3},
		n:    2,
		want: [][]int{{1, 2}, {1, 3}, {2, 3}},
	}, {
		inp:  []int{1, 2, 3},
		n:    3,
		want: [][]int{{1, 2, 3}},
	}}

	ctx := context.Background()

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case_%02d", i+1), func(t *testing.T) {
			it := Combinations(ctx, tc.inp, tc.n)
			got, err := ToSlice(it)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestCombinationsWithReplacement(t *testing.T) {
	cases := []struct {
		inp  []int
		n    int
		want [][]int
	}{{
		inp:  nil,
		n:    0,
		want: nil,
	}, {
		inp:  []int{1},
		n:    1,
		want: [][]int{{1}},
	}, {
		inp:  []int{1, 2},
		n:    1,
		want: [][]int{{1}, {2}},
	}, {
		inp:  []int{1, 2},
		n:    2,
		want: [][]int{{1, 1}, {1, 2}, {2, 2}},
	}, {
		inp:  []int{1, 2, 3},
		n:    1,
		want: [][]int{{1}, {2}, {3}},
	}, {
		inp:  []int{1, 2, 3},
		n:    2,
		want: [][]int{{1, 1}, {1, 2}, {1, 3}, {2, 2}, {2, 3}, {3, 3}},
	}, {
		inp:  []int{1, 2, 3},
		n:    3,
		want: [][]int{{1, 1, 1}, {1, 1, 2}, {1, 1, 3}, {1, 2, 2}, {1, 2, 3}, {1, 3, 3}, {2, 2, 2}, {2, 2, 3}, {2, 3, 3}, {3, 3, 3}},
	}}

	ctx := context.Background()

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case_%02d", i+1), func(t *testing.T) {
			it := CombinationsWithReplacement(ctx, tc.inp, tc.n)
			got, err := ToSlice(it)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
