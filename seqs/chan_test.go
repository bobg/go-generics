package seqs

import (
	"context"
	"errors"
	"iter"
	"slices"
	"testing"
)

func TestToChan(t *testing.T) {
	ints := Ints(0, 1)
	ch := ToChan(ints)
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
}

func TestToChanContext(t *testing.T) {
	var (
		ch1  = make(chan int, 1)
		seq1 = FromChan(ch1)
		ctx  = context.Background()
	)
	ctx, cancel := context.WithCancel(ctx)
	ch2, errptr := ToChanContext(ctx, seq1)

	ch1 <- 1
	got, ok := <-ch2
	if !ok {
		t.Fatal("channel closed")
	}
	if got != 1 {
		t.Errorf("got %d, want 1", got)
	}

	ch1 <- 2
	got, ok = <-ch2
	if !ok {
		t.Fatal("channel closed")
	}
	if got != 2 {
		t.Errorf("got %d, want 2", got)
	}

	cancel()

	ch1 <- 3

	if _, ok := <-ch2; ok {
		t.Fatal("channel still open after context cancellation")
	}

	if !errors.Is(*errptr, context.Canceled) {
		t.Errorf("got error %v, want %v", *errptr, context.Canceled)
	}
}

func TestFromChan(t *testing.T) {
	ch := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch)
	}()

	var (
		seq  = FromChan(ch)
		got  = slices.Collect(seq)
		want = []int{0, 1, 2}
	)
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFromChanContext(t *testing.T) {
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

	it, errptr := FromChanContext(ctx, ch)
	next, stop := iter.Pull(it)
	defer stop()
	if _, ok := next(); !ok {
		t.Fatal("no first value in iterator")
	}

	cancel()

	if _, ok := next(); ok {
		t.Fatal("next value available after context cancellation")
	}

	stop()

	err := *errptr
	if !errors.Is(err, context.Canceled) {
		t.Errorf("got error %v, want %v", err, context.Canceled)
	}
}
