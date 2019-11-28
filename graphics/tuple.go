package graphics

import (
	"errors"

	"github.com/smaTc/RayTracer/types"
)

//ConvertTuple Function
func ConvertTuple(x, y, z float64, w int) (interface{}, error) {
	if w == 1 {
		return types.Point{X: x, Y: y, Z: z, W: w}, nil
	} else if w == 0 {
		return types.Vector{X: x, Y: y, Z: z, W: w}, nil
	} else {
		return nil, errors.New("wrong W parameter")
	}
}

//Tuple Function
func Tuple(x, y, z float64, w int) types.Tuple {
	return types.Tuple{X: x, Y: y, Z: z, W: w}
}

//Vector Function
func Vector(x, y, z float64) types.Vector {
	return types.Vector{X: x, Y: y, Z: z, W: 0}
}

//Point Function
func Point(x, y, z float64) types.Point {
	return types.Point{X: x, Y: y, Z: z, W: 1}
}

//Equals checks if two Tupels,Points,Vectors are equal
func Equals(t1, t2 interface{}) bool {
	if isVector(t1) && isVector(t2) {
		return t1.(types.Vector).X == t2.(types.Vector).X && t1.(types.Vector).Y == t2.(types.Vector).Y && t1.(types.Vector).Z == t2.(types.Vector).Z
	} else if isPoint(t1) && isPoint(t2) {
		return t1.(types.Point).X == t2.(types.Point).X && t1.(types.Point).Y == t2.(types.Point).Y && t1.(types.Point).Z == t2.(types.Point).Z
	}
	return false
}

func isVector(i interface{}) bool {
	switch i.(type) {
	case types.Vector:
		return true
	default:
		return false
	}
}

func isPoint(i interface{}) bool {
	switch i.(type) {
	case types.Point:
		return true
	default:
		return false
	}
}
