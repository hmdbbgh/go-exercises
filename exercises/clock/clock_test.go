package clock

import "testing"

func TestClock(t *testing.T) {
    tests := []struct {
        h, m int
        add  int
        want string
    }{
        {23, 59, 2, "00:01"},
        {25, 0, 0, "01:00"},
        {-1, 15, 0, "23:15"},
    }

    for _, tc := range tests {
        c := New(tc.h, tc.m).Add(tc.add)
        if got := c.String(); got != tc.want {
            t.Fatalf("New(%d,%d).Add(%d) = %s, want %s",
                tc.h, tc.m, tc.add, got, tc.want)
        }
    }
}
