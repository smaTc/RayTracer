package graphics

import (
	"testing"
)

func TestTuple(t *testing.T) {
	v1 := NewVector(1, 2, 3)
	v2 := NewVector(1, 2, 3)
	v3 := NewVector(3, 2, 1)
	p1 := NewPoint(1, 2, 3)
	p2 := NewPoint(1, 2, 3)
	p3 := NewPoint(3, 2, 1)

	if !Equals(v1, v2) || !Equals(p1, p2) {
		t.Errorf("Vectors/Points were equal, incorrect eval")
	}

	if Equals(v1, v3) || Equals(p1, p3) {
		t.Errorf("Vectors/Points were not equal, incorrect eval")
	}

	if Equals(v1, p2) || Equals(p1, v2) {
		t.Errorf("Vectors/Points were not equal, incorrect eval")
	}
}
