package slices

import (
	"errors"
	"reflect"
	"testing"
)

func TestEach(t *testing.T) {
	inp := []int{2, 3, 5}

	type wanttype struct {
		idx, val int
	}
	want := []wanttype{{
		idx: 0, val: 2,
	}, {
		idx: 1, val: 3,
	}, {
		idx: 2, val: 5,
	}}

	var got []wanttype
	err := Each(inp, func(idx, val int) error {
		got = append(got, wanttype{idx: idx, val: val})
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}

	e := errors.New("error")

	err = Each(inp, func(_, _ int) error {
		return e
	})
	if !errors.Is(err, e) {
		t.Errorf("got %v, want error %v", err, e)
	}
}
