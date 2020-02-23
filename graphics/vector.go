package graphics

import "math"

//Vector struct
type Vector struct {
	X, Y, Z float32
	W       int
}

func isVector(i interface{}) bool {
	switch i.(type) {
	case Vector:
		return true
	default:
		return false
	}
}

//Equals Vector func
func (v1 *Vector) Equals(v2 Vector) bool {
	return *v1 == v2
}

//AddVectors Function
func AddVectors(v1, v2 *Vector) Vector {
	return NewVector(v1.X+v2.X, v1.Y+v2.Y, v1.Z+v2.Z)
}

//SubVector Function
func SubVector(v1, v2 *Vector) Vector {
	return NewVector(v1.X-v2.X, v1.Y-v2.Y, v1.Z-v2.Z)
}

//Scale Function
func (v1 *Vector) Scale(s float32) {
	v1.X *= s
	v1.Y *= s
	v1.Z *= s
}

//Normalize Function
func (v1 *Vector) Normalize() *Vector {
	n := math.Sqrt(math.Pow(float64(v1.X), 2) + math.Pow(float64(v1.Y), 2) + math.Pow(float64(v1.Z), 2) + math.Pow(float64(v1.W), 2))
	v1.Scale(float32(n))
	return v1
}

//Negate Function
func (v1 *Vector) Negate() {
	v1.X = -v1.X
	v1.Y = -v1.Y
	v1.Z = -v1.Z
	v1.W = -v1.W
}

//Dot Function
func Dot(v1, v2 *Vector) float32 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z + float32(v1.W*v2.W)
}

//Cross Function
func Cross(v1, v2 *Vector) Vector {
	return NewVector(
		v1.Y*v2.Z-v1.Z*v2.Y,
		v1.Z*v2.X-v1.X*v2.Z,
		v1.X*v2.Y-v1.Y*v2.X)
}
