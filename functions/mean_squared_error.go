package functions

import "math"

func MeanSquaredError(x []float64, t []float64) float64 {
	var sum float64 = 0

	for i := 0; i < len(x); i++ {
		sum += math.Pow(x[i]-t[i], 2)
	}
	return sum / 2
}
