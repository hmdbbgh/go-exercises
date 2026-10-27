# Nucleotide Count — Notes

## Overview
The task is to count nucleotides in a DNA string.
Valid nucleotides: `A`, `C`, `G`, `T`.

## Design Decisions

### Custom Types
Using types:
```go
type Histogram map[byte]int
type DNA string
```
gives type-safety and more expressive APIs.

### Initialization
The histogram is initialized with all nucleotides at zero,
ensuring output always contains all keys.

### Switch-based Counting
`switch` is preferred for:
- speed (no hashing)
- clarity
- small fixed set of cases

### Error Handling
Invalid nucleotides must produce an error.

### Complexity
- Time: O(n)
- Space: O(1)
Only a single pass with constant memory.
