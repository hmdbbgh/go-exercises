package raindrops

import "testing"

func TestConvert(t *testing.T) {
    tests := []struct {
        in   int
        want string
    }{
        {1, "1"},
        {28, "Plong"},
        {30, "PlingPlang"},
        {34, "34"},
    }

    for _, tc := range tests {
        if got := Convert(tc.in); got != tc.want {
            t.Fatalf("Convert(%d) = %q, want %q", tc.in, got, tc.want)
        }
    }
}
