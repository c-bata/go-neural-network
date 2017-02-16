package gradient

type funcType func([]float64) float64

func NumericalGradient(f funcType, x [][]float64) [][]float64 {
	var h float64 = 0.0001
	grad := make([][]float64, len(x))

	for i, a := range x {
		grad[i] = make([]float64, len(a))

		for j, b := range a {
			a[j] = b + h
			var fxh1 float64 = f(a)

			a[j] = b - h
			var fxh2 float64 = f(a)
			grad[i][j] = (fxh1 - fxh2) / (2 * h)

			a[j] = b
		}
	}
	return grad
}
