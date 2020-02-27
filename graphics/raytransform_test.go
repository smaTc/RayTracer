package graphics

import "testing"

func TestRayTranslation(t *testing.T) {
	r := NewRay(NewPoint(1, 2, 3), NewVector(0, 1, 0))
	m := Translation(3, 4, 5)
	r2 := r.Transform(m)

	if r2.Origin != NewPoint(4, 6, 8) || r2.Direction != NewVector(0, 1, 0) {
		t.Error("Ray Translation not working")
	}

}

func TestRayScaling(t *testing.T) {
	r := NewRay(NewPoint(1, 2, 3), NewVector(0, 1, 0))
	m := Scaling(2, 3, 4)
	r2 := r.Transform(m)

	if r2.Origin != NewPoint(2, 6, 12) || r2.Direction != NewVector(0, 3, 0) {
		t.Error("Ray Scaling not working")
	}
}

func TestSphereTransform(t *testing.T) {
	s := NewSphere(NewPoint(0, 0, 0), 1)
	idMat := [][]float32{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}

	if !CompareMatrixArrays(s.Transform.Values, idMat) {
		t.Error("Sphere default ID not correct")
	}

	tl := Translation(2, 3, 4)

	s.SetTransformationMatrix(tl)

	if !CompareMatrixArrays(tl.Values, s.Transform.Values) {
		t.Error("Setting sphere not working")
	}

}

func TestIntersectScaledSphereWithRay(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	s := NewSphere(NewPoint(0, 0, 0), 1)

	s.Transform = Scaling(2, 2, 2)
	xs := Intersect(r, s)

	if len(xs) != 2 || xs[0].T != 3 || xs[1].T != 7 {
		t.Error("Intersecting scaled sphere with ray not working")
	}
}

func TestIntersectTranslatedSphereWithRay(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	s := NewSphere(NewPoint(0, 0, 0), 1)

	s.Transform = Translation(5, 0, 0)
	xs := Intersect(r, s)

	if len(xs) != 0 {
		t.Error("Intersecting tranlated sphere with ray not working")
	}
}
