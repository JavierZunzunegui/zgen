# Helpers for Generic Go

[![Go Version](https://img.shields.io/github/go-mod/go-version/JavierZunzunegui/zgen)](./go.mod)
[![GoDoc](https://godoc.org/github.com/JavierZunzunegui/zgen?status.svg)](https://pkg.go.dev/github.com/JavierZunzunegui/zgen)
![Build Status](https://github.com/JavierZunzunegui/zgen/actions/workflows/go.yml/badge.svg)
[![Go report](https://goreportcard.com/badge/github.com/JavierZunzunegui/zgen)](https://goreportcard.com/report/github.com/JavierZunzunegui/zgen)
[![License](https://img.shields.io/github/license/JavierZunzunegui/zerrors)](./LICENSE)

Experimental packages with common helpers for generic Go (type parameters), and particularly for using iterators.

## Contents

The primary package is [ziter](https://pkg.go.dev/github.com/JavierZunzunegui/zgen/ziter). It provides many functions to work with iterators. They all use type parameters.

The root package [zgen](https://pkg.go.dev/github.com/JavierZunzunegui/zgen) provides a few other generic helpers for operations unrelated to iterators, but their value is small.

THIS REPO IS EXPERIMENTAL and subject to breaking changes. I use it to explore how much use of generics and iterators is actually practical (primarily from a readability perspective), and at what point it actually makes the code less understandable. Unconstrained use of ziter+zgen can actually result in very odd (and compact) go code.

## Relevant Tutorials

- New to iter (go 1.23+)? https://pkg.go.dev/iter
- New to generics (go 1.18+)? https://go.dev/doc/tutorial/generics