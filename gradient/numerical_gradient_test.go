package gradient

import (
	"testing"
	"math"
)


func SumSquared(x []float64) float64 {
	var sum float64 = 0
	for _, a := range x {
		sum += math.Pow(a, 2)
	}
	return sum
}


func TestNumericalGradient(t *testing.T) {
	const (
		NearEqualThreshold = 0.00000001
	)

	x := [][]float64{{3.0}, {4.0}}
	a := NumericalGradient(SumSquared, x)
	if math.Abs(a[0][0] - 6.0) > NearEqualThreshold || math.Abs(a[1][0] - 8.0) > NearEqualThreshold {
		t.Fatal("NumericalGradient([][]float64{{3.0}, {4.0}}) should be []float64{{6.0}, {8.0}}, but doesn't match.")
	}

	x = [][]float64{{0.0}, {2.0}}
	a = NumericalGradient(SumSquared, x)
	if math.Abs(a[0][0] - 0.0) > NearEqualThreshold || math.Abs(a[1][0] - 4.0) > NearEqualThreshold {
		t.Fatal("NumericalGradient([][]float64{{0.0}, {2.0}}) should be []float64{{0.0}, {4.0}}, but doesn't match.")
	}

	x = [][]float64{{3.0}, {0.0}}
	a = NumericalGradient(SumSquared, x)
	if math.Abs(a[0][0] - 6.0) > NearEqualThreshold || math.Abs(a[1][0] - 0.0) > NearEqualThreshold {
		t.Fatal("NumericalGradient([][]float64{{3.0}, {0.0}}) should be []float64{{6.0}, {0.0}}, but doesn't match.")
	}
}
