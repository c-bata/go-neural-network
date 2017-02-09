package matrix


func Subtract1d(a []float64, b []float64) []float64 {
	l := len(a)
	c := make([]float64, l)

	for i := 0; i < l; i++ {
		c[i] = a[i] - b[i]
	}
	return c
}


func Subtract2d(a [][]float64, b [][]float64) [][]float64 {
	l := len(a)
	c := make([][]float64, l)

	for i := 0; i < l; i++ {
		c[i] = Subtract1d(a[i], b[i])
	}
	return c
}
