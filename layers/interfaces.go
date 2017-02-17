package layers

type Layer interface {
	Forward([]float64) float64
	Backward(float64) []float64
}
