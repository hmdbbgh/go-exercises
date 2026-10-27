# Triangle — Technical Notes

## 1. Custom Type: `Kind`
The exercise defines an enumeration-like type:

```go
type Kind int
```

This creates a **new logical type**, backed by an integer.  
It provides:

- Type safety  
- Cleaner API  
- Prevention of invalid values  
- More idiomatic Go designs

Using a custom type prevents returning arbitrary strings like `"triangle"` or `"weird"`.

---

## 2. Why We Use `iota` for Enum Values

Go does not have a built‑in `enum` keyword.  
The idiomatic approach is to use a `const` block with `iota`:

```go
const (
    NaT Kind = iota
    Equ
    Iso
    Sca
)
```

### How it works:
- `iota` starts at **0** in each `const` block.
- Each new line increments it by **1**.
- Go automatically repeats the previous expression when omitted.

This block is equivalent to:

```
NaT = 0
Equ = 1
Iso = 2
Sca = 3
```

### Why not assign explicit numbers?
- Less typing  
- Less chance of mistakes  
- Automatically synchronized  
- Readable and standard practice  

This pattern is used widely across Go projects.

---

## 3. Validity Rules

### 3.1 Side positivity
A triangle requires strictly positive sides.

### 3.2 Triangle inequality
A triangle is valid if:

```
a + b >= c
a + c >= b
b + c >= a
```

The Exercism version accepts *degenerate* triangles (`>=` is allowed).

---

## 4. Float Comparison Notes

Even though floating‑point comparisons are usually tricky, the Exercism test inputs use simple values, so:

```
a == b
```

is safe.

For a production‑grade version, an epsilon is introduced.  
This is implemented in `alternative.go`.

---

## 5. Classification Order

Checking order is important:

1. Equilateral (`a == b && b == c`)
2. Isosceles (`a == b || a == c || b == c`)
3. Scalene (fallback)

Failing to check equilateral first will misclassify equilateral as isosceles.

---

## 6. Defensive Programming

We guard against:

- `NaN`  
- `Inf`  
- Zero/negative values  

These are not required by the basic exercise but reflect senior‑level safety.

---

## 7. Complexity

- Time: **O(1)**  
- Memory: **O(1)**  
- No allocations  

A straightforward and efficient implementation.

---

## 8. Summary

This exercise teaches:

- Custom types  
- `iota`‑based enumerations  
- Triangle inequality rules  
- Float safety  
- Clean branching  
- Idiomatic Go API design

This file serves as a deeper reference for your repository.
