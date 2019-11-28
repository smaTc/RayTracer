package types

import "errors"

//Tuple struct
type Tuple struct {
	X, Y, Z float64
	W       int
}

//Point struct
type Point struct {
	X, Y, Z float64
	W       int
}

//Vector struct
type Vector struct {
	X, Y, Z float64
	W       int
}

//NewTuple getnewTuple
func (t Tuple) NewTuple(x, y, z float64, w int) (interface{}, error) {
	if w == 1 {
		return Point{X: x, Y: y, Z: z, W: w}, nil
	} else if w == 0 {
		return Vector{X: x, Y: y, Z: z, W: w}, nil
	} else {
		return nil, errors.New("wrong W parameter")
	}
}
