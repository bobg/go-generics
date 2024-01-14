//go:build go1.23 || rangefunc

package iter

import "fmt"

func ExampleAll() {
	var (
		ints   = Ints(1, 1)      // All integers starting at 1
		first5 = FirstN(ints, 5) // First 5 integers
	)
	for val := range All(first5) {
		fmt.Println(val)
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}

func ExampleAllCount() {
	names := FromSlice([]string{"Alice", "Bob", "Carol"})

	for i, name := range AllCount(names) {
		fmt.Println(i, name)
	}
	// Output:
	// 0 Alice
	// 1 Bob
	// 2 Carol
}

func ExampleAllPairs() {
	var (
		letters = FromSlice([]string{"a", "b", "c", "d"})
		nums    = FromSlice([]int{1, 2, 3})
		pairs   = Zip(letters, nums)
	)

	for letter, num := range AllPairs(pairs) {
		fmt.Println(letter, num)
	}
	// Output:
	// a 1
	// b 2
	// c 3
	// d 0
}
