package graphics

import "math/rand"

//Sphere struct
type Sphere struct {
	Origin    Point
	Radius    float32
	Transform Matrix
}

//SetTransformationMatrix func
func (s *Sphere) SetTransformationMatrix(matrix Matrix) {
	s.Transform = matrix
}

//Equals func
func (s *Sphere) Equals(s2 Sphere) bool {
	if !s.Origin.Equals(s2.Origin) || s.Radius != s2.Radius || !s.Transform.Equals(s2.Transform.Values) {
		return false
	}
	return true
}

//NewSphere func
func NewSphere(orig Point, rad float32) Sphere {
	sphere := Sphere{Origin: orig, Radius: rad}
	var transMat Matrix
	transMat.GenerateMatrixWithDimension(4, 4)
	transMat.Values = [][]float32{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
	sphere.Transform = transMat
	return sphere
}

//RandSphere func
func RandSphere() float32 {
	return rand.Float32()
}
