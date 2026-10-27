
# Roman Numerals — README

This folder contains the documentation and tests for the **Roman Numerals** Go exercise.

The goal is to convert an integer in the range **1–3999** into a valid Roman numeral string.

## Files Included

### ✔ `readme.md`  
You are reading it now.  
This file explains the purpose of the exercise, the algorithm, and the structure.

### ✔ `roman_numerals_test.go`  
A complete, fully commented test suite that validates:
- basic Roman numeral outputs  
- subtractive cases (IV, IX, XL, XC, CD, CM)  
- larger combinations  
- edge cases  
- out-of-range errors  

## Running Tests

From the repository root:

```bash
go test ./...
```

## Algorithm Overview

This exercise typically uses a **greedy subtraction strategy**:
- Maintain a table of Roman symbols ordered from largest to smallest.
- While the remaining input is greater than or equal to the symbol value,
  append the symbol and subtract its value.
- Continue until the input reaches zero.

This produces a correct and canonical Roman numeral.

## Valid Input Range

Roman numerals traditionally support values up to **3999**, because:
- The largest meaningful symbol is `M` (1000)
- The canonical maximum is:  
  ```
  3999 = MMMCMXCIX
  ```

Values outside this range are rejected.

## Notes

The actual implementation (`solution.go`) is intentionally not included here,
since you only requested the documentation + tests.

