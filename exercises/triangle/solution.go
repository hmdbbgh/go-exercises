package triangle

import "math"

type Kind int

const (
    NaT Kind = iota // not a triangle
    Equ             // equilateral
    Iso             // isosceles
    Sca             // scalene
)

// KindFromSides determines the triangle kind.
func KindFromSides(a, b, c float64) Kind {
    if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
        return NaT
    }
    if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
        return NaT
    }
    if a <= 0 || b <= 0 || c <= 0 {
        return NaT
    }
    if a+b < c || a+c < b || b+c < a {
        return NaT
    }
    if a == b && b == c {
        return Equ
    }
    if a == b || a == c || b == c {
        return Iso
    }
    return Sca
}
