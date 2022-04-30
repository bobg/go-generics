# Slices - generic utility functions for operating on slices

[![Go Reference](https://pkg.go.dev/badge/github.com/bobg/slices.svg)](https://pkg.go.dev/github.com/bobg/slices)
[![Go Report Card](https://goreportcard.com/badge/github.com/bobg/slices)](https://goreportcard.com/report/github.com/bobg/slices)
[![Tests](https://github.com/bobg/slices/actions/workflows/go.yml/badge.svg)](https://github.com/bobg/slices/actions/workflows/go.yml)
[![Coverage Status](https://coveralls.io/repos/github/bobg/slices/badge.svg?branch=master)](https://coveralls.io/github/bobg/slices?branch=master)

This is slices,
a collection of generic utility functions for operating on Go slices.

This collection is useful in three ways:

- It encapsulates hard-to-remember Go idioms for inserting and removing elements to and from the middle of a slice;
- It adds the ability to index from the right end of a slice using negative integers
  (for example, Get(s, -1) is the same as s[len(s)-1]); and
- It includes `Map`, `Filter`, and a few other such functions
  for processing slice elements with callbacks
