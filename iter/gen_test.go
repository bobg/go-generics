package iter

import (
	"reflect"
	"testing"
)

func TestInts(t *testing.T) {
	ints := Ints(1, 2)
	got := ToSlice(FirstN(ints, 10))
	if !reflect.DeepEqual(got, []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}) {
		t.Errorf("got %v, want [1 3 5 7 9 11 13 15 17 19]", got)
	}
}

func TestRepeat(t *testing.T) {
	r := Repeat("foo")
	got := ToSlice(FirstN(r, 10))
	if !reflect.DeepEqual(got, []string{"foo", "foo", "foo", "foo", "foo", "foo", "foo", "foo", "foo", "foo"}) {
		t.Errorf("got %v, want 10 foos", got)
	}
}
