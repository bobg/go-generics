package seqs

import (
	"iter"
	"sync"
)

// Dup duplicates the contents of an iterator,
// producing n new iterators,
// each containing the members of the original.
//
// An internal buffer grows to roughly the size
// of the difference between the output iterator that is farthest ahead in the stream,
// and the one that is farthest behind.
func Dup[T any](inp iter.Seq[T], n int) []iter.Seq[T] {
	next, stop := iter.Pull(inp)

	var (
		mu        sync.Mutex
		buf       []T
		bufOffset int
		offsets   = make([]int, n)
		result    []iter.Seq[T]
		wg        sync.WaitGroup
	)

	for i := 0; i < n; i++ {
		wg.Add(1)

		result = append(result, func(yield func(T) bool) {
			defer wg.Done()

			helper := func() bool {
				mu.Lock()
				defer mu.Unlock()

				bufEnd := bufOffset + len(buf)
				for offsets[i] >= bufEnd {
					val, ok := next()
					if !ok {
						return false
					}
					buf = append(buf, val)
					bufEnd++
				}
				val := buf[offsets[i]-bufOffset]
				offsets[i]++

				minOffset := offsets[0]
				for j := 1; j < n; j++ {
					if offsets[j] < minOffset {
						minOffset = offsets[j]
					}
				}

				buf = buf[minOffset-bufOffset:]
				bufOffset = minOffset

				return yield(val)
			}

			for helper() {
			}
		})
	}

	go func() {
		wg.Wait()
		stop()
	}()

	return result
}
