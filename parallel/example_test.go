package parallel_test

import (
	"context"
	"fmt"
	"sync"

	"github.com/bobg/go-generics/parallel"
)

func ExampleConsumers() {
	ctx := context.Background()

	// One of three goroutines prints incoming values.
	send, closefn := parallel.Consumers(ctx, 3, func(_ context.Context, _, val int) error {
		fmt.Println(val)
		return nil
	})

	// Caller produces values.
	for i := 1; i <= 5; i++ {
		err := send(i)
		if err != nil {
			panic(err)
		}
	}
	if err := closefn(); err != nil {
		panic(err)
	}
	// Unordered output:
	// 1
	// 2
	// 3
	// 4
	// 5
}

func ExampleProducers() {
	ctx := context.Background()

	// Five goroutines each produce their worker number and then exit.
	it := parallel.Producers(ctx, 5, func(_ context.Context, n int, send func(int) error) error {
		return send(n)
	})

	// Caller consumes the produced values.
	for it.Next() {
		fmt.Println(it.Val())
	}
	if err := it.Err(); err != nil {
		panic(err)
	}
	// Unordered output:
	// 0
	// 1
	// 2
	// 3
	// 4
}

func ExampleValues() {
	ctx := context.Background()

	// Five goroutines, each placing its worker number in the corresponding slot of the result slice.
	values, err := parallel.Values(ctx, 5, func(_ context.Context, n int) (int, error) {
		return n, nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(values)
	// Output:
	// [0 1 2 3 4]
}

func ExamplePool() {
	// Three workers available, each negating its input.
	pool := parallel.Pool(3, func(n int) (int, error) {
		return -n, nil
	})

	var wg sync.WaitGroup

	// Ten goroutines requesting work from those three workers.
	for i := 1; i <= 10; i++ {
		i := i // Go loop-var pitfall
		wg.Add(1)
		go func() {
			neg, err := pool(i)
			if err != nil {
				panic(err)
			}
			fmt.Println(neg)
			wg.Done()
		}()
	}

	wg.Wait()

	// Unordered output:
	// -1
	// -2
	// -3
	// -4
	// -5
	// -6
	// -7
	// -8
	// -9
	// -10
}

func ExampleProtect() {
	// A caller is supplied with a reader and a writer
	// for purposes of accessing and updating the protected value safely
	// (in this case an int, initially 4).
	reader, writer, closer := parallel.Protect(4)
	defer closer()

	// Call the reader in three concurrent goroutines, each printing the protected value.
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(reader())
			wg.Done()
		}()
	}
	wg.Wait()

	// Increment the protected value.
	writer(reader() + 1)

	// Call the reader in three concurrent goroutines, each printing the protected value.
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(reader())
			wg.Done()
		}()
	}
	wg.Wait()

	// Output:
	// 4
	// 4
	// 4
	// 5
	// 5
	// 5
}
