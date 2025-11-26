package isogram

import "testing"

func TestIsIsogram(t *testing.T) {
    tests := []struct {
        word string
        want bool
    }{
        {"background", true},
        {"six-year-old", true},
        {"isograms", false},
    }

    for _, tc := range tests {
        if got := IsIsogram(tc.word); got != tc.want {
            t.Fatalf("IsIsogram(%q) = %t, want %t", tc.word, got, tc.want)
        }
    }
}
