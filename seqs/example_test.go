package seqs_test

import (
	"fmt"
	"slices"

	"github.com/bobg/go-generics/v4/seqs"
)

func ExampleAccum() {
	var (
		ints   = seqs.Ints(1, 1)      // All integers starting at 1
		first5 = seqs.FirstN(ints, 5) // First 5 integers
		sums   = seqs.Accum(first5, func(a, b int) int { return a + b })
	)
	for val := range sums {
		fmt.Println(val)
	}
	// Output:
	// 1
	// 3
	// 6
	// 10
	// 15
}

func ExampleDup() {
	var (
		ints       = seqs.Ints(1, 1)       // All integers starting at 1
		first10    = seqs.FirstN(ints, 10) // First 10 integers
		dups       = seqs.Dup(first10, 2)  // Two copies of the first10 iterator
		evens      = seqs.Filter(dups[0], func(val int) bool { return val%2 == 0 })
		odds       = seqs.Filter(dups[1], func(val int) bool { return val%2 == 1 })
		evensSlice = slices.Collect(evens)
		oddsSlice  = slices.Collect(odds)
	)
	fmt.Println(evensSlice)
	fmt.Println(oddsSlice)
	// Output:
	// [2 4 6 8 10]
	// [1 3 5 7 9]
}

func ExampleGo() {
	it, errptr := seqs.Go(func(ch chan<- int) error {
		ch <- 1
		ch <- 2
		ch <- 3
		return nil
	})
	slice := slices.Collect(it)
	if *errptr != nil {
		panic(*errptr)
	}
	fmt.Println(slice)
	// Output:
	// [1 2 3]
}

func ExampleZip() {
	var (
		letters = slices.Values([]string{"a", "b", "c", "d"})
		nums    = slices.Values([]int{1, 2, 3})
		pairs   = seqs.Zip(letters, nums)
	)
	for x, y := range pairs {
		fmt.Println(x, y)
	}
	// Output:
	// a 1
	// b 2
	// c 3
	// d 0
}

func ExamplePages() {
	var (
		ints    = seqs.Ints(1, 1)       // All integers starting at 1
		first10 = seqs.FirstN(ints, 10) // First 10 integers
	)
	pages := seqs.Pages(first10, 3)
	for page := range pages {
		fmt.Println(page)
	}
	// Output:
	// [1 2 3]
	// [4 5 6]
	// [7 8 9]
	// [10]
}
