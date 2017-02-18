package matrix

import (
	"fmt"
	"time"
)

type ArraySizeError struct {
	When time.Time
	What string
}

func (e *ArraySizeError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func Dot1x1(a, b []float64) (sum float64, err error) {
	la, lb := len(a), len(b)
	if la != lb {
		err = &ArraySizeError{
			time.Now(),
			fmt.Sprintf("Couldn't dot calcuration: %d・%d", la, lb),
		}
		return
	}

	sum = 0
	for i := 0; i < la; i++ {
		sum += a[i] * b[i]
	}
	return
}

func Dot1x2(a []float64, b [][]float64) (c []float64, err error) {
	la, lb1, lb2 := len(a), len(b), len(b[0])
	c = make([]float64, lb2)
	if la != lb1 {
		err = &ArraySizeError{
			time.Now(),
			fmt.Sprintf("Couldn't dot calcuration: %d・%dx%d", la, lb1, lb2),
		}
		return
	}

	for i := 0; i < lb2; i++ {
		var sum float64 = 0
		for j := 0; j < la; j++ {
			sum += a[j] * b[j][i]
		}
		c[i] = sum
	}
	return
}

func Dot2x2(a [][]float64, b [][]float64) (c [][]float64, err error) {
	la1, la2, lb1, lb2 := len(a), len(a[0]), len(b), len(b[0])
	if la2 != lb1 {
		err = &ArraySizeError{
			time.Now(),
			fmt.Sprintf("Couldn't dot calcuration: %dx%d・%dx%d", la1, la2, lb1, lb2),
		}
		return
	}

	c = make([][]float64, lb2)

	for i := 0; i < la1; i++ {
		c[i] = make([]float64, la1)
		for j := 0; j < lb2; j++ {
			var sum float64 = 0
			for k := 0; k < la2; k++ {
				sum += a[i][k] * b[k][j]
			}
			c[i][j] = sum
		}
	}
	return
}
