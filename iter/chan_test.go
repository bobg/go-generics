package iter

import (
	"context"
	"errors"
	"reflect"
	"testing"
)

func TestToChan(t *testing.T) {
	ints := Ints(0, 1)
	ch, errfn := ToChan(ints)
	want := 0
	for got := range ch {
		if got != want {
			t.Fatalf("got %d, want %d", got, want)
		}
		want++
		if want > 10 {
			break
		}
	}
	if err := errfn(); err != nil {
		t.Fatal(err)
	}
}

func TestFromChan(t *testing.T) {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()
	it := FromChan(ch)
	ints, err := ToSlice(it)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(ints, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
		t.Errorf("got %v, want 1 through 10", ints)
	}
}

func TestChanContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 1; i <= 10; i++ {
			select {
			case <-ctx.Done():
				return
			case ch <- i:
			}
		}
	}()

	it := FromChan(ch, WithContext(ctx))
	if !it.Next() {
		t.Fatal("no first value in iterator")
	}

	cancel()

	if it.Next() {
		t.Fatal("next value available after context cancellation")
	}
	err := it.Err()
	if !errors.Is(err, context.Canceled) {
		t.Errorf("got error %v, want %v", err, context.Canceled)
	}
}
