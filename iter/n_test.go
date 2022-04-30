package iter

import (
	"reflect"
	"testing"
)

func TestLast(t *testing.T) {
	ints := Ints(0, 1)
	first100 := FirstN(ints, 100)
	nineties := LastN(first100, 10)
	got := ToSlice(nineties)
	if !reflect.DeepEqual(got, []int{90, 91, 92, 93, 94, 95, 96, 97, 98, 99}) {
		t.Errorf("got %v, want [90 91 92 93 94 95 96 97 98 99]", got)
	}
}

func TestSkip(t *testing.T) {
	ints := Ints(0, 1)
	twenties := FirstN(SkipN(ints, 20), 10)
	got := ToSlice(twenties)
	if !reflect.DeepEqual(got, []int{20, 21, 22, 23, 24, 25, 26, 27, 28, 29}) {
		t.Errorf("got %v, want [20 21 22 23 24 25 26 27 28 29]", got)
	}
}
