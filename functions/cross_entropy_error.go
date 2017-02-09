package functions

import "math"

func CrossEntropyError(x []float64, t []float64) float64 {
	var d float64 = 1e-7  // to avoid log(0)
	var sum float64 = 0

	for i := 0; i < len(x); i++ {
		sum += t[i] * math.Log(x[i] + d)
	}
	return -sum
}

