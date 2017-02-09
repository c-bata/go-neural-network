package matrix

import "testing"

func TestSubtract1d(t *testing.T) {
	x := Subtract1d([]float64{1, 2}, []float64{2, 2})
	if x[0] != -1 {
		t.Fatal("Subtract1d should be return -1")
	}
	if x[1] != 0 {
		t.Fatal("Subtract1d should be return 0")
	}
}


func TestSubtract2d(t *testing.T) {
	x := Subtract2d([][]float64{{1, 2}, {3, 4}}, [][]float64{{1, 2}, {3, 4}})
	if x[0][0] != 0 {
		t.Fatal("Subtract2d should be return 0")
	}
}
