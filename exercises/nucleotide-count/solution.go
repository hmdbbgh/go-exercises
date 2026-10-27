package dna

import "errors"

// Histogram counts occurrences of nucleotides.
type Histogram map[byte]int

// DNA represents a DNA strand.
type DNA string

// Counts computes the frequency of each nucleotide in the DNA string.
// Returns an error if an invalid nucleotide is encountered.
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}

	for i := 0; i < len(d); i++ {
		switch d[i] {
		case 'A', 'C', 'G', 'T':
			h[d[i]]++
		default:
			return nil, errors.New("invalid nucleotide")
		}
	}

	return h, nil
}
