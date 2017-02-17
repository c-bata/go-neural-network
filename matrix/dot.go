package matrix

func Dot1x1(a []float64, b []float64) float64 {
	var sum float64 = 0
	for i := 0; i < len(a); i++ {
		sum += a[i] * b[i]
	}
	return sum
}

func Dot1x2(a []float64, b [][]float64) []float64 {
	w := len(a)
	h := len(b[0])
	c := make([]float64, h)

	for i := 0; i < h; i++ {
		var sum float64 = 0
		for j := 0; j < w; j++ {
			sum += a[j] * b[j][i]
		}
		c[i] = sum
	}
	return c
}

func Dot2x2(a [][]float64, b [][]float64) [][]float64 {
	w := len(a)
	h := len(b[0])
	c := make([][]float64, h)

	for i := 0; i < w; i++ {
		c[i] = make([]float64, w)
		for j := 0; j < h; j++ {
			var sum float64 = 0
			for k := 0; k < len(a[0]); k++ {
				sum += a[i][k] * b[k][j]
			}
			c[i][j] = sum
		}
	}
	return c
}
