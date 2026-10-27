package leap

// IsLeapYearReadable is a more explicit and human-readable version.
func IsLeapYearReadable(year int) bool {
	if year%400 == 0 {
		return true
	}
	if year%100 == 0 {
		return false
	}
	return year%4 == 0
}
