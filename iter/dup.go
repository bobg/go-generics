package iter

import "sync"

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
		result = append(result, Gen(func() (T, bool) {
			mu.Lock()
			defer mu.Unlock()

			bufEnd := bufOffset + len(buf)
			for offsets[i] >= bufEnd {
				if inp.Next() {
					buf = append(buf, inp.Val())
					bufEnd++
				} else {
					var zero T
					return zero, false
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

			return val, true
		}))
	}

	return result
}
