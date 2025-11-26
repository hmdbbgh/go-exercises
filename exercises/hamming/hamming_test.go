package hamming

import "testing"

func TestDistance(t *testing.T) {
    tests := []struct {
        a, b string
        want int
        ok   bool
    }{
        {"GGACTGA", "GGACTGA", 0, true},
        {"GGACTGA", "GGACTEA", 1, true},
    }

    for _, tc := range tests {
        got, err := Distance(tc.a, tc.b)
        if (err == nil) != tc.ok {
            t.Fatalf("Distance(%q,%q): unexpected error: %v", tc.a, tc.b, err)
        }
        if err == nil && got != tc.want {
            t.Fatalf("Distance(%q,%q) = %d, want %d", tc.a, tc.b, got, tc.want)
        }
    }
}
