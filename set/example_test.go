package set_test

import (
	"fmt"

	"github.com/bobg/go-generics/set"
)

func ExampleDiff() {
	s1 := set.New(1, 2, 3, 4, 5)
	s2 := set.New(4, 5, 6, 7, 8)
	diff := set.Diff(s1, s2)
	err := diff.Each(func(val int) error {
		fmt.Println(val)
		return nil
	})
	if err != nil {
		panic(err)
	}
	// Unordered output:
	// 1
	// 2
	// 3
}

func ExampleIntersect() {
	s1 := set.New(1, 2, 3, 4, 5)
	s2 := set.New(4, 5, 6, 7, 8)
	inter := set.Intersect(s1, s2)
	err := inter.Each(func(val int) error {
		fmt.Println(val)
		return nil
	})
	if err != nil {
		panic(err)
	}
	// Unordered output:
	// 4
	// 5
}

func ExampleUnion() {
	s1 := set.New(1, 2, 3, 4, 5)
	s2 := set.New(4, 5, 6, 7, 8)
	union := set.Union(s1, s2)
	err := union.Each(func(val int) error {
		fmt.Println(val)
		return nil
	})
	if err != nil {
		panic(err)
	}
	// Unordered output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
}

func ExampleOf() {
	s := set.New(1, 2, 3, 4, 5)
	fmt.Println("1 is in the set?", s.Has(1))
	fmt.Println("100 is in the set?", s.Has(100))
	s.Add(100)
	fmt.Println("100 is in the set?", s.Has(100))
	fmt.Println("set size is", s.Len())
	s.Del(100)
	fmt.Println("100 is in the set?", s.Has(100))
	fmt.Println("set size is", s.Len())
	// Output:
	// 1 is in the set? true
	// 100 is in the set? false
	// 100 is in the set? true
	// set size is 6
	// 100 is in the set? false
	// set size is 5
}
