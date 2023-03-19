package internal

// RotateSlice rotates a slice in place by n places to the right.
// (With negative n, it's to the left.)
// Example: RotateSlice([D, E, A, B, C], 3) -> [A, B, C, D, E]
//
// This function is here
// in order to resolve an import cycle
// that would otherwise exist
// between iter and slices.
func RotateSlice[S ~[]T, T any](s S, n int) {
	if n < 0 {
		// Convert left-rotation to right-rotation.
		n = -n
		n %= len(s)
		n = len(s) - n
	} else {
		n %= len(s)
	}
	if n == 0 {
		return
	}
	tmp := make([]T, n)
	copy(tmp, s[len(s)-n:])
	copy(s[n:], s)
	copy(s, tmp)
}
