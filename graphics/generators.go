package graphics

import (
	"errors"
)

//ConvertTuple Function
func convertTuple(x, y, z, w float32) (interface{}, error) {
	if w == 1 {
		return Point{X: x, Y: y, Z: z, W: w}, nil
	} else if w == 0 {
		return Vector{X: x, Y: y, Z: z, W: w}, nil
	} else {
		return nil, errors.New("wrong W parameter")
	}
}

//NewTuple Function
func NewTuple(x, y, z, w float32) Tuple {
	return Tuple{X: x, Y: y, Z: z, W: w}
}

//NewVector Function
func NewVector(x, y, z float32) Vector {
	return Vector{X: x, Y: y, Z: z, W: 0}
}

//NewPoint Function
func NewPoint(x, y, z float32) Point {
	return Point{X: x, Y: y, Z: z, W: 1}
}
