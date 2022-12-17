package iter

import (
	"context"

	"github.com/bobg/go-generics/slices"
)

// Permutations produces an iterator over all permutations of s.
// It uses Heap's Algorithm.
// See https://en.wikipedia.org/wiki/Heap%27s_algorithm.
//
// If s is [1 2 3], this function will produce:
//
//	[1 2 3] [2 1 3] [3 1 2] [1 3 2] [2 3 1] [3 2 1]
func Permutations[T any](ctx context.Context, s []T) Of[[]T] {
	if len(s) == 0 {
		return FromSlice[[]T](nil)
	}
	return Go(ctx, func(ch chan<- []T) error {
		return permutations(ctx, slices.Dup(s), len(s), ch)
	})
}

func permutations[T any](ctx context.Context, s []T, n int, ch chan<- []T) error {
	if n == 1 {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case ch <- slices.Dup(s):
		}
		return nil
	}

	if err := permutations(ctx, s, n-1, ch); err != nil {
		return err
	}

	for i := 0; i < n-1; i++ {
		if n%2 == 0 {
			s[i], s[n-1] = s[n-1], s[i]
		} else {
			s[0], s[n-1] = s[n-1], s[0]
		}

		if err := permutations(ctx, s, n-1, ch); err != nil {
			return err
		}
	}
	return nil
}

// Combinations produces an iterator over all n-length combinations of distinct elements from s.
//
// If s is [1 2 3] and n is 2, this function will produce:
//
//	[1 2] [1 3] [2 3]
func Combinations[T any](ctx context.Context, s []T, n int) Of[[]T] {
	if n == 0 {
		return FromSlice[[]T](nil)
	}
	if n > len(s) {
		return FromSlice[[]T](nil)
	}
	if n == len(s) {
		return FromSlice([][]T{s})
	}
	return Go(ctx, func(ch chan<- []T) error {
		counters := make([]int, n)
		for i := 0; i < n; i++ {
			counters[i] = i
		}
		buf := make([]T, n)

	OUTER:
		for {
			for i := 0; i < n; i++ {
				buf[i] = s[counters[i]]
			}
			select {
			case <-ctx.Done():
				return ctx.Err()
			case ch <- slices.Dup(buf):
			}

			for i := n - 1; i >= 0; i-- {
				maxForThisPos := len(s) - 1 - ((n - 1) - i)
				if counters[i] < maxForThisPos {
					counters[i]++
					for j := i + 1; j < n; j++ {
						counters[j] = counters[j-1] + 1
					}
					continue OUTER
				}
			}

			return nil
		}
	})
}

// CombinationsWithReplacement produces an iterator over all n-length combinations of possibly repeated elements from s.
//
// If s is [1 2 3] and n is 2, this function will produce:
//
//	[1 1] [1 2] [1 3] [2 2] [2 3] [3 3]
func CombinationsWithReplacement[T any](ctx context.Context, s []T, n int) Of[[]T] {
	if n == 0 {
		return FromSlice[[]T](nil)
	}
	if n > len(s) {
		return FromSlice[[]T](nil)
	}
	return Go(ctx, func(ch chan<- []T) error {
		counters := make([]int, n)
		buf := make([]T, n)

	OUTER:
		for {
			for i := 0; i < n; i++ {
				buf[i] = s[counters[i]]
			}
			select {
			case <-ctx.Done():
				return ctx.Err()
			case ch <- slices.Dup(buf):
			}

			for i := n - 1; i >= 0; i-- {
				if counters[i] < len(s)-1 {
					counters[i]++
					for j := i + 1; j < n; j++ {
						counters[j] = counters[i]
					}
					continue OUTER
				}
			}

			return nil
		}
	})
}
