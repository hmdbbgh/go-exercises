Idiomatic Go exercises with clean solutions, alternative implementations, tests, and deep learning notes.

# Go Exercises – Professional Track

This repository contains senior-level, idiomatic Go solutions for a selection of Exercism Go track exercises.

The goals:

- Practice writing **clean, idiomatic Go**
- Explore **multiple implementations** for the same problem (trade-offs: clarity vs performance)
- Capture **design notes** and **language concepts** in `docs/`
- Provide a repo that can be used as a **personal reference** or shared with others

## Layout

- `exercises/<exercise>/solution.go` – primary, idiomatic solution
- `exercises/<exercise>/alternative.go` – optional alternative implementation(s)
- `exercises/<exercise>/notes.md` – design notes, reasoning, Go vs Python comparisons
- `exercises/<exercise>/*_test.go` – tests
- `docs/*.md` – deeper conceptual notes (pointers, slices, maps, iota, enums, etc.)
- `Makefile` – helpers for formatting and testing

## Running Tests

From the repo root:

```bash
go test ./...
```

To format:

```bash
go fmt ./...
```
