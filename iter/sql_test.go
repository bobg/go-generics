package iter

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"reflect"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

const schema = `
CREATE TABLE employees (
  name TEXT NOT NULL,
  salary INT NOT NULL
);
`

func TestSQL(t *testing.T) {
	f, err := os.CreateTemp("", "sqltest")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())
	if err = f.Close(); err != nil {
		t.Fatal(err)
	}

	db, err := sql.Open("sqlite3", f.Name())
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	_, err = db.ExecContext(ctx, schema)
	if err != nil {
		t.Fatal(err)
	}

	type employee struct {
		Name   string
		Salary int
	}

	want := []employee{{
		Name:   "alice",
		Salary: 100000,
	}, {
		Name:   "bill",
		Salary: 80000,
	}, {
		Name:   "carol",
		Salary: 90000,
	}, {
		Name:   "dave",
		Salary: 40000,
	}}

	for _, emp := range want {
		_, err := db.ExecContext(ctx, "INSERT INTO employees (name, salary) VALUES ($1, $2)", emp.Name, emp.Salary)
		if err != nil {
			t.Fatal(err)
		}
	}

	const q = `SELECT name, salary FROM employees ORDER BY name`

	t.Run("Scalar", func(t *testing.T) {
		it, err := SQL[string](ctx, db, `SELECT name FROM employees ORDER BY name DESC`)
		if err != nil {
			t.Fatal(err)
		}
		got, err := ToSlice(it)
		if err != nil {
			t.Fatal(err)
		}
		want := []string{"dave", "carol", "bill", "alice"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("SQL", func(t *testing.T) {
		it, err := SQL[employee](ctx, db, q)
		if err != nil {
			t.Fatal(err)
		}
		got, err := ToSlice(it)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Prepare", func(t *testing.T) {
		stmt, err := db.PrepareContext(ctx, q)
		if err != nil {
			t.Fatal(err)
		}
		it, err := Prepared[employee](ctx, stmt)
		if err != nil {
			t.Fatal(err)
		}
		got, err := ToSlice(it)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("KindError", func(t *testing.T) {
		_, err := SQL[*int](ctx, db, q)

		var e sqlKindError
		if !errors.As(err, &e) {
			e.kind = reflect.TypeOf(0).Kind()
			t.Errorf("got %v, want %v", err, e)
		}
	})
}
