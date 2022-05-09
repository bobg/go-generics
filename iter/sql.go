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
func SQL[T any](ctx context.Context, db QueryerContext, query string, args ...any) (Of[T], func() error) {
	var t T
	tt := reflect.TypeOf(t)
	if tt.Kind() != reflect.Struct {
		return nil, func() error { return fmt.Errorf("type parameter to SQL has %s kind but must be struct", tt.Kind()) }
	}
	nfields := tt.NumField()

	var (
		err   error
		errfn = func() error { return err }
	)

	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, errfn
	}

	ch := make(chan T)
	go func() {
		defer rows.Close()
		defer close(ch)

		for rows.Next() {
			var (
				row    T
				rowval = reflect.ValueOf(row)
				ptrs   = make([]any, 0, nfields)
			)
			for i := 0; i < nfields; i++ {
				addr := rowval.Field(i).Addr()
				ptrs = append(ptrs, addr.Interface())
			}
			err = rows.Scan(ptrs...)
			if err != nil {
				return
			}
			select {
			case <-ctx.Done():
				err = ctx.Err()
				return
			case ch <- row:
			}
		}

		err = rows.Err()
	}()

	return FromChanContext(ctx, ch), errfn
}
