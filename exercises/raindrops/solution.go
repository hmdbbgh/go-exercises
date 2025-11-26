package raindrops

import "strconv"

// Convert converts a number into its raindrop sound.
func Convert(n int) string {
    result := ""

    factors := []struct {
        divisor int
        sound   string
    }{
        {3, "Pling"},
        {5, "Plang"},
        {7, "Plong"},
    }

    for _, f := range factors {
        if n%f.divisor == 0 {
            result += f.sound
        }
    }

    if result == "" {
        return strconv.Itoa(n)
    }

    return result
}
