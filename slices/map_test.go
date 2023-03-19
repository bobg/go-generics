package slices

import (
	"reflect"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	inp := []int{2, 3, 5}
	want := []string{"2", "3", "5"}
	got := Map(inp, func(val int) string { return strconv.Itoa(val) })
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
