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
	var t T
	tt := reflect.TypeOf(t)
	if tt.Kind() != reflect.Struct {
		return nil, fmt.Errorf("type parameter to SQL has %s kind but must be struct", tt.Kind())
	}
	nfields := tt.NumField()

	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("executing query: %w", err)
	}

	res := Go(ctx, func(ch chan<- T) error {
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
			err = rows.Scan(ptrs...)
			if err != nil {
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
