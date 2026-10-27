package triangle

import "math"

type KindAlt int

const (
    NaTAlt KindAlt = iota
    EquAlt
    IsoAlt
    ScaAlt
)

const eps = 1e-9

func equal(x, y float64) bool {
    return math.Abs(x-y) < eps
}

func KindAltFromSides(a, b, c float64) KindAlt {
    if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
        return NaTAlt
    }
    if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
        return NaTAlt
    }
    if a <= 0 || b <= 0 || c <= 0 {
        return NaTAlt
    }
    if a+b < c-eps || a+c < b-eps || b+c < a-eps {
        return NaTAlt
    }
    if equal(a, b) && equal(b, c) {
        return EquAlt
    }
    if equal(a, b) || equal(a, c) || equal(b, c) {
        return IsoAlt
    }
    return ScaAlt
}
