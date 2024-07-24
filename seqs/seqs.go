package seqs

import (
	"fmt"
	"iter"
	"math"
	"reflect"
)

type Pair[T, U any] struct {
	X T
	Y U
}

// Seq1 changes an [iter.Seq2] to an [iter.Seq] of [Pair]s.
func Seq1[T, U any](inp iter.Seq2[T, U]) iter.Seq[Pair[T, U]] {
	return func(yield func(Pair[T, U]) bool) {
		for x, y := range inp {
			if !yield(Pair[T, U]{X: x, Y: y}) {
				return
			}
		}
	}
}

// Enumerate changes an [iter.Seq] to an [iter.Seq2] of (index, val) pairs.
func Enumerate[T any](inp iter.Seq[T]) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		i := 0
		for x := range inp {
			if !yield(i, x) {
				return
			}
			i++
		}
	}
}

// Seq2 changes an [iter.Seq] of [Pair]s to an [iter.Seq2].
func Seq2[T, U any](inp iter.Seq[Pair[T, U]]) iter.Seq2[T, U] {
	return func(yield func(T, U) bool) {
		for val := range inp {
			if !yield(val.X, val.Y) {
				return
			}
		}
	}
}

func IntRanger(inp any) iter.Seq[int] {
	val := reflect.ValueOf(inp)

	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		return func(yield func(int) bool) {
			for i := 0; i < val.Len(); i++ {
				if !yield(i) {
					return
				}
			}
		}

	case reflect.Func:
		return val.Interface().(func(func(int) bool))

	case reflect.String:
		return func(yield func(int) bool) {
			for pos := range val.String() {
				if !yield(pos) {
					return
				}
			}
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if val.Kind() == reflect.Int64 && val.Int() > math.MaxInt {
			panic("int64 too large for int")
		}

		return func(yield func(int) bool) {
			for i := int64(0); i < val.Int(); i++ {
				if !yield(int(i)) {
					return
				}
			}
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if val.Kind() == reflect.Uint64 && val.Uint() > math.MaxInt {
			panic("uint64 too large for int")
		}

		return func(yield func(int) bool) {
			for i := uint64(0); i < val.Uint(); i++ {
				if !yield(int(i)) {
					return
				}
			}
		}

	default:
		panic(fmt.Sprintf("%s not int-rangeable", val.Kind()))
	}
}

func Ranger[T any](inp any) iter.Seq[T] {
	val := reflect.ValueOf(inp)

	switch val.Kind() {
	case reflect.Chan:
		return func(yield func(T) bool) {
			for {
				x, ok := val.Recv()
				if !ok {
					return
				}
				if !yield(x.Interface().(T)) {
					return
				}
			}
		}

	case reflect.Func:
		return val.Interface().(func(func(T) bool))

	default:
		panic(fmt.Sprintf("%s not T-rangeable", val.Kind()))
	}
}

func MapRanger[K comparable, V any](inp map[K]V) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range inp {
			if !yield(k, v) {
				return
			}
		}
	}
}

func EnumRanger[T any](inp any) iter.Seq2[int, T] {
	val := reflect.ValueOf(inp)

	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		return func(yield func(int, T) bool) {
			for i := 0; i < val.Len(); i++ {
				if !yield(i, val.Index(i).Interface().(T)) {
					return
				}
			}
		}

	case reflect.Func:
		return val.Interface().(func(func(int, T) bool))

	default:
		panic(fmt.Sprintf("%s not T-rangeable", val.Kind()))
	}
}

func StringRanger(inp string) iter.Seq2[int, rune] {
	return func(yield func(int, rune) bool) {
		for pos, r := range inp {
			if !yield(pos, r) {
				return
			}
		}
	}
}

func Empty[T any](func(T) bool)        {}
func Empty2[T, U any](func(T, U) bool) {}
