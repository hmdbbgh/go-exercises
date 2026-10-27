package dna

import "errors"

// ToCountsAlternative is a variation using a manual check instead of switch.
func (d DNA) ToCountsAlternative() (Histogram, error) {
	h := Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}

	for i := 0; i < len(d); i++ {
		c := d[i]
		if c == 'A' || c == 'C' || c == 'G' || c == 'T' {
			h[c]++
		} else {
			return nil, errors.New("invalid nucleotide")
		}
	}

	return h, nil
}
