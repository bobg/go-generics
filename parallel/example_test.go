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

func ExampleRW() {
	ctx := context.Background()

	// A cancelable context releases resources from the RW call
	// after no new reader or writer calls are needed.
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// A caller is supplied only with reader and writer
	// for purposes of accessing and updating the protected value
	// (in this case an int, initially 4).
	reader, writer := parallel.RW(ctx, 4)

	// This channel makes sure the writer call doesn't happen
	// before at least one of the reader calls is under way.
	ch := make(chan struct{})

	for i := 0; i < 3; i++ {
		i := i // Go loop var pitfall
		go reader(func(x int) {
			if i == 0 {
				close(ch)
			}
			fmt.Println(x)
		})
	}

	// Wait for at least one reader to be running.
	// (Otherwise it's possible for the writer call to happen
	// before any of the reader goroutines manages to launch.)
	<-ch

	// Once there are some readers in progress,
	// RW will make sure that they all complete
	// before allowing a writer to run.
	writer(func(x int) int { return x + 1 })

	// Similarly, no additional readers (or writers) can run
	// until the writer is finished.
	// The WaitGroup here ensures the test does not exit before all readers complete.
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go reader(func(x int) {
			fmt.Println(x)
			wg.Done()
		})
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
