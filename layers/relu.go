package layers

type ReLULayer struct {
	dx []float64
}

func (l *ReLULayer) Forward(x []float64) []float64 {
	lx := len(x)
	l.dx = make([]float64, lx)
	y := make([]float64, lx)
	for i, v := range x {
		if v > 0 {
			y[i] = v
			l.dx[i] = 1.0
		}
	}
	return y
}

func (l *ReLULayer) Backward(y []float64) []float64 {
	return l.dx
}
