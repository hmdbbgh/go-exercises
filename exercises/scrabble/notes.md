# Scrabble – Notes

- Primary implementation uses a `map[rune]int` for clarity.
- Alternative uses a `[26]int` array indexed by uppercase letters for speed.
- Demonstrates:
  - rune iteration
  - Unicode normalization (`unicode.ToUpper`)
  - map lookup with the `value, ok := map[key]` idiom
