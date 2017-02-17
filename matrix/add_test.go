package matrix

import (
	"testing"
	"reflect"
)

func TestAdd1d(t *testing.T) {
	x := Add1d([]float64{1, 2}, []float64{2, 2})
	if !reflect.DeepEqual(x, []float64{3, 4}) {
		t.Errorf("Add1d want %#v got %#v", []float64{3, 4}, x)
	}
}

func TestAdd2d(t *testing.T) {
	x := Add2d([][]float64{{1, 2}, {3, 4}}, [][]float64{{1, 2}, {3, 4}})
	if !reflect.DeepEqual(x, [][]float64{{2, 4}, {6, 8}}) {
		t.Errorf("Add2d want %#v got %#v", [][]float64{{2, 4}, {6, 8}}, x)
	}
}
