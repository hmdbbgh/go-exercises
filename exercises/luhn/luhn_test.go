package luhn

import "testing"

func TestValid(t *testing.T) {
    valid := []string{
        "055 444 285",
        "059",
    }
    invalid := []string{
        "055 444 286",
        "abc",
        "0",
    }

    for _, s := range valid {
        if !Valid(s) {
            t.Fatalf("Valid(%q) = false, want true", s)
        }
    }
    for _, s := range invalid {
        if Valid(s) {
            t.Fatalf("Valid(%q) = true, want false", s)
        }
    }
}
