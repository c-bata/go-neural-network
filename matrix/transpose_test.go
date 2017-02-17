package matrix

import (
	"testing"
	"reflect"
)

func TestTranspose(t *testing.T) {
	x := [][]float64{{1, 2}, {3, 4}}
	y := Transpose(x)
	if !reflect.DeepEqual(y, [][]float64{{1, 3}, {2, 4}}) {
		t.Errorf("Transpose() want %#v, got %#v\n", [][]float64{{1, 3}, {2, 4}}, y)
	}
}
