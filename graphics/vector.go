package graphics

import "math"

//Vector struct
type Vector struct {
	X, Y, Z float32
	W       float32
}

//Add func
func (v1 *Vector) Add(v2 Vector) Vector {
	vec := AddVectors(v1, &v2)
	*v1 = vec
	return vec
}

//Substract func
func (v1 *Vector) Substract(v2 Vector) Vector {
	vec := SubVectors(v1, &v2)
	*v1 = vec
	return vec
}

//Equals Vector func
func (v1 *Vector) Equals(v2 Vector) bool {
	return *v1 == v2
}

//Scale Function
func (v1 *Vector) Scale(s float32) Vector {
	scaledVec := NewVector(v1.X, v1.Y, v1.Z)
	scaledVec.X *= s
	scaledVec.Y *= s
	scaledVec.Z *= s
	return scaledVec
	//v1.X *= s
	//v1.Y *= s
	//v1.Z *= s
}

//Normalize Function
func (v1 *Vector) Normalize() *Vector {
	n := math.Sqrt(math.Pow(float64(v1.X), 2) + math.Pow(float64(v1.Y), 2) + math.Pow(float64(v1.Z), 2) + math.Pow(float64(v1.W), 2))
	v1.Scale(float32(n))
	return v1
}

//Negate Function
func (v1 *Vector) Negate() {
	//vec := Vector{X: -v1.X, Y: -v1.Y, Z: -v1.Z, W: 1}
	v1.X = -v1.X
	v1.Y = -v1.Y
	v1.Z = -v1.Z
	v1.W = -v1.W
}
