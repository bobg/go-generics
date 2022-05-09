package iter

import "testing"

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
