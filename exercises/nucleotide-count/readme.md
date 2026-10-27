# Nucleotide Count — README

This folder contains a complete Go solution for the Exercism
**Nucleotide Count** exercise.

## Problem Summary
Count occurrences of nucleotides in a DNA string.

Valid nucleotides:
- A
- C
- G
- T

Invalid characters must raise an error.

Example:
```
"ACGTACGAA" → A:4 C:2 G:2 T:1
```

## Files Included

### ✔ `solution.go`
Primary implementation using `switch`.

### ✔ `alternative.go`
Alternative counting style using manual conditional checks.

### ✔ `notes.md`
Design and performance considerations.

### ✔ `dna_test.go`
Unit tests with validation for valid and invalid input.

## Run Tests
```
go test ./...
```
