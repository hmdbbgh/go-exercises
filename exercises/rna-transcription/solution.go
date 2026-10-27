package strand

import "strings"

// ToRNA converts a DNA strand into its RNA complement.
// Mapping:
//   G → C
//   C → G
//   T → A
//   A → U
func ToRNA(dna string) string {
	var builder strings.Builder
	builder.Grow(len(dna)) // RNA length equals DNA length

	for i := 0; i < len(dna); i++ {
		switch dna[i] {
		case 'G':
			builder.WriteByte('C')
		case 'C':
			builder.WriteByte('G')
		case 'T':
			builder.WriteByte('A')
		case 'A':
			builder.WriteByte('U')
		}
	}

	return builder.String()
}
