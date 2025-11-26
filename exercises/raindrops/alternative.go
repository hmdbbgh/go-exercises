package raindrops

import "strconv"

// ConvertIf is an alternative implementation using explicit if-chains.
func ConvertIf(n int) string {
    result := ""

    if n%3 == 0 {
        result += "Pling"
    }
    if n%5 == 0 {
        result += "Plang"
    }
    if n%7 == 0 {
        result += "Plong"
    }

    if result == "" {
        return strconv.Itoa(n)
    }

    return result
}
