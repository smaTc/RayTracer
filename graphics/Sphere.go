package graphics

import "math/rand"

//Sphere struct
type Sphere struct {
	Origin Point
	Radius float32
}

//NewSphere func
func NewSphere(orig Point, rad float32) Sphere {
	return Sphere{Origin: orig, Radius: rad}
}

//RandSphere func
func RandSphere() float32 {
	return rand.Float32()
}
