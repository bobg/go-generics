package set

import (
	"reflect"
	"slices"
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
	s.Each(func(val int) { got[val] = struct{}{} })
	if !reflect.DeepEqual(got, map[int]struct{}{1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}}) {
		t.Errorf("got %v, want [1 2 3 4 5 6]", got)
	}

	s2 := New[int](5, 6, 7, 8)
	i := Intersect(s, s2)
	if !reflect.DeepEqual(i, Of[int](map[int]struct{}{5: {}, 6: {}})) {
		t.Errorf("got %v, want [5 6]", i)
	}
	i = Intersect(s2, nil)
	if i.Len() != 0 {
		t.Errorf("got %v, want []", i)
	}

	u := Union(s, s2, nil)
	if !reflect.DeepEqual(u, Of[int](map[int]struct{}{1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}, 7: {}, 8: {}})) {
		t.Errorf("got %v, want [1 2 3 4 5 6 7 8]", u)
	}

	d := Diff(s, s2)
	if !reflect.DeepEqual(d, Of[int](map[int]struct{}{1: {}, 2: {}, 3: {}, 4: {}})) {
		t.Errorf("got %v, want [1 2 3 4]", d)
	}

	var (
		it = s.All()
		m  = make(map[int]struct{})
	)
	for val := range it {
		m[val] = struct{}{}
	}
	if !reflect.DeepEqual(m, map[int]struct{}{1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}}) {
		t.Errorf("got %v, want [1 2 3 4 5 6]", m)
	}

	m = make(map[int]struct{})
	for _, val := range s.Slice() {
		m[val] = struct{}{}
	}
	if !reflect.DeepEqual(m, map[int]struct{}{1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}}) {
		t.Errorf("got %v, want [1 2 3 4 5 6]", m)
	}
}

func TestEqual(t *testing.T) {
	var (
		a = New[int](1, 2, 3)
		b = New[int](1, 2, 3)
		c = New[int](1, 3, 5)
		d = New[int](1, 5, 9)
	)
	if !a.Equal(b) {
		t.Error("got a != b")
	}
	if !b.Equal(a) {
		t.Error("got b != a")
	}
	if a.Equal(c) {
		t.Error("got a == c")
	}
	if c.Equal(a) {
		t.Error("got c == a")
	}
	if a.Equal(d) {
		t.Error("got a == d")
	}
	if d.Equal(a) {
		t.Error("got d == a")
	}
}

func TestCollect(t *testing.T) {
	var (
		nums = slices.Values([]int{1, 2, 3, 4, 5, 6})
		got  = Collect(nums)
		want = New[int](1, 2, 3, 4, 5, 6)
	)
	if !got.Equal(want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
