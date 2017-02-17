package matrix

import (
	"testing"
	"reflect"
)

func TestDot1x1(t *testing.T) {
	a := []float64{1, 2, 3}
	b := []float64{1, 2, 3}
	var c float64 = Dot1x1(a, b)

	if c != 1*1+2*2+3*3 {
		t.Errorf("Dot1x1() want %#v got %#v.", 1*1+2*2+3*3, c)
	}
}

func TestDot1x2(t *testing.T) {
	a := []float64{1, 2, 3}
	b := [][]float64{{1, 4}, {2, 5}, {3, 6}}
	var c []float64 = Dot1x2(a, b)

	if !reflect.DeepEqual(c, []float64{14, 32}) {
		t.Errorf("Dot1x2() want %#v got %#v.", []float64{14, 32}, c)
	}
}

func TestDot2x2(t *testing.T) {
	a := [][]float64{{1, 2, 3}, {4, 5, 6}}
	b := [][]float64{{1, 2}, {3, 4}, {5, 6}}
	c := Dot2x2(a, b)

	if len(c) != 2 || len(c[0]) != 2 {
		t.Fatal("Dot2x2() should return 2x2 arrays.")
	}

	if !reflect.DeepEqual(c, [][]float64{{22, 28}, {49, 64}}) {
		t.Errorf("Dot2x2() want %#v got %#v.", [][]float64{{22, 28}, {49, 64}}, c)
	}
}
