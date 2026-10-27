package strand

import "testing"

// TestBasic verifies the core DNA → RNA mapping logic.
func TestBasic(t *testing.T) {
	tests := []struct {
		dna  string
		want string
	}{
		{"G", "C"},
		{"C", "G"},
		{"T", "A"},
		{"A", "U"},
	}

	for _, tc := range tests {
		got := ToRNA(tc.dna)
		if got != tc.want {
			t.Errorf("ToRNA(%q) = %q; want %q", tc.dna, got, tc.want)
		}
	}
}

// TestSequences ensures multi-character transcription works correctly.
func TestSequences(t *testing.T) {
	tests := []struct {
		dna  string
		want string
	}{
		{"ACGT", "UGCA"},
		{"GATTACA", "CUAAUGU"},
		{"CCCC", "GGGG"},
		{"TTTT", "AAAA"},
	}

	for _, tc := range tests {
		got := ToRNA(tc.dna)
		if got != tc.want {
			t.Errorf("ToRNA(%q) = %q; want %q", tc.dna, got, tc.want)
		}
	}
}

// TestEmptyInput validates the behavior on empty strings.
func TestEmptyInput(t *testing.T) {
	got := ToRNA("")
	if got != "" {
		t.Errorf("ToRNA(\"\") = %q; want \"\"", got)
	}
}

// TestMapImplementation cross-checks alternative implementation.
func TestMapImplementation(t *testing.T) {
	dna := "ACGTACGT"
	want := "UGCAUGCA"

	got := ToRNAMap(dna)
	if got != want {
		t.Errorf("ToRNAMap(%q) = %q; want %q", dna, got, want)
	}
}
