package iter

import (
	"reflect"
	"testing"
)

func TestAll(t *testing.T) {
	var (
		ints   = Ints(1, 1)      // All integers starting at 1
		first5 = FirstN(ints, 5) // First 5 integers
	)

	testSeq(t, All(first5), []int{1, 2, 3, 4, 5})
}

func TestAllCount(t *testing.T) {
	names := FromSlice([]string{"Alice", "Bob", "Carol"})

	testSeq2(t, AllCount(names), []int{0, 1, 2}, []string{"Alice", "Bob", "Carol"})
}

func TestAllPairs(t *testing.T) {
	var (
		letters = FromSlice([]string{"a", "b", "c", "d"})
		nums    = FromSlice([]int{1, 2, 3})
		pairs   = Zip(letters, nums)
	)

	testSeq2(t, AllPairs(pairs), []string{"a", "b", "c", "d"}, []int{1, 2, 3, 0})
}

func testSeq[T any](t *testing.T, seq Seq[T], want []T) {
	var got []T

	seq(func(val T) bool {
		got = append(got, val)
		return true
	})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func testSeq2[T, U any](t *testing.T, seq Seq2[T, U], wantT []T, wantU []U) {
	var (
		gotT []T
		gotU []U
	)

	seq(func(valT T, valU U) bool {
		gotT = append(gotT, valT)
		gotU = append(gotU, valU)
		return true
	})

	if !reflect.DeepEqual(gotT, wantT) {
		t.Errorf("got %v, want %v", gotT, wantT)
	}

	if !reflect.DeepEqual(gotU, wantU) {
		t.Errorf("got %v, want %v", gotU, wantU)
	}
}

func TestFromSeq(t *testing.T) {
	seq := func(yield func(int) bool) {
		for i := 0; i < 5; i++ {
			if !yield(i) {
				break
			}
		}
	}
	it := FromSeq(seq)
	got, err := ToSlice(it)
	if err != nil {
		t.Fatal(err)
	}
	want := []int{0, 1, 2, 3, 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
