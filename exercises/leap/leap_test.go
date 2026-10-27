package leap

import "testing"

func TestLeapYears(t *testing.T) {
	tests := []struct {
		year int
		want bool
	}{
		{1996, true},
		{2000, true},
		{2020, true},
		{1900, false},
		{2100, false},
		{2023, false},
	}

	for _, tc := range tests {
		got := IsLeapYear(tc.year)
		if got != tc.want {
			t.Errorf("IsLeapYear(%d) = %v; want %v", tc.year, got, tc.want)
		}
	}
}

func TestReadableVersion(t *testing.T) {
	for _, y := range []int{1996, 2000, 2020, 1900, 2100, 2023} {
		if IsLeapYear(y) != IsLeapYearReadable(y) {
			t.Errorf("Mismatch between versions for year %d", y)
		}
	}
}

// Benchmark for performance
func BenchmarkLeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsLeapYear(2024)
	}
}
