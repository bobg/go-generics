package slices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPrefix(t *testing.T) {
	cases := []struct {
		inp  []int
		idx  int
		want []int
	}{{
		inp:  []int{4, 5, 6},
		idx:  0,
		want: []int{},
	}, {
		inp:  []int{4, 5, 6},
		idx:  1,
		want: []int{4},
	}, {
		inp:  []int{4, 5, 6},
		idx:  2,
		want: []int{4, 5},
	}, {
		inp:  []int{4, 5, 6},
		idx:  -1,
		want: []int{4, 5},
	}, {
		inp:  []int{4, 5, 6},
		idx:  -2,
		want: []int{4},
	}}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case_%02d", i+1), func(t *testing.T) {
			got := Prefix(tc.inp, tc.idx)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestSuffix(t *testing.T) {
	cases := []struct {
		inp  []int
		idx  int
		want []int
	}{{
		inp:  []int{4, 5, 6},
		idx:  0,
		want: []int{4, 5, 6},
	}, {
		inp:  []int{4, 5, 6},
		idx:  1,
		want: []int{5, 6},
	}, {
		inp:  []int{4, 5, 6},
		idx:  2,
		want: []int{6},
	}, {
		inp:  []int{4, 5, 6},
		idx:  -1,
		want: []int{6},
	}, {
		inp:  []int{4, 5, 6},
		idx:  -2,
		want: []int{5, 6},
	}}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case_%02d", i+1), func(t *testing.T) {
			got := Suffix(tc.inp, tc.idx)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestSliceN(t *testing.T) {
	cases := []struct {
		inp    []int
		idx, n int
		want   []int
	}{{
		inp:  []int{4, 5, 6, 7},
		idx:  0,
		n:    0,
		want: []int{},
	}, {
		inp:  []int{4, 5, 6, 7},
		idx:  0,
		n:    1,
		want: []int{4},
	}, {
		inp:  []int{4, 5, 6, 7},
		idx:  1,
		n:    1,
		want: []int{5},
	}, {
		inp:  []int{4, 5, 6, 7},
		idx:  0,
		n:    2,
		want: []int{4, 5},
	}, {
		inp:  []int{4, 5, 6, 7},
		idx:  1,
		n:    2,
		want: []int{5, 6},
	}, {
		inp:  []int{4, 5, 6, 7},
		idx:  -1,
		n:    1,
		want: []int{7},
	}, {
		inp:  []int{4, 5, 6, 7},
		idx:  -2,
		n:    1,
		want: []int{6},
	}, {
		inp:  []int{4, 5, 6, 7},
		idx:  -2,
		n:    2,
		want: []int{6, 7},
	}}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case_%02d", i+1), func(t *testing.T) {
			got := SliceN(tc.inp, tc.idx, tc.n)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestSliceTo(t *testing.T) {
	cases := []struct {
		inp      []int
		from, to int
		want     []int
	}{{
		inp:  []int{4, 5, 6, 7},
		from: 0,
		to:   0,
		want: []int{4, 5, 6, 7},
	}, {
		inp:  []int{4, 5, 6, 7},
		from: 0,
		to:   1,
		want: []int{4},
	}, {
		inp:  []int{4, 5, 6, 7},
		from: 1,
		to:   2,
		want: []int{5},
	}, {
		inp:  []int{4, 5, 6, 7},
		from: 0,
		to:   2,
		want: []int{4, 5},
	}, {
		inp:  []int{4, 5, 6, 7},
		from: 1,
		to:   3,
		want: []int{5, 6},
	}, {
		inp:  []int{4, 5, 6, 7},
		from: -1,
		to:   0,
		want: []int{7},
	}, {
		inp:  []int{4, 5, 6, 7},
		from: -2,
		to:   -1,
		want: []int{6},
	}, {
		inp:  []int{4, 5, 6, 7},
		from: -2,
		to:   0,
		want: []int{6, 7},
	}}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case_%02d", i+1), func(t *testing.T) {
			got := SliceTo(tc.inp, tc.from, tc.to)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
