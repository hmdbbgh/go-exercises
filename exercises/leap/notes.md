# Leap — Technical Notes

## 1. Leap Year Rules
A year is a leap year if:

1. It is divisible by 4  
2. If divisible by 100, it must *also* be divisible by 400

In boolean form:

```
year % 4 == 0 && (year % 100 != 0 || year % 400 == 0)
```

This follows the Gregorian calendar rules.

---

## 2. Implementation Approaches

### 2.1 Compact Boolean Expression (solution.go)
This is the idiomatic one-liner:

```go
return year%4 == 0 && (year%100 != 0 || year%400 == 0)
```

Pros:
- Very fast
- No branching
- Zero allocations
- Perfectly readable for Go developers

### 2.2 Expanded Readable Version (alternative.go)

```go
if year%400 == 0 { return true }
if year%100 == 0 { return false }
return year%4 == 0
```

Pros:
- Extremely readable
- Matches the rule order as described in natural language

Both versions produce identical behavior.

---

## 3. Performance (Benchmarking)

Benchmarks measure **speed**, **allocations**, and **CPU behavior**, not correctness.

Example benchmark:

```go
func BenchmarkLeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsLeapYear(2024)
	}
}
```

Example output:

```
BenchmarkLeap-8   1000000000   0.25 ns/op   0 allocs/op
```

Meaning:
- Function was run 1 billion times
- Average cost = 0.25 nanoseconds per call
- Zero memory allocations
- Extremely efficient (as expected)

---

## 4. Complexity
Both versions are:

- Time: O(1)
- Space: O(1)
- No heap allocations
- Single-digit CPU instructions

---

## 5. Summary
The leap year function is simple but a perfect example of:

- Clear boolean logic
- Branching vs no-branching implementation
- Benchmark analysis
- Clean and idiomatic Go design
