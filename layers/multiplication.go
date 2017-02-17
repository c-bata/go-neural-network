package layers

import "github.com/c-bata/go-neural-network/matrix"

type MulLayer struct {
	x []float64
	y []float64
}

func (m *MulLayer) Forward(x []float64, y []float64) float64 {
	m.x = x
	m.y = y

	var sum float64 = 0
	for i := 0; i < len(x); i++ {
		sum += x[i] * y[i]
	}
	return sum
}

func (m *MulLayer) Backward(d float64) ([]float64, []float64) {
	dx := matrix.Scalar1d(m.y, d)
	dy := matrix.Scalar1d(m.x, d)

	return dx, dy
}
