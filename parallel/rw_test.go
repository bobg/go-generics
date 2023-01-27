package parallel

import (
	"context"
	"reflect"
	"sync"
	"testing"
)

func TestRW(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	reader, writer := RW(ctx, 4)

	var (
		mu   sync.Mutex // protects vals
		vals []int
		ch   = make(chan struct{}) // signals that the first reader is under way
	)

	for i := 0; i < 3; i++ {
		i := i
		go reader(func(x int) {
			mu.Lock()
			if i == 0 {
				close(ch)
			}
			vals = append(vals, x)
			mu.Unlock()
		})
	}

	<-ch

	writer(func(x int) int { return x + 1 })

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go reader(func(x int) {
			mu.Lock()
			vals = append(vals, x)
			mu.Unlock()
			wg.Done()
		})
	}
	wg.Wait()

	want := []int{4, 4, 4, 5, 5, 5}
	if !reflect.DeepEqual(vals, want) {
		t.Errorf("got %v, want %v", vals, want)
	}
}
