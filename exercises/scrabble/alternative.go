package scrabble

import "unicode"

var scoreTable = [26]int{
    1,  // A
    3,  // B
    3,  // C
    2,  // D
    1,  // E
    4,  // F
    2,  // G
    4,  // H
    1,  // I
    8,  // J
    5,  // K
    1,  // L
    3,  // M
    1,  // N
    1,  // O
    3,  // P
    10, // Q
    1,  // R
    1,  // S
    1,  // T
    1,  // U
    4,  // V
    4,  // W
    8,  // X
    4,  // Y
    10, // Z
}

// ScoreAlt is a more low-level, array-based implementation.
func ScoreAlt(word string) int {
    total := 0

    for _, r := range word {
        upper := unicode.ToUpper(r)
        if upper < 'A' || upper > 'Z' {
            continue
        }
        total += scoreTable[upper-'A']
    }

    return total
}
