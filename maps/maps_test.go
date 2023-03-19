package maps

import (
	"reflect"
	"testing"

	"github.com/bobg/go-generics/v2/iter"
)

var (
	testMap     = map[int]int{1: 2, 3: 6, 7: 14}
	testMapDups = map[int]int{1: 2, 3: 2, 7: 14}
)

func TestEach(t *testing.T) {
	var sum int
	Each(testMap, func(k, v int) {
		sum += k * v
	})
	if sum != 118 {
		t.Errorf("got %d, want 118", sum)
	}
}

func TestFromPairs(t *testing.T) {
	inp := []iter.Pair[int, int]{{
		X: 1, Y: 2,
	}, {
		X: 3, Y: 6,
	}, {
		X: 7, Y: 14,
	}}
	got, err := FromPairs(iter.FromSlice(inp))
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, testMap) {
		t.Errorf("got %v, want %v", got, testMap)
	}
}

func TestInvert(t *testing.T) {
	var (
		want = map[int]int{2: 1, 6: 3, 14: 7}
		got  = Invert(testMap)
	)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestInvertMulti(t *testing.T) {
	var (
		// Both are possible. The order of keys in the slices is undefined.
		want1 = map[int][]int{2: {1, 3}, 14: {7}}
		want2 = map[int][]int{2: {3, 1}, 14: {7}}

		got = InvertMulti(testMapDups)
	)

	if !reflect.DeepEqual(got, want1) && !reflect.DeepEqual(got, want2) {
		t.Errorf("got %v, want %v", got, want1)
	}
}
