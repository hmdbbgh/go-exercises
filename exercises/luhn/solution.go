package luhn

// Valid reports whether the input string is valid per the Luhn formula.
func Valid(id string) bool {
    sum := 0
    double := false
    digitCount := 0

    for i := len(id) - 1; i >= 0; i-- {
        r := id[i]

        if r == ' ' {
            continue
        }
        if r < '0' || r > '9' {
            return false
        }

        digit := int(r - '0')
        digitCount++

        if double {
            digit *= 2
            if digit > 9 {
                digit -= 9
            }
        }

        sum += digit
        double = !double
    }

    if digitCount < 2 {
        return false
    }

    return sum%10 == 0
}
