package proverb

import "fmt"

// ProverbSimple is a beginner‑friendly version.
func ProverbSimple(rhyme []string) []string {
	if len(rhyme) == 0 {
		return []string{}
	}

	result := []string{}

	for i := 0; i < len(rhyme)-1; i++ {
		line := fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
		result = append(result, line)
	}

	result = append(result,
		fmt.Sprintf("And all for the want of a %s.", rhyme[0]),
	)

	return result
}
