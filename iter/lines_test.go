package iter

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLines(t *testing.T) {
	indep, err := os.ReadFile("testdata/indep.txt")
	if err != nil {
		t.Fatal(err)
	}

	var (
		lines = Lines(bytes.NewReader(indep))
		got   = new(bytes.Buffer)
	)
	for lines.Next() {
		line := lines.Val()
		fmt.Fprintln(got, line)
	}
	if err := lines.Err(); err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(string(indep), got.String()); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestLongLines(t *testing.T) {
	indep, err := os.ReadFile("testdata/indep.txt")
	if err != nil {
		t.Fatal(err)
	}

	var (
		lines = LongLines(context.Background(), bytes.NewReader(indep))
		got   = new(bytes.Buffer)
	)
	for lines.Next() {
		r := lines.Val()
		line, err := io.ReadAll(r)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Fprintln(got, string(line))
	}
	if err := lines.Err(); err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(string(indep), got.String()); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}
