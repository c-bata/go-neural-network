package layers

import (
	"reflect"
	"testing"
)

func TestMulLayer(t *testing.T) {
	const (
		NearEqualThreshold = 1e-8
	)

	m := &MulLayer{}
	p := m.Forward([]float64{100}, []float64{2})

	if p != 200 {
		t.Errorf("Forward() want %#v, got %#v", 200, p)
	}
	if !reflect.DeepEqual(m.x, []float64{100}) {
		t.Errorf("Forward() should set %#v, got %#v", 100, m.x)
	}
	if !reflect.DeepEqual(m.y, []float64{2}) {
		t.Errorf("Forward() should set %#v, got %#v", 2, m.y)
	}

	dx, dy := m.Backward(1.1)
	if (dx[0] - 2.2) > NearEqualThreshold {
		t.Errorf("Backward() returns %#v as dx, but should be %#v", dx, []float64{2.2})
	}
	if (dy[0] - 110) > NearEqualThreshold {
		t.Errorf("Backward() returns %#v as dy, but should be %#v", dy, []float64{110})
	}
}
