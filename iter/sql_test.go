package iter

import (
	"context"
	"database/sql"
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

	it, err := SQL[employee](ctx, db, "SELECT name, salary FROM employees ORDER BY name")
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
}
