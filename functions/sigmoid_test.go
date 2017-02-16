package functions

import "testing"

func TestSigmoid(t *testing.T) {
	const (
		NearEqualThreshold = 1e-8
	)

	if Sigmoid(0) != 0.5 {
		t.Fatal("Sigmoid(0) should be 0.5, but doesn't match.")
	}

	if Sigmoid(0.3)-0.57444252 > NearEqualThreshold {
		t.Fatal("Sigmoid(0.3) should be 0.57444252 but doesn't match.")
	}
	if Sigmoid(0.7)-0.66818777 > NearEqualThreshold {
		t.Fatal("Sigmoid(0.7) should be 0.66818777 but doesn't match.")
	}

	if Sigmoid(1.1)-0.75026011 > NearEqualThreshold {
		t.Fatal("Sigmoid(1.1) should be 0.75026011 but doesn't match.")
	}
}
