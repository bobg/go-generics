package seqs

import (
	"reflect"
	"slices"
	"sync"
	"testing"
)

func TestPartition(t *testing.T) {
	var (
		ints       = Ints(0, 1)
		first10    = FirstN(ints, 10)
		partitions = Partition(first10, func(n int) int { return n % 3 })
		mu         sync.Mutex
		m          = make(map[int][]int)
		wg         sync.WaitGroup
	)

	for k, seq := range partitions {
		wg.Add(1)
		go func() {
			members := slices.Collect(seq)
			mu.Lock()
			m[k] = members
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	want := map[int][]int{
		0: {0, 3, 6, 9},
		1: {1, 4, 7},
		2: {2, 5, 8},
	}

	if !reflect.DeepEqual(m, want) {
		t.Errorf("got %v, want %v", m, want)
	}
}
