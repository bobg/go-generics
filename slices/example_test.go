package slices_test

import (
	"fmt"
	"sort"

	"github.com/bobg/go-generics/slices"
)

func ExampleGet() {
	s := []int{1, 2, 3, 4}
	last := slices.Get(s, -1)
	fmt.Println(last)
	// Output: 4
}

func ExampleInsert() {
	s1 := []int{10, 15, 16}
	s2 := slices.Insert(s1, 1, 11, 12, 13, 14)
	fmt.Println(s2)
	// Output: [10 11 12 13 14 15 16]
}

func ExampleReplaceN() {
	s1 := []int{99, 0, 0, 0, 97}
	s2 := slices.ReplaceN(s1, 1, 3, 98)
	fmt.Println(s2)
	// Output: [99 98 97]
}

func ExampleReplaceTo() {
	s1 := []int{99, 0, 0, 0, 97}
	s2 := slices.ReplaceTo(s1, 1, -1, 98)
	fmt.Println(s2)
	// Output: [99 98 97]
}

func ExampleRemoveN() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := slices.RemoveN(s1, -2, 2)
	fmt.Println(s2)
	// Output: [1 2 3]
}

func ExampleRemoveTo() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := slices.RemoveTo(s1, -2, 0)
	fmt.Println(s2)
	// Output: [1 2 3]
}

func ExampleEach() {
	s := []int{100, 200, 300}
	slices.Each(s, func(idx, val int) error {
		fmt.Println(idx, val)
		return nil
	})
	// Output:
	// 0 100
	// 1 200
	// 2 300
}

func ExampleMap() {
	s1 := []int{1, 2, 3, 4, 5}
	s2, err := slices.Map(s1, func(idx, val int) (string, error) {
		return string([]byte{byte('a' + val - 1)}), nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(s2)
	// Output: [a b c d e]
}

func ExampleAccum() {
	s := []int{1, 2, 3, 4, 5}
	sum, err := slices.Accum(s, func(a, b int) (int, error) { return a + b, nil })
	if err != nil {
		panic(err)
	}
	fmt.Println(sum)
	// Output: 15
}

func ExampleFilter() {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	evens, err := slices.Filter(s, func(val int) (bool, error) { return val%2 == 0, nil })
	if err != nil {
		panic(err)
	}
	fmt.Println(evens)
	// Output: [2 4 6]
}

func ExampleGroup() {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	groups, err := slices.Group(s, func(val int) (string, error) {
		if val%2 == 0 {
			return "even", nil
		}
		return "odd", nil
	})
	if err != nil {
		panic(err)
	}

	for key, slice := range groups {
		fmt.Println(key, slice)
	}
	// Unordered output:
	// even [2 4 6]
	// odd [1 3 5 7]
}

func ExampleRotate() {
	s := []int{3, 4, 5, 1, 2}
	slices.Rotate(s, 2)
	fmt.Println(s)
	// Output: [1 2 3 4 5]
}

func ExampleKeyedSort() {
	var (
		nums  = []int{1, 2, 3, 4, 5}
		names = []string{"one", "two", "three", "four", "five"}
	)

	// Sort the numbers in `nums` according to their names in `names`.
	slices.KeyedSort(nums, sort.StringSlice(names))

	fmt.Println(nums)
	// Output: [5 4 1 3 2]
}
