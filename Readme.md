# Go-generics - Generic slice, map, set, iterator, and goroutine utilities for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/bobg/go-generics.svg)](https://pkg.go.dev/github.com/bobg/go-generics)
[![Go Report Card](https://goreportcard.com/badge/github.com/bobg/go-generics)](https://goreportcard.com/report/github.com/bobg/go-generics)
[![Tests](https://github.com/bobg/go-generics/actions/workflows/go.yml/badge.svg)](https://github.com/bobg/go-generics/actions/workflows/go.yml)
[![Coverage Status](https://coveralls.io/repos/github/bobg/go-generics/badge.svg?branch=master)](https://coveralls.io/github/bobg/go-generics?branch=master)

This is go-generics,
a collection of typesafe generic utilities
for slices, maps, sets, iterators, and goroutine patterns in Go.

# Slices

The `slices` package is useful in three ways:

- It encapsulates hard-to-remember Go idioms for inserting and removing elements to and from the middle of a slice;
- It adds the ability to index from the right end of a slice using negative integers
  (for example, Get(s, -1) is the same as s[len(s)-1]); and
- It includes `Map`, `Filter`, and a few other such functions
  for processing slice elements with callbacks.

It also includes combinatorial operations:
`Permutations`, `Combinations`, and `CombinationsWithReplacement`.

The `slices` package is a drop-in replacement for `golang.org/x/exp/slices`
(which will become part of the stdlib [in Go 1.21](https://github.com/golang/go/issues/57433))
with one difference:
the index values passed to `Insert`, `Delete`, and `Replace`
may be negative for counting backward from the end of the slice.

# Maps

The `maps` package has a few convenience functions
for duplicating, inverting, constructing, and iterating over maps,
as well as for testing their equality.

The `maps` package is a drop-in replacement for `golang.org/x/exp/maps`,
(which will become part of the stdlib [in Go 1.21](https://github.com/golang/go/issues/57436)).

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
