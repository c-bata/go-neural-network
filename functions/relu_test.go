package functions

import "testing"

func TestReLU(t *testing.T) {
	if ReLU(-0.5) != 0 {
		t.Fatal("ReLU(-0.5) should be 0, but doesn't match.")
	}

	if ReLU(0.5) != 0.5 {
		t.Fatal("ReLU(0.5) should be 0.5, but doesn't match.")
	}
}
