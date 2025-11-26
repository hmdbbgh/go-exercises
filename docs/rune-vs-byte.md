# rune vs byte in Go — Understanding Strings and Unicode

Key points:
- `byte` is an alias for `uint8` and represents a raw byte
- `rune` is an alias for `int32` and represents a Unicode code point
- indexing a string (`s[i]`) yields a byte
- ranging over a string (`for _, r := range s`) yields runes

This is essential for correctly handling UTF-8 text and for deciding whether you care about bytes or characters.
