package graphics

import "github.com/smaTc/RayTracer/types"

import "math"

/*func AddTuple(t1,t2 types.Tuple) types.Tuple{
	return Tuple(v1.X+v2.X, v1.Y+v2.Y, v1.Z+v2.Z)
}*/

//
/*Tuple Functions*/

//AddTuple Function
func AddTuple(t1, t2 *types.Tuple) types.Tuple {
	return Tuple(t1.X+t2.X, t1.Y+t2.Y, t1.Z+t2.Z, t1.W+t2.W)
}

//SubTuple Function
func SubTuple(t1, t2 *types.Tuple) types.Tuple {
	return Tuple(t1.X-t2.X, t1.Y-t2.Y, t1.Z-t2.Z, t1.W-t2.W)
}

//
/*Vector Functions*/

//AddVector Function
func AddVector(v1, v2 *types.Vector) types.Vector {
	return Vector(v1.X+v2.X, v1.Y+v2.Y, v1.Z+v2.Z)
}

//SubVector Function
func SubVector(v1, v2 *types.Vector) types.Vector {
	return Vector(v1.X-v2.X, v1.Y-v2.Y, v1.Z-v2.Z)
}

//ScalVector Function
func ScalVector(s *float64, v *types.Vector) types.Vector {
	return types.Vector{X: *s * v.X, Y: *s * v.Y, Z: *s * v.Z, W: int(*s) * v.W}
}

//NormVector Function
func NormVector(v *types.Vector) types.Vector {
	n := math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2) + math.Pow(v.Z, 2) + math.Pow(float64(v.W), 2))
	return ScalVector(&n, v)
}

//NegVector Function
func NegVector(v *types.Vector) types.Vector {
	return types.Vector{X: -v.X, Y: -v.Y, Z: -v.Z, W: -v.W}
}

//Dot Function
func Dot(v1, v2 *types.Vector) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z + float64(v1.W*v2.W)
}

//Cross Function
func Cross(v1, v2 *types.Vector) types.Vector {
	return Vector(
		v1.Y*v2.Z-v1.Z*v2.Y,
		v1.Z*v2.X-v1.X*v2.Z,
		v1.X*v2.Y-v1.Y*v2.X)
}

//
/*Point Functions*/

//AddPoint Function
func AddPoint(v1, v2 types.Vector) types.Point {
	return Point(v1.X+v2.X, v1.Y+v2.Y, v1.Z+v2.Z)
}

//SubPoint Function
func SubPoint(v1, v2 types.Vector) types.Point {
	return Point(v1.X-v2.X, v1.Y-v2.Y, v1.Z-v2.Z)
}
