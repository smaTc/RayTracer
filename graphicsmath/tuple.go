package graphicsmath

import (
	"errors"
)

//ConvertTuple Function
func convertTuple(x, y, z float64, w int) (interface{}, error) {
	if w == 1 {
		return Point{X: x, Y: y, Z: z, W: w}, nil
	} else if w == 0 {
		return Vector{X: x, Y: y, Z: z, W: w}, nil
	} else {
		return nil, errors.New("wrong W parameter")
	}
}

//NewTuple Function
func NewTuple(x, y, z float64, w int) Tuple {
	return Tuple{X: x, Y: y, Z: z, W: w}
}

//NewVector Function
func NewVector(x, y, z float64) Vector {
	return Vector{X: x, Y: y, Z: z, W: 0}
}

//NewPoint Function
func NewPoint(x, y, z float64) Point {
	return Point{X: x, Y: y, Z: z, W: 1}
}

//Equals checks if two Tupels,Points,Vectors are equal
func Equals(t1, t2 interface{}) bool {
	if isVector(t1) && isVector(t2) {
		return t1.(Vector).X == t2.(Vector).X && t1.(Vector).Y == t2.(Vector).Y && t1.(Vector).Z == t2.(Vector).Z
	} else if isPoint(t1) && isPoint(t2) {
		return t1.(Point).X == t2.(Point).X && t1.(Point).Y == t2.(Point).Y && t1.(Point).Z == t2.(Point).Z
	}
	return false
}

func isVector(i interface{}) bool {
	switch i.(type) {
	case Vector:
		return true
	default:
		return false
	}
}

func isPoint(i interface{}) bool {
	switch i.(type) {
	case Point:
		return true
	default:
		return false
	}
}
