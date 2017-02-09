package matrix


func Scalar1d(x []float64, s float64) []float64 {
	for i, a := range x {
		x[i] = s * a
	}
	return x
}


func Scalar2d(x [][]float64, s float64) [][]float64 {
	for i, a := range x {
		x[i] = Scalar1d(a, s)
	}
	return x
}
