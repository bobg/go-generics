package seqs

import "iter"

// Partition splits a sequence of values into multiple sequences by computing a partition key for each value.
// The output is a sequence of key-subsequence pairs.
// Each subsequence contains all the values that have the same key,
// in their original order.
//
// If the caller fails to consume values from one subsequence,
// it could block the production of values in other sequences.
// To avoid deadlock it is safest to launch a separate goroutine to consume each subsequence.
func Partition[T any, K comparable, F ~func(T) K](inp iter.Seq[T], f F) iter.Seq2[K, iter.Seq[T]] {
	pairs, _ := Go(func(outerChan chan<- Pair[K, iter.Seq[T]]) error {
		m := make(map[K]chan<- T)

		defer func() {
			for _, outerChan := range m {
				close(outerChan)
			}
		}()

		for val := range inp {
			k := f(val)
			innerChan, ok := m[k]
			if !ok {
				var (
					ch  = make(chan T)
					seq = FromChan(ch)
				)
				m[k] = ch
				outerChan <- Pair[K, iter.Seq[T]]{X: k, Y: seq}
				innerChan = ch
			}
			innerChan <- val
		}

		return nil
	})

	return FromPairs(pairs)
}
