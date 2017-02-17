package matrix

import (
	"testing"
	"reflect"
)

func TestSubtract1d(t *testing.T) {
	x := Subtract1d([]float64{1, 2}, []float64{2, 2})
	if !reflect.DeepEqual(x, []float64{-1, 0}) {
		t.Errorf("Subtract1d want %#v got %#v", []float64{-1, 0}, x)
	}
}

func TestSubtract2d(t *testing.T) {
	x := Subtract2d([][]float64{{1, 2}, {3, 4}}, [][]float64{{1, 3}, {3, 4}})
	if !reflect.DeepEqual(x, [][]float64{{0, -1}, {0, 0}}) {
		t.Errorf("Substract2d want %#v got %#v", [][]float64{{0, -1}, {0, 0}}, x)
	}
}
