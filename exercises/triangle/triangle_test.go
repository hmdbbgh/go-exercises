package triangle

import "testing"

func TestTriangleKinds(t *testing.T) {
	tests := []struct {
		a, b, c float64
		want    Kind
	}{
		{2, 2, 2, Equ},
		{2, 2, 3, Iso},
		{3, 4, 5, Sca},
		{0, 2, 2, NaT},
		{-1, 2, 2, NaT},
		{1, 1, 3, NaT},
		{7, 3, 2, NaT},
	}

	for _, tc := range tests {
		got := KindFromSides(tc.a, tc.b, tc.c)
		if got != tc.want {
			t.Errorf("KindFromSides(%v,%v,%v) = %v; want %v",
				tc.a, tc.b, tc.c, got, tc.want)
		}
	}
}
