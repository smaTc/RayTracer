package graphics

import "testing"

func TestPosition(t *testing.T) {

	r := NewRay(NewPoint(2, 3, 4), NewVector(1, 0, 0))
	p1 := NewPoint(2, 3, 4)
	p2 := NewPoint(3, 3, 4)
	p3 := NewPoint(1, 3, 4)
	p4 := NewPoint(4.5, 3, 4)

	rp1 := r.Position(0)
	rp2 := r.Position(1)
	rp3 := r.Position(-1)
	rp4 := r.Position(2.5)

	if !rp1.Equals(p1) && !rp2.Equals(p2) && !rp3.Equals(p3) && !rp4.Equals(p4) {
		t.Error("Ray Position not working")
	}

}

func TestIntersection(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	s := NewSphere(NewPoint(0, 0, 0), 1)
	xs := r.Intersects(s)

	if xs[0] != 4.0 && xs[1] != 6.0 {
		t.Error("intersect not working")
	}

}

func TestIntersectionTangent(t *testing.T) {
	r := NewRay(NewPoint(0, 1, -5), NewVector(0, 0, 1))
	s := NewSphere(NewPoint(0, 0, 0), 1)
	xs := r.Intersects(s)

	if xs[0] != 5.0 && xs[1] != 5.0 {
		t.Error("intersect not working")
	}
}

func TestRayMissesSpehere(t *testing.T) {
	r := NewRay(NewPoint(0, 2, -5), NewVector(0, 0, 1))
	s := NewSphere(NewPoint(0, 0, 0), 1)
	xs := r.Intersects(s)

	if len(xs) != 0 {
		t.Error("intersect not working")
	}
}

func TestRayOriginSphere(t *testing.T) {
	r := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
	s := NewSphere(NewPoint(0, 0, 0), 1)
	xs := r.Intersects(s)

	if xs[0] != -1.0 && xs[1] != 1.0 {
		t.Error("intersect not working")
	}
}

func TestSphereBehindRay(t *testing.T) {
	r := NewRay(NewPoint(0, 0, 5), NewVector(0, 0, 1))
	s := NewSphere(NewPoint(0, 0, 0), 1)
	xs := r.Intersects(s)

	if xs[0] != -6.0 && xs[1] != -4.0 {
		t.Error("intersect not working")
	}
}
