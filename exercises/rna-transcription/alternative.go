package strand

import "strings"

// ToRNAMap is an alternative implementation using a lookup map.
// Clean, expressive, but slightly slower due to map access overhead.
func ToRNAMap(dna string) string {
	var builder strings.Builder
	builder.Grow(len(dna))

	m := map[byte]byte{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}

	for i := 0; i < len(dna); i++ {
		if v, ok := m[dna[i]]; ok {
			builder.WriteByte(v)
		}
	}

	return builder.String()
}
