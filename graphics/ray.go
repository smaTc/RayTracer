package graphics

import (
	"fmt"
	"math"
)

//Ray struct
type Ray struct {
	Origin    Point
	Direction Vector
}

//NewRay func
func NewRay(o Point, d Vector) Ray {
	return Ray{Origin: o, Direction: d}
}

//SetRayProps func
func (r *Ray) SetRayProps(o Point, d Vector) {
	r.Origin = o
	r.Direction = d
}

//Position func
func (r *Ray) Position(t float32) Point {
	vec := r.Direction.Scale(t)
	newPoint, _ := r.Origin.Add(vec)
	return newPoint
}

//Intersects func
func (r *Ray) Intersects(s Sphere) []float32 {
	return Intersect(*r, s)
}

//Intersect func
func Intersect(r Ray, s Sphere) []float32 {
	discr, a, b := Discriminant(r, s)
	if discr < 0 {
		return nil
	}

	t1 := (-b - float32(math.Sqrt(float64(discr)))) / (2 * a)
	t2 := (-b + float32(math.Sqrt(float64(discr)))) / (2 * a)
	fmt.Println("intersections:", t1, ",", t2)

	if t1 <= t2 {
		return []float32{t1, t2}
	}
	return []float32{t2, t1}
}

//Discriminant func
func Discriminant(r Ray, s Sphere) (float32, float32, float32) {
	sphereToRay, _ := r.Origin.Subtract(s.Origin)
	//phereToRay, _ := s.Origin.Subtract(r.Origin)

	sphereToRayVector := sphereToRay.(Vector)
	fmt.Println("SpheretoRay", sphereToRayVector)
	a := Dot(&r.Direction, &r.Direction)
	b := 2 * Dot(&r.Direction, &sphereToRayVector)
	c := Dot(&sphereToRayVector, &sphereToRayVector) - 1

	discriminant := b*b - 4*a*c
	fmt.Println("dab", discriminant, a, b)
	return discriminant, a, b
}
