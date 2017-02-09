package functions

import "testing"

func TestCrossEntropyError(t *testing.T) {
	const (
		NearEqualThreshold = 1e-8
	)

	tk := []float64{0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}

	xk := []float64{0.1, 0.05, 0.6, 0.0, 0.05, 0.1, 0.0, 0.1, 0.0, 0.0}
	if CrossEntropyError(xk, tk) - 0.510825457 > NearEqualThreshold {
		t.Fatal("CrossEntropyError() should be 0.510825457, but doesn't match. %v", CrossEntropyError(xk, tk) - 0.510825457)
	}


	xk = []float64{0.1, 0.05, 0.1, 0.0, 0.05, 0.1, 0.0, 0.6, 0.0, 0.0}
	if MeanSquaredError(xk, tk) - 2.3025840929 > NearEqualThreshold {
		t.Fatal("CrossEntropyError() should be 2.3025840929, but doesn't match.")
	}
}

