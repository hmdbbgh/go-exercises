package proverb

import "fmt"

const (
	causeLine = "For want of a %s the %s was lost."
	finalLine = "And all for the want of a %s."
)

// ProverbOptimized is a performance‑focused implementation with no slice growth.
func ProverbOptimized(rhyme []string) []string {
	n := len(rhyme)
	if n == 0 {
		return []string{}
	}

	out := make([]string, n)

	for i := 0; i < n-1; i++ {
		out[i] = fmt.Sprintf(causeLine, rhyme[i], rhyme[i+1])
	}

	out[n-1] = fmt.Sprintf(finalLine, rhyme[0])
	return out
}
