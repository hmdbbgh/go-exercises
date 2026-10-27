# strings.Builder — Efficient String Builder

Go strings are immutable.
Doing this:

```go
s += "A"
```

creates a new string every time → slow.

`strings.Builder` solves this:
- Uses internal growable byte slice.
- Writes are O(1) amortized.
- Final `String()` creates one string.

## Best Practices
- Do not reuse after calling `.String()`.
- Use `Grow(n)` when approximate size is known.
- Great for loops, encoders, Roman numeral conversion.

