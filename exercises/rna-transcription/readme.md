# RNA Transcription — README

This folder contains a complete, idiomatic Go solution for the Exercism
**RNA Transcription** exercise.

## Problem Summary
Given a DNA string (e.g., `"ACGTGGTCTTAA"`), convert each nucleotide to its RNA complement:

| DNA | RNA |
|-----|-----|
| G   | C   |
| C   | G   |
| T   | A   |
| A   | U   |

Example:
```
"ACGT" → "UGCA"
```

## Files Included

### ✔ `solution.go`
Main implementation using a `switch` statement and `strings.Builder`
for optimal performance.

### ✔ `alternative.go`
A map-based implementation — slightly slower but more expressive.

### ✔ `notes.md`
Detailed explanation of the design decisions, performance notes, and Go best practices.

### ✔ `rna_transcription_test.go`
A fully commented test suite covering:
- core conversions
- multi-character strings
- empty input
- performance sanity checks

## Running Tests
From the repository root:

```bash
go test ./...
```
