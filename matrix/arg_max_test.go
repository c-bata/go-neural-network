package matrix

import "testing"

func TestArgMax(t *testing.T) {
	a := []float64{0.5, 1.0, -0.7}

	if ArgMax(a) != 1 {
		t.Errorf("ArgMax() want %#v got %#v.", 1, ArgMax(a))
	}
}

