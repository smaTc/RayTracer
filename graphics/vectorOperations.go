package graphics

//AddVectors Function
func AddVectors(v1, v2 *Vector) Vector {
	return NewVector(v1.X+v2.X, v1.Y+v2.Y, v1.Z+v2.Z)
}

//SubVectors Function
func SubVectors(v1, v2 *Vector) Vector {
	return NewVector(v1.X-v2.X, v1.Y-v2.Y, v1.Z-v2.Z)
}

//NegateVector func
func NegateVector(v1 *Vector) Vector {
	return Vector{X: -v1.X, Y: -v1.Y, Z: -v1.Z, W: 1}
}

func isVector(i interface{}) bool {
	switch i.(type) {
	case Vector:
		return true
	default:
		return false
	}
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
