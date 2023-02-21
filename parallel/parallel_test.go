package parallel

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/bobg/go-generics/v2/set"
)

func TestValues(t *testing.T) {
	got, err := Values(context.Background(), 100, func(_ context.Context, n int) (int, error) { return n, nil })
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 100 {
		t.Errorf("got len %d, want 100", len(got))
	}
	for i := 0; i < 100; i++ {
		if got[i] != i {
			t.Errorf("got[%d] is %d, want %d", i, got[i], i)
		}
	}
}

func TestProducers(t *testing.T) {
	it := Producers(context.Background(), 10, func(_ context.Context, n int, send func(int) error) error {
		for i := 0; i < 10; i++ {
			err := send(10*n + i)
			if err != nil {
				return err
			}
		}
		return nil
	})
	got := set.New[int]()
	for it.Next() {
		got.Add(it.Val())
	}
	if err := it.Err(); err != nil {
		t.Fatal(err)
	}
	if got.Len() != 100 {
		t.Errorf("got %d values, want 100", got.Len())
	}
	for i := 0; i < 100; i++ {
		if !got.Has(i) {
			t.Errorf("%d missing from result", i)
		}
	}
}

func TestConsumers(t *testing.T) {
	var (
		mu  sync.Mutex
		got = set.New[int]()
	)

	sendfn, closefn := Consumers(context.Background(), 10, func(_ context.Context, n, val int) error {
		mu.Lock()
		got.Add(val)
		mu.Unlock()
		return nil
	})
	for i := 0; i < 100; i++ {
		err := sendfn(i)
		if err != nil {
			t.Fatal(err)
		}
	}
	err := closefn()
	if err != nil {
		t.Fatal(err)
	}
	if got.Len() != 100 {
		t.Errorf("got %d values, want 100", got.Len())
	}
	for i := 0; i < 100; i++ {
		if !got.Has(i) {
			t.Errorf("%d missing from result", i)
		}
	}
}

func TestPool(t *testing.T) {
	var (
		running, max int
		mu           sync.Mutex
		cond         = sync.NewCond(&mu)
		unblocked    = make(map[int]bool)
	)

	call := Pool(10, func(n int) (int, error) {
		mu.Lock()
		running++
		if running > max {
			max = running
		}

		for !unblocked[n] {
			cond.Wait()
		}

		running--
		mu.Unlock()

		return n, nil
	})

	var (
		errch = make(chan error, 100)
		wg    sync.WaitGroup
	)
	for i := 0; i < 100; i++ {
		i := i // Go loop var pitfall
		wg.Add(1)
		go func() {
			got, err := call(i)
			if err != nil {
				errch <- err
			}
			if got != i {
				errch <- fmt.Errorf("got %d, want %d", got, i)
			}
			wg.Done()
		}()
	}

	for i := 0; i < 100; i++ {
		mu.Lock()
		unblocked[i] = true
		cond.Broadcast()
		mu.Unlock()
	}

	go func() {
		wg.Wait()
		close(errch)
	}()

	for err := range errch {
		t.Error(err)
	}

	if max > 10 {
		t.Errorf("max is %d, want <=10", max)
	}
}
