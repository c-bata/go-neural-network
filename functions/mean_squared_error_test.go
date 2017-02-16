package functions

import "testing"

func TestMeanSquaredError(t *testing.T) {
	const (
		NearEqualThreshold = 1e-8
	)

	tk := []float64{0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}

	xk := []float64{0.1, 0.05, 0.6, 0.0, 0.05, 0.1, 0.0, 0.1, 0.0, 0.0}
	if MeanSquaredError(xk, tk)-0.09750000000 > NearEqualThreshold {
		t.Fatal("MeanSquaredError() should be 0.09750000000, but doesn't match.")
	}

	xk = []float64{0.1, 0.05, 0.1, 0.0, 0.05, 0.1, 0.0, 0.6, 0.0, 0.0}
	if MeanSquaredError(xk, tk)-0.59750000000 > NearEqualThreshold {
		t.Fatal("MeanSquaredError() should be 0.59750000000, but doesn't match.")
	}
}
