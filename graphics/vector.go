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

//Multiply func
func (v1 *Vector) Multiply(input interface{}) Vector {
	if isMatrix(input) {
		m := input.(Matrix)

		if r, c := m.GetDimensions(); r != c && r != 4 {
			return Vector{}
		}
		var nv Vector
		col0 := m.GetColumn(0)
		col1 := m.GetColumn(1)
		col2 := m.GetColumn(2)
		col3 := m.GetColumn(3)

		nv.X = v1.X*col0[0] + v1.Y*col0[1] + v1.Z*col0[2] + v1.W*col0[3]
		nv.Y = v1.X*col1[0] + v1.Y*col1[1] + v1.Z*col1[2] + v1.W*col1[3]
		nv.Z = v1.X*col2[0] + v1.Y*col2[1] + v1.Z*col2[2] + v1.W*col2[3]
		nv.W = v1.X*col3[0] + v1.Y*col3[1] + v1.Z*col3[2] + v1.W*col3[3]

		return nv
	}
	return Vector{}
}
