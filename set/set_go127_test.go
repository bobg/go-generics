//go:build go1.27

package set_test

import (
	"encoding/json/v2"
	"github.com/bobg/go-generics/v4/set"
	"testing"
)

func TestSetJSON(t *testing.T) {
	t.Parallel()

	testData := []struct {
		name    string
		set     set.Of[string]
		json    string
		jsonAlt string
	}{
		{
			name: "null",
			set:  nil,
			json: `null`,
		},
		{
			name: "empty",
			set:  make(set.Of[string]),
			json: `[]`,
		},
		{
			name: "single",
			set:  set.New("one"),
			json: `["one"]`,
		},
		{
			name: "multiple",
			set:  set.New("one", "two"),

			// sets are unordered so either of the two outputs are acceptable
			// unfortunately this cannot be made deterministic without making T cmp.Orderable instead of just comparable
			json:    `["one","two"]`,
			jsonAlt: `["two","one"]`,
		},
	}

	for _, d := range testData {
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()

			t.Run("marshal", func(t *testing.T) {
				t.Parallel()

				b, err := json.Marshal(d.set)
				if err != nil {
					t.Fatal("unexpected error:", err)
				}

				if string(b) != d.json && (d.jsonAlt == "" || string(b) != d.jsonAlt) {
					t.Errorf("got %s, want %s", string(b), d.json)
				}
			})

			t.Run("unmarshal", func(t *testing.T) {
				t.Parallel()

				var s set.Of[string]
				err := json.Unmarshal([]byte(d.json), &s)
				if err != nil {
					t.Fatal("unexpected error:", err)
				}

				if !s.Equal(d.set) {
					t.Errorf("got %s, want %s", s, d.set)
				}
			})
		})
	}
}
