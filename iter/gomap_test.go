package iter

import (
	"reflect"
	"testing"
)

func TestGomaps(t *testing.T) {
	inp := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	iter := FromMap(inp)
	got, err := ToMap(iter)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, inp) {
		t.Errorf("got %v, want [one:1 two:2 three:3]", got)
	}
}
