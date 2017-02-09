package functions

import "testing"

func TestSoftMax(t *testing.T) {
	const (
		NearEqualThreshold = 1e-8
	)

	e := []float64{0.01821127, 0.24519181, 0.73659691}
	x := []float64{0.3, 2.9, 4.0}
	a := SoftMax(x)
	for i := 0; i < 3; i ++ {
		if a[i] - e[i] > NearEqualThreshold {
			t.Fatalf("Softmax(%v) should be %v, but return %v.", x, e, a)
		}
	}
}
