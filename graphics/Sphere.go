package graphics

import "math/rand"

//Sphere struct
type Sphere struct {
	Origin Point
	Radius float32
}

//NewSphere func
func NewSphere() Sphere {
	return Sphere{}
}

//RandSphere func
func RandSphere() float32 {
	return rand.Float32()
}
