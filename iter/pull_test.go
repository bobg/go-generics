//go:build go1.23 || goexperiment.rangefunc

package iter

import (
	"reflect"
	"testing"
)

func TestPull(t *testing.T) {
	var (
		ints       = Ints(1, 1) // All integers starting at 1
		next, stop = Pull(All(ints))
		want       = []int{1, 2, 3, 4, 5}
		got        []int
	)

	for range 5 {
		v, ok := next()
		if !ok {
			break
		}
		got = append(got, v)
	}
	stop()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("have %v, want %v", got, want)
	}

	v, ok := next()
	if v != 0 || ok != false {
		t.Errorf("next() after stop() gives %d, %v, want 0, false", v, ok)
	}
}

func TestPull2(t *testing.T) {
	var (
		names      = FromSlice([]string{"Alice", "Bob", "Carol", "Dave"})
		namelens   = FromSlice([]int{5, 3, 5, 4})
		pairs      = Zip(names, namelens)
		next, stop = Pull2(AllPairs(pairs))
		want       = []any{
			"Alice", 5,
			"Bob", 3,
			"Carol", 5,
			"Dave", 4,
			"", 0,
		}
		got []any
	)

	for {
		name, namelen, ok := next()
		got = append(got, name, namelen)
		if !ok {
			break
		}
	}
	stop()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
