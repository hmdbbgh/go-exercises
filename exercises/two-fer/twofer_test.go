package twofer

import "testing"

func TestShareWith(t *testing.T) {
    tests := []struct {
        name string
        want string
    }{
        {"Alice", "One for Alice, one for me."},
        {"Bob", "One for Bob, one for me."},
        {"", "One for you, one for me."},
    }

    for _, tc := range tests {
        if got := ShareWith(tc.name); got != tc.want {
            t.Fatalf("ShareWith(%q) = %q, want %q", tc.name, got, tc.want)
        }
    }
}
