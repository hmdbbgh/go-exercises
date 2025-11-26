package luhn

import "unicode"

// normalize extracts digits and validates input. It ignores whitespace.
func normalize(id string) ([]int, bool) {
    digits := make([]int, 0, len(id))

    for _, r := range id {
        if unicode.IsSpace(r) {
            continue
        }
        if r < '0' || r > '9' {
            return nil, false
        }
        digits = append(digits, int(r-'0'))
    }

    if len(digits) < 2 {
        return nil, false
    }

    return digits, true
}

// ValidAlt performs Luhn validation on a normalized digit slice.
func ValidAlt(id string) bool {
    digits, ok := normalize(id)
    if !ok {
        return false
    }

    sum := 0
    double := false

    for i := len(digits) - 1; i >= 0; i-- {
        d := digits[i]
        if double {
            d *= 2
            if d > 9 {
                d -= 9
            }
        }
        sum += d
        double = !double
    }

    return sum%10 == 0
}
