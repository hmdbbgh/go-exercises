# RNA Transcription — Notes

## Overview
The RNA transcription problem requires converting a DNA string into its RNA complement:
- G → C
- C → G
- T → A
- A → U

The output string always has the same length as the input.

## Design Considerations

### 1. Using `strings.Builder`
We use `strings.Builder` for:
- Efficient string construction
- Avoiding repeated allocation
- O(n) performance

`builder.Grow(len(dna))` ensures the builder pre-allocates enough capacity,
reducing memory growth operations.

### 2. Using `switch` vs `map`
The primary solution uses `switch`, which is:
- Faster (no map lookup)
- Clear and idiomatic
- Perfect for small fixed sets of cases

The alternative implementation uses a map for readability.

### 3. Input validation
Exercism guarantees valid input (`A C G T`) so no error-handling is required.

### 4. Time & Space Complexity
- Time: O(n)
- Space: O(n)
- No intermediate allocations due to Builder growth hint.
