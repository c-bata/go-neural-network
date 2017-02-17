package matrix

func Add1d(a []float64, b []float64) []float64 {
	l := len(a)
	c := make([]float64, l)

	for i := 0; i < l; i++ {
		c[i] = a[i] + b[i]
	}
	return c
}

func Add2d(a [][]float64, b [][]float64) [][]float64 {
	l := len(a)
	c := make([][]float64, l)

	for i := 0; i < l; i++ {
		c[i] = Add1d(a[i], b[i])
	}
	return c
}
