package hamming

import "errors"

// Distance returns the Hamming distance between a and b.
// It returns an error if the strings are of unequal length.
func Distance(a, b string) (int, error) {
    if len(a) != len(b) {
        return 0, errors.New("strings must be of equal length")
    }

    distance := 0
    for i := 0; i < len(a); i++ {
        if a[i] != b[i] {
            distance++
        }
    }

    return distance, nil
}
