package graphics

import "testing"

func TestAggreagatingIntersections(t *testing.T) {
	s := NewSphere(NewPoint(0, 0, 0), 1)
	i1 := NewIntersection(1, s)
	i2 := NewIntersection(2, s)
	xs := Intersections(i1, i2)

	if len(xs) != 2 || xs[0].T != 1 || xs[1].T != 2 {
		t.Error("aggregation not working")
	}

}

func TestIntersectionsSetsObject(t *testing.T) {
	s := NewSphere(NewPoint(0, 0, 0), 1)
	i1 := NewIntersection(1, s)
	i2 := NewIntersection(2, s)
	xs := Intersections(i1, i2)

	if len(xs) != 2 || xs[0].Object != s || xs[1].Object != s {
		t.Error("aggregation not working")
	}

}
