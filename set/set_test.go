package set

import (
	"reflect"
	"testing"
)

func TestSet(t *testing.T) {
	s := New[int](1, 2, 3)
	s.Add(4, 5, 6)
	if s.Has(0) {
		t.Error("set should not contain 0")
	}
	if !s.Has(1) {
		t.Error("set should contain 1")
	}
	if s.Len() != 6 {
		t.Errorf("got len %d, want 6", s.Len())
	}

	got := make(map[int]struct{})
	err := s.Each(func(val int) error {
		got[val] = struct{}{}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, map[int]struct{}{1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}}) {
		t.Errorf("got %v, want [1 2 3 4 5 6]", got)
	}

	s2 := New[int](5, 6, 7, 8)
	i := Intersect(s, s2)
	if !reflect.DeepEqual(i, Of[int](map[int]struct{}{5: {}, 6: {}})) {
		t.Errorf("got %v, want [5 6]", i)
	}
	u := Union(s, s2)
	if !reflect.DeepEqual(u, Of[int](map[int]struct{}{1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}, 7: {}, 8: {}})) {
		t.Errorf("got %v, want [1 2 3 4 5 6 7 8]", u)
	}
	d := Diff(s, s2)
	if !reflect.DeepEqual(d, Of[int](map[int]struct{}{1: {}, 2: {}, 3: {}, 4: {}})) {
		t.Errorf("got %v, want [1 2 3 4]", d)
	}
}
