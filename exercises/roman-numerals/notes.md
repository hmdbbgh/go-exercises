# Roman Numerals — Notes

## Algorithm
We use a greedy algorithm:
- Iterate Roman symbols from largest to smallest.
- Repeatedly append while `input >= value`.
- Subtract and continue.

This is optimal because Roman numeral system is canonical.

## Why strings.Builder?
- String concatenation is O(n²).
- `Builder` keeps a mutable byte slice.
- Final string is constructed at once.
- Cleaner and idiomatic for Go.

## Range
Valid Roman numerals are 1–3999:
- Largest: MMMCMXCIX = 3999
- Below 1 or above 3999 → error.

## Alternative Implementation
`alternative.go` uses a `[]byte` buffer:
- Slightly faster for heavy workloads.
- Less idiomatic.
