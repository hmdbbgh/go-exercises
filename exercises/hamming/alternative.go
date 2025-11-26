package hamming

import (
    "errors"
    "unicode/utf8"
)

// DistanceRunes is an alternative implementation that works at the rune level.
func DistanceRunes(a, b string) (int, error) {
    if utf8.RuneCountInString(a) != utf8.RuneCountInString(b) {
        return 0, errors.New("strings must be of equal rune length")
    }

    ra := []rune(a)
    rb := []rune(b)

    distance := 0
    for i := range ra {
        if ra[i] != rb[i] {
            distance++
        }
    }

    return distance, nil
}
