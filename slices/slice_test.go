package slices

import (
	"fmt"
	"reflect"
	"slices"
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

func isEven(n int) bool { return n%2 == 0 }

func TestPrefixFunc(t *testing.T) {
	cases := []struct {
		inp, want []int
	}{{
		inp:  nil,
		want: nil,
	}, {
		inp:  []int{4},
		want: []int{4},
	}, {
		inp:  []int{5},
		want: nil,
	}, {
		inp:  []int{4, 5, 6, 7},
		want: []int{4},
	}, {
		inp:  []int{4, 6},
		want: []int{4, 6},
	}, {
		inp:  []int{4, 6, 5, 7},
		want: []int{4, 6},
	}, {
		inp:  []int{7, 5, 6, 4},
		want: nil,
	}}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case_%02d", i+1), func(t *testing.T) {
			got := PrefixFunc(tc.inp, isEven)
			if !slices.Equal(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestSuffixFunc(t *testing.T) {
	cases := []struct {
		inp, want []int
	}{{
		inp:  nil,
		want: nil,
	}, {
		inp:  []int{4},
		want: []int{4},
	}, {
		inp:  []int{5},
		want: nil,
	}, {
		inp:  []int{4, 5, 6, 7},
		want: nil,
	}, {
		inp:  []int{4, 6},
		want: []int{4, 6},
	}, {
		inp:  []int{7, 6, 5, 4},
		want: []int{4},
	}}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case_%02d", i+1), func(t *testing.T) {
			got := SuffixFunc(tc.inp, isEven)
			if !slices.Equal(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestRindex(t *testing.T) {
	cases := []struct {
		inp  []int
		want int
	}{{
		inp:  nil,
		want: -1,
	}, {
		inp:  []int{4, 5, 6, 7},
		want: 1,
	}, {
		inp:  []int{4, 5, 6, 7, 5},
		want: 4,
	}, {
		inp:  []int{6, 7, 8},
		want: -1,
	}, {
		inp:  []int{6, 5, 4, 3},
		want: 1,
	}, {
		inp:  []int{5, 5, 5, 4},
		want: 2,
	}}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case_%02d", i+1), func(t *testing.T) {
			got := Rindex(tc.inp, 5)
			if got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
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

func TestRotate(t *testing.T) {
	cases := []struct {
		inp  []int
		n    int
		want []int
	}{{
		inp:  []int{4, 5, 1, 2, 3},
		n:    -2,
		want: []int{1, 2, 3, 4, 5},
	}, {
		inp:  []int{4, 5, 1, 2, 3},
		n:    -7,
		want: []int{1, 2, 3, 4, 5},
	}, {
		inp:  []int{4, 5, 1, 2, 3},
		n:    3,
		want: []int{1, 2, 3, 4, 5},
	}, {
		inp:  []int{4, 5, 1, 2, 3},
		n:    8,
		want: []int{1, 2, 3, 4, 5},
	}}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case_%02d", i+1), func(t *testing.T) {
			got := make([]int, len(tc.inp))
			copy(got, tc.inp)
			Rotate(got, tc.n)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
