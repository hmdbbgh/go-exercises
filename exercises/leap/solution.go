package leap

// IsLeapYear returns true if the given year is a leap year.
// Rule:
//   - divisible by 4 AND
//   - (not divisible by 100 OR divisible by 400)
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
