package isogram

import "unicode"

// IsIsogramASCII is an alternative implementation using a fixed-size array.
func IsIsogramASCII(word string) bool {
    var seen [26]bool

    for _, r := range word {
        if r == '-' || r == ' ' {
            continue
        }
        if !unicode.IsLetter(r) {
            continue
        }

        lower := unicode.ToLower(r)
        idx := int(lower - 'a')
        if idx < 0 || idx >= len(seen) {
            return false
        }

        if seen[idx] {
            return false
        }
        seen[idx] = true
    }

    return true
}
