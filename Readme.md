# Go-generics - Generic slice, map, set, iterator, and goroutine utilities for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/bobg/go-generics/v4.svg)](https://pkg.go.dev/github.com/bobg/go-generics/v4)
[![Go Report Card](https://goreportcard.com/badge/github.com/bobg/go-generics/v4)](https://goreportcard.com/report/github.com/bobg/go-generics/v4)
[![Tests](https://github.com/bobg/go-generics/actions/workflows/go.yml/badge.svg)](https://github.com/bobg/go-generics/actions/workflows/go.yml)
[![Coverage Status](https://coveralls.io/repos/github/bobg/go-generics/badge.svg?branch=master)](https://coveralls.io/github/bobg/go-generics?branch=master)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)

This is go-generics,
a collection of typesafe generic utilities
for slices, sets, and goroutine patterns in Go.

# Compatibility note

This is version 4 of this library,
for the release of Go 1.23.

Earlier versions of this library included a package,
`iter`,
that defined an iterator type over several types of containers,
and functions for operating with iterators.
However, Go 1.23 defines its own, better iterator mechanism
via the new “range over function” language feature,
plus [a new standard-library package](https://pkg.go.dev/iter) also called `iter`.
This version of the go-generics library therefore does away with its `iter` package.
The handy functions that `iter` contained for working with iterators
(`Filter`, `Map`, `FirstN`, and many more)
can now be found in the [github.com/bobg/seqs](https://pkg.go.dev/github.com/bobg/seqs) library,
adapted for Go 1.23 iterators.

(This version of go-generics might have kept `iter` as a drop-in replacement for the standard-library package,
but was unable because the standard library defines two types,
`iter.Seq[K]` and `iter.Seq2[K, V]`,
that go-generics would have had to reference using type aliases;
but Go type aliases [do not yet permit type parameters](https://github.com/golang/go/issues/46477#issuecomment-2101270785).)

Earlier versions of this library included combinatorial operations in the `slices` package.
Those have now been moved to their own library,
[github.com/bobg/combo](https://pkg.go.dev/github.com/bobg/combo).

Earlier versions of this library included a `maps` package,
which was a drop-in replacement for the stdlib `maps`
(added in Go 1.21)
plus a few convenience functions.
With the advent of Go 1.23 iterators,
those few convenience functions are mostly redundant
(and a couple of them − `Keys` and `Values` − conflict with new functions in the standard library),
so `maps` has been removed.

Earlier versions of this library included a `Find` method on the `set.Of[T]` type,
for finding some element in the set that satisfies a given predicate.
This method has been removed in favor of composing operations with functions from [github.com/bobg/seqs](https://pkg.go.dev/github.com/bobg/seqs).
For example, `s.Find(pred)` can now be written as `seqs.First(seqs.Filter(s.All(), pred))`.

# Slices

The `slices` package is useful in three ways:

- It encapsulates hard-to-remember Go idioms for inserting and removing elements to and from the middle of a slice;
- It adds the ability to index from the right end of a slice using negative integers
  (for example, Get(s, -1) is the same as s[len(s)-1]); and
- It includes `Map`, `Filter`, and a few other such functions
  for processing slice elements with callbacks.

The `slices` package is a drop-in replacement
for the `slices` package
added to the Go stdlib
in [Go 1.21](https://go.dev/doc/go1.21#slices).
There is one difference:
this version of `slices`
allows the index value passed to `Insert`, `Delete`, and `Replace`
to be negative for counting backward from the end of the slice.

# Set

The `set` package implements the usual collection of functions for sets:
`Intersect`, `Union`, `Diff`, etc.,
as well as member functions for adding and removing items,
checking for the presence of items,
and iterating over items.

# Iter

The `iter` package implements efficient, typesafe iterators
that can convert to and from Go slices, maps, and channels,
and produce iterators over the values from function calls and goroutines.
There is also an iterator over the results of a SQL query;
the usual collection of functions on iterators
(`Filter`, `Map`, `Concat`, `Accum`, etc.).

# Parallel

The `parallel` package contains functions for coordinating parallel workers:

- `Consumers` manages a set of N workers consuming a stream of values produced by the caller.
- `Producers` manages a set of N workers producing a stream of values consumed by the caller.
- `Values` concurrently produces a set of N values.
- `Pool` manages access to a pool of concurrent workers.
- `Protect` manages concurrent access to a protected data value.
