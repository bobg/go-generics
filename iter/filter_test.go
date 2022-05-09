package iter

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	ints := Ints(1, 1)
	evens := Filter(ints, func(n int) bool { return n%2 == 0 })
	got, err := ToSlice(FirstN(evens, 3))
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, []int{2, 4, 6}) {
		t.Errorf("got %v, [2 4 6]", got)
	}
}

func TestSkipUntil(t *testing.T) {
	ints := Ints(1, 1)
	first10 := FirstN(ints, 10)
	latter := SkipUntil(first10, func(x int) bool { return x > 7 })
	got, err := ToSlice(latter)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, []int{8, 9, 10}) {
		t.Errorf("got %v, want [8 9 10]", got)
	}
}
