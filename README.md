# Happy Fun Books

A tiny Go library and CLI for a hard-coded book catalog used for learning Go patterns and simple CLI programs.

Purpose
- Demonstrates a small package (`books`) with a `Book` type and a simple in-memory `Catalog`.
- Provides two CLI commands under `cmd/` to list the catalog and find a book by ID.

Repository layout

- `books.go` — Core package implementing `Book`, `Catalog`, getters, and a small hard-coded catalog.
- `books_test.go` — Unit tests for the `books` package.
- `cmd/list/main.go` — CLI to list all books in stock.
- `cmd/find/main.go` — CLI to find a book by its ID.

Quick start

Install Go (project targets modern Go 1.x) then run the CLIs directly with `go run`:

```bash
go run ./cmd/list
go run ./cmd/find abc
```

Run tests

```bash
go test ./...
```

Working with the catalog

The catalog is returned by `GetCatalog()` in `books.go`. To add or change books edit that function and update tests as needed.

Conventions

- Functions in `books` return `(value, bool)` where appropriate (e.g., `GetBook(ID string) (Book, bool)`).
- String formatting for display is implemented on the `Book` type.

Notes

This repository is intentionally minimal and uses only the Go standard library. It's intended for learning and demonstration purposes.
