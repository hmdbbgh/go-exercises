package proverb

import "fmt"

// Proverb returns the generated proverb using an idiomatic Go implementation.
func Proverb(rhyme []string) []string {
	n := len(rhyme)
	if n == 0 {
		return []string{}
	}

	lines := make([]string, 0, n)

	for i := 0; i < n-1; i++ {
		lines = append(lines,
			fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]),
		)
	}

	lines = append(lines,
		fmt.Sprintf("And all for the want of a %s.", rhyme[0]),
	)

	return lines
}
