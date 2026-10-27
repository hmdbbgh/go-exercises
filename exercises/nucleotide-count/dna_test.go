package dna

import "testing"

// TestValidCounts ensures counting works with correct DNA input.
func TestValidCounts(t *testing.T) {
	tests := []struct {
		dna  DNA
		want Histogram
	}{
		{"A", Histogram{'A': 1, 'C': 0, 'G': 0, 'T': 0}},
		{"ACGT", Histogram{'A': 1, 'C': 1, 'G': 1, 'T': 1}},
		{"AACCGGTT", Histogram{'A': 2, 'C': 2, 'G': 2, 'T': 2}},
		{"GATTACA", Histogram{'A': 3, 'C': 1, 'G': 1, 'T': 2}},
	}

	for _, tc := range tests {
		got, err := tc.dna.Counts()
		if err != nil {
			t.Fatalf("unexpected error for %q: %v", tc.dna, err)
		}

		for k := range tc.want {
			if got[k] != tc.want[k] {
				t.Errorf("Counts(%q)[%q] = %d; want %d", tc.dna, k, got[k], tc.want[k])
			}
		}
	}
}

// TestInvalid ensures invalid characters produce an error.
func TestInvalid(t *testing.T) {
	tests := []DNA{"X", "ABZ", "123", "ACGTX"}

	for _, d := range tests {
		_, err := d.Counts()
		if err == nil {
			t.Errorf("expected error for %q, got nil", d)
		}
	}
}

// TestAlternative ensures the alternative implementation matches expected behavior.
func TestAlternative(t *testing.T) {
	dna := DNA("ACGTACGT")
	want := Histogram{'A':2, 'C':2, 'G':2, 'T':2}

	got, err := dna.ToCountsAlternative()
	if err != nil {
		t.Fatalf("unexpected error in alternative implementation: %v", err)
	}

	for k := range want:
		if got[k] != want[k] {
			t.Errorf("ToCountsAlternative mismatch: [%q] = %d; want %d", k, got[k], want[k])
		}
	}
}
