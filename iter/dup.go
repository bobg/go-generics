package iter

import "sync"

// Dup duplicates the contents of an iterator,
// producing n new iterators,
// each containing the members of the original.
//
// An internal buffer grows to roughly the size
// of the difference between the output iterator that is farthest ahead in the stream,
// and the one that is farthest behind.
func Dup[T any](inp Of[T], n int) []Of[T] {
	var (
		mu        sync.Mutex
		buf       []T
		bufOffset int
		offsets   = make([]int, n)
		result    []Of[T]
	)

	for i := 0; i < n; i++ {
		i := i // Go loop-var pitfall
		result = append(result, Gen(func() (T, bool, error) {
			mu.Lock()
			defer mu.Unlock()

			bufEnd := bufOffset + len(buf)
			for offsets[i] >= bufEnd {
				if inp.Next() {
					buf = append(buf, inp.Val())
					bufEnd++
				} else {
					var zero T
					return zero, false, inp.Err()
				}
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

			return val, true, nil
		}))
	}

	return result
}
