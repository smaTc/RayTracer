package graphics

import (
	"testing"
)

func Test(t *testing.T) {
	v1 := Vector(1, 2, 3)
	v2 := Vector(1, 2, 3)
	v3 := Vector(3, 2, 1)
	p1 := Point(1, 2, 3)
	p2 := Point(1, 2, 3)
	p3 := Point(3, 2, 1)

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
