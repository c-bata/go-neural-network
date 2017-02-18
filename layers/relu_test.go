package layers

import (
	"reflect"
	"testing"
)

func TestMulLayer(t *testing.T) {
	const (
		NearEqualThreshold = 1e-8
	)

	l := &ReLULayer{}
	p := l.Forward([]float64{1, -1})

	if !reflect.DeepEqual(p, []float64{1, 0}) {
		t.Errorf("Forward() want %#v, got %#v", []float64{1, 0}, p)
	}
	if !reflect.DeepEqual(l.dx, []float64{1, 0}) {
		t.Errorf("Forward() should set %#v, got %#v", []float64{1, 0}, l.dx)
	}

	d := l.Backward([]float64{1.1})
	if (d[0] - 2.2) > NearEqualThreshold {
		t.Errorf("Backward() returns %#v as dx, but should be %#v", d[0], 2.2)
	}
	if (d[1] - 110) > NearEqualThreshold {
		t.Errorf("Backward() returns %#v as dy, but should be %#v", d[1], 110)
	}
}
