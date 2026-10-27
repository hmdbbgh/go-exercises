package romannumerals

import "errors"

// Faster alternative using raw []byte buffer.
var romanTable = []struct {
    Value  int
    Symbol string
}{
    {1000, "M"},
    {900, "CM"},
    {500, "D"},
    {400, "CD"},
    {100, "C"},
    {90, "XC"},
    {50, "L"},
    {40, "XL"},
    {10, "X"},
    {9, "IX"},
    {5, "V"},
    {4, "IV"},
    {1, "I"},
}

func ToRomanAlternative(input int) (string, error) {
    if input <= 0 || input > 3999 {
        return "", errors.New("input must be between 1 and 3999")
    }

    buf := make([]byte, 0, 32)

    for _, n := range romanTable {
        for input >= n.Value {
            buf = append(buf, n.Symbol...)
            input -= n.Value
        }
    }

    return string(buf), nil
}
