package gradient

import "testing"

func TestGradientDecent(t *testing.T) {
	const (
		NearEqualThreshold = 0.00000001
	)

	init := [][]float64{{-3.0}, {4.0}}
	x := GradientDecent(SumSquared, init, 0.1, 100)
	if x[0][0] > NearEqualThreshold {
		t.Fatal("failure")
	}
	if x[1][0] > NearEqualThreshold {
		t.Fatal("failure")
	}
}
