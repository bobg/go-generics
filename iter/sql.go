package iter

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
)

// QueryerContext is a minimal interface satisfied by *sql.DB
// (from database/sql).
type QueryerContext interface {
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
}

// SQL performs a query against db and returns the results as an iterator of type T.
// T must be a struct type whose fields have the same types,
// in the same order,
// as the values being queried.
// The values produced by the iterator will be instances of that struct type,
// with fields populated by the queried values.
func SQL[T any](ctx context.Context, db QueryerContext, query string, args ...any) (Of[T], error) {
	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("executing query: %w", err)
	}

	return sqlhelper[T](ctx, rows)
}

// Prepared is like [SQL] but uses a prepared [sql.Stmt] instead of a database and string query.
// It is the caller's responsibility to close the statement.
func Prepared[T any](ctx context.Context, stmt *sql.Stmt, args ...any) (Of[T], error) {
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, fmt.Errorf("executing query: %w", err)
	}

	return sqlhelper[T](ctx, rows)
}

type sqlKindError struct {
	kind reflect.Kind
}

func (e sqlKindError) Error() string {
	return fmt.Sprintf("type parameter has %s kind but must be struct", e.kind)
}

func sqlhelper[T any](ctx context.Context, rows *sql.Rows) (Of[T], error) {
	var t T
	tt := reflect.TypeOf(t)
	if tt.Kind() != reflect.Struct {
		return nil, sqlKindError{kind: tt.Kind()}
	}
	nfields := tt.NumField()

	res := Go(func(ch chan<- T) error {
		defer rows.Close()

		for rows.Next() {
			var (
				row T

				// Note: We cannot use:
				//   rowval = reflect.ValueOf(row)
				// because the result is not addressable.
				// We have to let the reflect package create its own instance of T.

				rowtype   = reflect.TypeOf(row)
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

			select {
			case <-ctx.Done():
				return ctx.Err()
			case ch <- rowval.Interface().(T):
			}
		}
		return rows.Err()
	})

	return res, nil
}
