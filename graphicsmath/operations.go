package graphicsmath

import (
	"math"
	"github.com/smaTc/RayTracer/graphicsmath"
)


//Add Function
func Add(t1, t2 interface{}) interface{} {
	if isVector(t1) && isVector(t2) {
		tu1 := t1.(Vector)
		tu2 := t2.(Vector)
		return AddVector(&tu1, &tu2)
	} else if isPoint(t1) && isVector(t2) {
		tu1 := Tuple(t1.(Point))
		tu2 := Tuple(t2.(Vector))
		res := Point(AddTuple(&tu1, &tu2))
		return res
	} else if isPoint(t2) && isVector(t1) {
		tu1 := Tuple(t2.(Point))
		tu2 := Tuple(t1.(Vector))
		res := Point(AddTuple(&tu1, &tu2))
		return res
	}
	return nil
}

//Sub Function
func Sub(t1, t2 interface{}) interface{} {
	if isVector(t1) && isVector(t2) {
		tu1 := t1.(Vector)
		tu2 := t2.(Vector)
		return SubVector(&tu1, &tu2)
	} else if isPoint(t1) && isVector(t2) {
		tu1 := Tuple(t1.(Point))
		tu2 := Tuple(t2.(Vector))
		res := Point(SubTuple(&tu1, &tu2))
		return res
	} else if isPoint(t2) && isVector(t1) {
		tu1 := Tuple(t2.(Point))
		tu2 := Tuple(t1.(Vector))
		res := Point(SubTuple(&tu1, &tu2))
		return res
	}
	return nil
}

//
/*Tuple Functions*/

//AddTuple Function
func AddTuple(t1, t2 *Tuple) Tuple {
	return NewTuple(t1.X+t2.X, t1.Y+t2.Y, t1.Z+t2.Z, t1.W+t2.W)
}

//SubTuple Function
func SubTuple(t1, t2 *Tuple) Tuple {
	return NewTuple(t1.X-t2.X, t1.Y-t2.Y, t1.Z-t2.Z, t1.W-t2.W)
}

//
/*Vector Functions*/

//AddVector Function
func AddVector(v1, v2 *Vector) Vector {
	return NewVector(v1.X+v2.X, v1.Y+v2.Y, v1.Z+v2.Z)
}

//SubVector Function
func SubVector(v1, v2 *Vector) Vector {
	return NewVector(v1.X-v2.X, v1.Y-v2.Y, v1.Z-v2.Z)
}

//ScalVector Function
func ScalVector(s *float64, v *Vector) Vector {
	return Vector{X: *s * v.X, Y: *s * v.Y, Z: *s * v.Z, W: int(*s) * v.W}
}

//NormVector Function
func NormVector(v *Vector) Vector {
	n := math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2) + math.Pow(v.Z, 2) + math.Pow(float64(v.W), 2))
	return ScalVector(&n, v)
}

//NegVector Function
func NegVector(v *Vector) Vector {
	return Vector{X: -v.X, Y: -v.Y, Z: -v.Z, W: -v.W}
}

//Dot Function
func Dot(v1, v2 *Vector) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z + float64(v1.W*v2.W)
}

//Cross Function
func Cross(v1, v2 *Vector) Vector {
	return NewVector(
		v1.Y*v2.Z-v1.Z*v2.Y,
		v1.Z*v2.X-v1.X*v2.Z,
		v1.X*v2.Y-v1.Y*v2.X)
}

//
/*Point Functions*/

//AddPoint Function
func AddPoint(v1, v2 Point) Point {
	return NewPoint(v1.X+v2.X, v1.Y+v2.Y, v1.Z+v2.Z)
}

//SubPoint Function
func SubPoint(v1, v2 Point) Point {
	return NewPoint(v1.X-v2.X, v1.Y-v2.Y, v1.Z-v2.Z)
}
