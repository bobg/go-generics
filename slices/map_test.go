package slices

import (
	"reflect"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	inp := []int{2, 3, 5}
	want := []string{"2", "3", "5"}
	got, err := Map(inp, func(idx, val int) (string, error) {
		return strconv.Itoa(val), nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
