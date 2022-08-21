package maps

import (
	"reflect"
	"sort"
	"testing"

	"github.com/bobg/go-generics/iter"
)

var testMap = map[int]int{1: 2, 3: 6, 7: 14}

func TestDup(t *testing.T) {
	got := Dup(testMap)
	if !reflect.DeepEqual(got, testMap) {
		t.Errorf("got %v, want %v", got, testMap)
	}
}

func TestEach(t *testing.T) {
	var sum int
	err := Each(testMap, func(k, v int) error {
		sum += k * v
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
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

func TestEqual(t *testing.T) {
	inp := Dup(testMap)
	if !Equal(inp, testMap) {
		t.Errorf("Equal says %v and %v are not equal", inp, testMap)
	}
	inp = Invert(inp)
	if Equal(inp, testMap) {
		t.Errorf("Equal says %v and %v are equal", inp, testMap)
	}
}

func TestKeys(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	keys := Keys(m)
	sort.Strings(keys)
	if !reflect.DeepEqual(keys, []string{"a", "b", "c"}) {
		t.Errorf("got %v, want [a b c]", keys)
	}
}
