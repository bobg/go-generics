package iter

import "testing"

func BenchmarkIter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var (
			ints      = Ints(1, 1)
			first1000 = FirstN(ints, 1000)
			odds      = Filter(first1000, func(x int) bool { return x%2 == 1 })
			squares   = Map(odds, func(x int) (int, error) { return x * x, nil })
			sums      = Accum(squares, func(x, y int) (int, error) { return x + y, nil })
		)
		sum, err := LastN(sums, 1)
		if err != nil {
			b.Fatal(err)
		}
		if len(sum) != 1 {
			b.Fatalf("len(sum) = %d, want 1", len(sum))
		}
	}
}
