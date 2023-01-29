package parallel

import (
	"context"
	"reflect"
	"testing"
)

func TestProtect(t *testing.T) {
	reader, writer, closer := Protect(4)
	defer closer()

	ctx := context.Background()

	vals, err := Values(ctx, 3, func(_ context.Context, _ int) (int, error) { return reader(), nil })
	if err != nil {
		t.Fatal(err)
	}

	writer(reader() + 1)

	vals2, err := Values(ctx, 3, func(_ context.Context, _ int) (int, error) { return reader(), nil })
	if err != nil {
		t.Fatal(err)
	}

	vals = append(vals, vals2...)

	want := []int{4, 4, 4, 5, 5, 5}
	if !reflect.DeepEqual(vals, want) {
		t.Errorf("got %v, want %v", vals, want)
	}
}
