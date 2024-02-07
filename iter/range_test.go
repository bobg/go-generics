//go:build go1.23 || goexperiment.rangefunc

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

func ExamplePull() {
	var (
		ints       = Ints(1, 1) // All integers starting at 1
		next, stop = Pull(All(ints))
	)
	for i := 0; i < 5; i++ {
		val, ok := next()
		fmt.Println(val, ok)
	}
	stop()
	val, ok := next()
	fmt.Println(val, ok)
	// Output:
	// 1 true
	// 2 true
	// 3 true
	// 4 true
	// 5 true
	// 0 false
}

func ExamplePull2() {
	var (
		names      = FromSlice([]string{"Alice", "Bob", "Carol", "Dave"})
		namelens   = FromSlice([]int{5, 3, 5, 4})
		pairs      = Zip(names, namelens)
		next, stop = Pull2(AllPairs(pairs))
	)
	defer stop()

	for {
		name, namelen, ok := next()
		fmt.Println(name, namelen, ok)
		if !ok {
			break
		}
	}
	// Output:
	// Alice 5 true
	// Bob 3 true
	// Carol 5 true
	// Dave 4 true
	//  0 false
}
