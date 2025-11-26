# Isogram – Notes

- Uses `map[rune]bool` as a natural "set".
- Iterates over runes with `for _, r := range`.
- Alternative uses a fixed array for ASCII only, which is more efficient but less general.
