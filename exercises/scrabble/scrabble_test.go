package scrabble

import "testing"

func TestScore(t *testing.T) {
    tests := []struct {
        word string
        want int
    }{
        {"a", 1},
        {"f", 4},
        {"street", 6},
        {"quirky", 22},
        {"Cabbage", 14},
        {"", 0},
    }

    for _, tc := range tests {
        if got := Score(tc.word); got != tc.want {
            t.Fatalf("Score(%q) = %d, want %d", tc.word, got, tc.want)
        }
    }
}
