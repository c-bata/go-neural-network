package matrix

import (
	"reflect"
	"testing"
)

func TestDot1x1(t *testing.T) {
	a := []float64{1, 2, 3}
	b := []float64{1, 2, 3}
	var ex float64 = 1*1 + 2*2 + 3*3
	c, err := Dot1x1(a, b)
	if err != nil {
		t.Errorf("Dot1x1() should not be err, but %#v", err)
	}

	if c != ex {
		t.Errorf("Dot1x1() want %#v got %#v.", ex, c)
	}
}

func TestDot1x2(t *testing.T) {
	a := []float64{
		1, 2, 3,
	}
	b := [][]float64{
		{1, 4},
		{2, 5},
		{3, 6},
	}
	ex := []float64{
		14, 32,
	}
	c, err := Dot1x2(a, b)
	if err != nil {
		t.Errorf("Dot1x2() should not be err, but %#v", err)
	}

	if !reflect.DeepEqual(c, ex) {
		t.Errorf("Dot1x2() want %#v got %#v.", ex, c)
	}
}

func TestDot2x2(t *testing.T) {
	a := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}
	b := [][]float64{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	ex := [][]float64{
		{22, 28},
		{49, 64},
	}
	c, err := Dot2x2(a, b)

	if err != nil {
		t.Errorf("Dot2x2() should not be err, but %#v", err)
	}

	if len(c) != 2 || len(c[0]) != 2 {
		t.Fatal("Dot2x2() should return 2x2 arrays.")
	}

	if !reflect.DeepEqual(c, ex) {
		t.Errorf("Dot2x2() want %#v got %#v.", ex, c)
	}
}

func BenchmarkDot1x1(b *testing.B) {
	a1, a2 := make([]float64, 1024), make([]float64, 1024)
	sum := 0.0
	for i := 0; i < b.N; i++ {
		s, _ := Dot1x1(a1, a2)
		sum += s
	}
}
