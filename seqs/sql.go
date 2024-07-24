package seqs

import (
	"context"
	"database/sql"
	"fmt"
	"iter"
	"reflect"
)

// QueryerContext is a minimal interface satisfied by *sql.DB
// (from database/sql).
type QueryerContext interface {
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
}

// SQL performs a query against db and returns the results as an iterator of type T.
//
// If the query produces a single value per row,
// T may be any scalar type (bool, int, float, string)
// into which the values can be scanned.
//
// Otherwise T must be a struct type whose fields have the same types,
// in the same order,
// as the values being queried.
// The values produced by the iterator will be instances of that struct type,
// with fields populated by the queried values.
func SQL[T any](ctx context.Context, db QueryerContext, query string, args ...any) (iter.Seq[T], *error) {
	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return Empty[T], &err
	}

	f := func(yield func(T) bool) {
		defer rows.Close()
		err = sqlHelper[T](ctx, rows, yield)
	}

	return f, &err
}

// Prepared is like [SQL] but uses a prepared [sql.Stmt] instead of a database and string query.
// It is the caller's responsibility to close the statement.
func Prepared[T any](ctx context.Context, stmt *sql.Stmt, args ...any) (iter.Seq[T], *error) {
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return Empty[T], &err
	}

	f := func(yield func(T) bool) {
		defer rows.Close()
		err = sqlHelper[T](ctx, rows, yield)
	}

	return f, &err
}

type sqlKindError struct {
	kind reflect.Kind
}

func (e sqlKindError) Error() string {
	return fmt.Sprintf("type parameter has %s kind but must be struct", e.kind)
}

func sqlHelper[T any](ctx context.Context, rows *sql.Rows, yield func(T) bool) error {
	var t T
	tt := reflect.TypeOf(t)
	switch tt.Kind() {
	case reflect.Struct:
		return sqlHelperStruct[T](ctx, tt, rows, yield)

	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.String:
		return sqlHelperScalar[T](ctx, rows, yield)

	default:
		return sqlKindError{kind: tt.Kind()}
	}
}

func sqlHelperStruct[T any](ctx context.Context, rowtype reflect.Type, rows *sql.Rows, yield func(T) bool) error {
	nfields := rowtype.NumField()

	for rows.Next() {
		var (
			// Note: We cannot use:
			//   var row T
			//   rowval = reflect.ValueOf(row)
			// because the result is not addressable.
			// We have to let the reflect package create its own instance of T.

			rowptrval = reflect.New(rowtype)
			rowval    = rowptrval.Elem()
			ptrs      = make([]any, 0, nfields)
		)

		for i := 0; i < nfields; i++ {
			addr := rowval.Field(i).Addr()
			ptrs = append(ptrs, addr.Interface())
		}
		if err := rows.Scan(ptrs...); err != nil {
			return fmt.Errorf("scanning row: %w", err)
		}

		if err := ctx.Err(); err != nil {
			return err
		}

		if !yield(rowval.Interface().(T)) {
			return nil
		}
	}

	return rows.Err()
}

func sqlHelperScalar[T any](ctx context.Context, rows *sql.Rows, yield func(T) bool) error {
	for rows.Next() {
		var val T
		if err := rows.Scan(&val); err != nil {
			return fmt.Errorf("scanning row: %w", err)
		}

		if err := ctx.Err(); err != nil {
			return err
		}

		if !yield(val) {
			return nil
		}
	}
	return rows.Err()
}
