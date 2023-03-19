package slices_test

import (
	"fmt"
	"sort"

	"github.com/bobg/go-generics/v2/slices"
)

func ExampleGet() {
	var (
		s    = []int{1, 2, 3, 4}
		last = slices.Get(s, -1)
	)
	fmt.Println(last)
	// Output: 4
}

func ExampleInsert() {
	var (
		s1 = []int{10, 15, 16}
		s2 = slices.Insert(s1, 1, 11, 12, 13, 14)
	)
	fmt.Println(s2)
	// Output: [10 11 12 13 14 15 16]
}

func ExampleReplaceN() {
	var (
		s1 = []int{99, 0, 0, 0, 97}
		s2 = slices.ReplaceN(s1, 1, 3, 98)
	)
	fmt.Println(s2)
	// Output: [99 98 97]
}

func ExampleReplaceTo() {
	var (
		s1 = []int{99, 0, 0, 0, 97}
		s2 = slices.ReplaceTo(s1, 1, -1, 98)
	)
	fmt.Println(s2)
	// Output: [99 98 97]
}

func ExampleRemoveN() {
	var (
		s1 = []int{1, 2, 3, 4, 5}
		s2 = slices.RemoveN(s1, -2, 2)
	)
	fmt.Println(s2)
	// Output: [1 2 3]
}

func ExampleRemoveTo() {
	var (
		s1 = []int{1, 2, 3, 4, 5}
		s2 = slices.RemoveTo(s1, -2, 0)
	)
	fmt.Println(s2)
	// Output: [1 2 3]
}

func ExampleEachx() {
	s := []int{100, 200, 300}
	_ = slices.Eachx(s, func(idx, val int) error {
		fmt.Println(idx, val)
		return nil
	})
	// Output:
	// 0 100
	// 1 200
	// 2 300
}

func ExampleMap() {
	var (
		s1 = []int{1, 2, 3, 4, 5}
		s2 = slices.Map(s1, func(val int) string { return string([]byte{byte('a' + val - 1)}) })
	)
	fmt.Println(s2)
	// Output: [a b c d e]
}

func ExampleAccum() {
	var (
		s   = []int{1, 2, 3, 4, 5}
		sum = slices.Accum(s, func(a, b int) int { return a + b })
	)
	fmt.Println(sum)
	// Output: 15
}

func ExampleFilter() {
	var (
		s     = []int{1, 2, 3, 4, 5, 6, 7}
		evens = slices.Filter(s, func(val int) bool { return val%2 == 0 })
	)
	fmt.Println(evens)
	// Output: [2 4 6]
}

func ExampleGroup() {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	groups := slices.Group(s, func(val int) string {
		if val%2 == 0 {
			return "even"
		}
		return "odd"
	})

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
