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

	if len(xs) != 2 || !s.Equals(xs[0].Object.(Sphere)) || !s.Equals(xs[1].Object.(Sphere)) {
		t.Error("object set not working")
	}

}

func TestHitWithPosTs(t *testing.T) {
	s := NewSphere(NewPoint(0, 0, 0), 1)
	i1 := NewIntersection(1, s)
	i2 := NewIntersection(2, s)
	xs := Intersections(i1, i2)
	i := Hit(xs)

	if !i.Equals(i1) {
		t.Error("Hit with pos. Ts not working")
	}

}

func TestHitWithSomeNegTs(t *testing.T) {
	s := NewSphere(NewPoint(0, 0, 0), 1)
	i1 := NewIntersection(-1, s)
	i2 := NewIntersection(1, s)
	xs := Intersections(i1, i2)
	i := Hit(xs)

	if !i.Equals(i2) {
		t.Error("Hit with some neg. Ts not working")
	}
}

func TestHitWithNegTs(t *testing.T) {
	s := NewSphere(NewPoint(0, 0, 0), 1)
	i1 := NewIntersection(-2, s)
	i2 := NewIntersection(-1, s)
	xs := Intersections(i1, i2)
	i := Hit(xs)
	iRes := Intersection{}

	if !i.Equals(iRes) {
		t.Error("Hit with neg. Ts not working")
	}
}

func TestHitWithLowestPosIntersec(t *testing.T) {
	s := NewSphere(NewPoint(0, 0, 0), 1)
	i1 := NewIntersection(5, s)
	i2 := NewIntersection(7, s)
	i3 := NewIntersection(-3, s)
	i4 := NewIntersection(2, s)
	xs := Intersections(i1, i2, i3, i4)
	i := Hit(xs)

	if !i.Equals(i4) {
		t.Error("Hit with lowest non neg not working")
	}
}
