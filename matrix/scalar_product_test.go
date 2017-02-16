package matrix

import "testing"

func TestScalar1d(t *testing.T) {
	x := Scalar1d([]float64{1, 2}, 2)
	if x[0] != 2 {
		t.Fatal("Scalar1d should be return 2")
	}
	if x[1] != 4 {
		t.Fatal("Scalar1d should be return 2")
	}
}

func TestScalar2d(t *testing.T) {
	x := Scalar2d([][]float64{{1, 2}, {3, 4}}, 3)
	if x[1][0] != 9 {
		t.Fatal("Scalar1d should be return 9")
	}
	if x[1][1] != 12 {
		t.Fatal("Scalar1d should be return 12")
	}
}
