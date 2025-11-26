package twofer

import "fmt"

// defaultString returns fallback if s is empty.
func defaultString(s, fallback string) string {
    if s == "" {
        return fallback
    }
    return s
}

// ShareWithAlt is an alternative implementation using a helper.
func ShareWithAlt(name string) string {
    name = defaultString(name, "you")
    return fmt.Sprintf("One for %s, one for me.", name)
}
