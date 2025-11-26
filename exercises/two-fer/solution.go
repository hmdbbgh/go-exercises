package twofer

import "fmt"

// ShareWith returns "One for <name>, one for me.".
// If name is empty, it defaults to "you".
func ShareWith(name string) string {
    if name == "" {
        name = "you"
    }
    return fmt.Sprintf("One for %s, one for me.", name)
}
