package iter

// Page consumes inp one "page" at a time of up to pageSize elements,
// repeatedly calling a callback with a slice of the items consumed.
// The callback also gets a second argument that is false until the final call,
// when it is true.
//
// An error from the callback will terminate Page and return that error.
//
// The space for the slice is reused on each call to the callback.
//
// The slice in every non-final call of the callback is guaranteed to have a length of pageSize.
// The final call of the callback may contain an empty slice.
func Page[T any](inp Of[T], pageSize int, f func([]T, bool) error) error {
	page := make([]T, 0, pageSize)
	for inp.Next() {
		page = append(page, inp.Val())
		if len(page) >= pageSize {
			err := f(page, false)
			if err != nil {
				return err
			}
			page = page[:0]
		}
	}
	if err := inp.Err(); err != nil {
		return err
	}
	return f(page, true)
}
