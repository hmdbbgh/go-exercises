
package romannumerals

import "testing"

// TestRomanNumerals verifies that valid integers are correctly converted
// to their expected Roman numeral representations.
//
// This test suite follows the classical Roman numeral rules:
// - Use subtractive notation (IV, IX, XL, XC, CD, CM)
// - Build numerals from left to right with the largest symbols first
// - Only values 1–3999 are valid
func TestRomanNumerals(t *testing.T) {
    tests := []struct {
        input int
        want  string
    }{
        // Basic numerals
        {1, "I"},
        {2, "II"},
        {3, "III"},
        {4, "IV"},
        {5, "V"},
        {6, "VI"},
        {9, "IX"},

        // Multi‑symbol numerals
        {27, "XXVII"},
        {48, "XLVIII"},
        {49, "XLIX"},
        {59, "LIX"},
        {93, "XCIII"},
        {141, "CXLI"},
        {163, "CLXIII"},

        // Larger values
        {402, "CDII"},
        {575, "DLXXV"},
        {911, "CMXI"},
        {1024, "MXXIV"},
        {3000, "MMM"},
        {3999, "MMMCMXCIX"},
    }

    for _, tc := range tests {
        got, err := ToRomanNumeral(tc.input)
        if err != nil {
            t.Fatalf("unexpected error for %d: %v", tc.input, err)
        }
        if got != tc.want {
            t.Errorf("ToRomanNumeral(%d) = %q; want %q", tc.input, got, tc.want)
        }
    }
}

// TestRomanNumeralsOutOfRange verifies that invalid inputs (zero,
// negative numbers, and numbers larger than 3999) return an error.
func TestRomanNumeralsOutOfRange(t *testing.T) {
    tests := []int{-1, 0, 4000, 5000}

    for _, input := range tests {
        _, err := ToRomanNumeral(input)
        if err == nil {
            t.Errorf("expected error for input %d, got nil", input)
        }
    }
}
