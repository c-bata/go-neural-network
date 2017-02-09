package gradient

import "github.com/c-bata/go-neural-network/matrix"

func GradientDecent(f funcType, init [][]float64, lr float64, s int) [][]float64 {
	x := init

	for i := 0; i < s; i++ {
		var grad [][]float64 = NumericalGradient(f, x)
		x = matrix.Subtract2d(x, matrix.Scalar2d(grad, lr))
	}
	return x
}
