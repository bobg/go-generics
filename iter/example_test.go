package iter_test

import (
	"context"
	"fmt"

	"github.com/bobg/go-generics/iter"
)

func ExampleAccum() {
	ints := iter.Ints(1, 1)        // All integers starting at 1
	first5 := iter.FirstN(ints, 5) // First 5 integers
	sums := iter.Accum(first5, func(a, b int) (int, error) { return a + b, nil })
	for sums.Next() {
		fmt.Println(sums.Val())
	}
	if err := sums.Err(); err != nil {
		panic(err)
	}
	// Output:
	// 1
	// 3
	// 6
	// 10
	// 15
}

func ExampleFromMap() {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	it := iter.FromMap(m)
	for it.Next() {
		val := it.Val()
		fmt.Println(val.X, val.Y)
	}
	if err := it.Err(); err != nil {
		panic(err)
	}
	// Unordered output:
	// one 1
	// two 2
	// three 3
}

func ExampleDup() {
	var (
		ints    = iter.Ints(1, 1)       // All integers starting at 1
		first10 = iter.FirstN(ints, 10) // First 10 integers
		dups    = iter.Dup(first10, 2)  // Two copies of the first10 iterator
		evens   = iter.Filter(dups[0], func(val int) bool { return val%2 == 0 })
		odds    = iter.Filter(dups[1], func(val int) bool { return val%2 == 1 })
	)
	evensSlice, err := iter.ToSlice(evens)
	if err != nil {
		panic(err)
	}
	fmt.Println(evensSlice)
	oddsSlice, err := iter.ToSlice(odds)
	if err != nil {
		panic(err)
	}
	fmt.Println(oddsSlice)
	// Output:
	// [2 4 6 8 10]
	// [1 3 5 7 9]
}

func ExampleGo() {
	it := iter.Go(context.Background(), func(send func(val int) error) error {
		if err := send(1); err != nil {
			return err
		}
		if err := send(2); err != nil {
			return err
		}
		return send(3)
	})
	slice, err := iter.ToSlice(it)
	if err != nil {
		panic(err)
	}
	fmt.Println(slice)
	// Output:
	// [1 2 3]
}

func ExampleZip() {
	var (
		letters = iter.FromSlice([]string{"a", "b", "c", "d"})
		nums    = iter.FromSlice([]int{1, 2, 3})
		pairs   = iter.Zip(letters, nums)
	)
	for pairs.Next() {
		val := pairs.Val()
		fmt.Println(val.X, val.Y)
	}
	if err := pairs.Err(); err != nil {
		panic(err)
	}
	// Output:
	// a 1
	// b 2
	// c 3
	// d 0
}
