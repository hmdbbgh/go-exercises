package diffsquares

import "testing"

func TestDifference(t *testing.T) {
    tests := []struct {
        n    int
        want int
    }{
        {5, 170},
        {10, 2640},
    }

    for _, tc := range tests {
        if got := Difference(tc.n); got != tc.want {
            t.Fatalf("Difference(%d) = %d, want %d", tc.n, got, tc.want)
        }
    }
}
