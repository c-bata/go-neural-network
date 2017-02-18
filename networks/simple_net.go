package networks

import (
	"github.com/c-bata/go-neural-network/matrix"
	"github.com/c-bata/go-neural-network/functions"
	"math/rand"
)

type SimpleNet struct {
	Weight [][]float64
}

func (s SimpleNet) Predict(x []float64) []float64 {
	y, _ := matrix.Dot1x2(x, s.Weight)
	return y
}

func (s SimpleNet) Loss(x []float64, t []float64) float64 {
	var z []float64 = s.Predict(x)
	var y []float64 = functions.SoftMax(z)
	l := functions.CrossEntropyError(y, t)
	return l
}

func NewSimpleNet() *SimpleNet {
	n := &SimpleNet{}

	w := make([][]float64, 2)
	for i := 0; i < 2; i++ {
		w[i] = make([]float64, 3)
		for j := 0; j < 3; j++ {
			w[i][j] = rand.Float64()
		}
	}
	n.Weight = w
	return n
}
