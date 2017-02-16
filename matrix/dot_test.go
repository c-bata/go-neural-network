package matrix

import "testing"

func TestDot(t *testing.T) {
	a := [][]float64{{1, 2, 3}, {4, 5, 6}}
	b := [][]float64{{1, 2}, {3, 4}, {5, 6}}
	c := Dot(a, b)

	if len(c) != 2 || len(c[0]) != 2 {
		t.Fatal("Dot() should return 2x2 arrays.")
	}

	if c[0][0] != 1*1+2*3+3*5 {
		t.Fatal("c[0][0] should be 22(= 1*1 + 2*3 + 3*5).")
	}
}
