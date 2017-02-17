package matrix

func Transpose(x [][]float64) [][]float64 {
	c := len(x)
	r := len(x[0])
	y := make([][]float64, r)
	for i := 0; i < r; i++ {
		y[i] = make([]float64, c)
		for j := 0; j < c; j++ {
			y[i][j] = x[j][i]
		}
	}
	return y
}
