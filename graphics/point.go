package graphics

import "errors"

//Point struct
type Point struct {
	X, Y, Z float32
	W       float32
}

func isPoint(i interface{}) bool {
	switch i.(type) {
	case Point:
		return true
	default:
		return false
	}
}

//Equals Function
func (p1 *Point) Equals(p2 Point) bool {
	return *p1 == p2
}

//AddPoint Function
func AddPoint(v1, v2 Point) Point {
	return NewPoint(v1.X+v2.X, v1.Y+v2.Y, v1.Z+v2.Z)
}

//SubPoint Function
func SubPoint(v1, v2 Point) Point {
	return NewPoint(v1.X-v2.X, v1.Y-v2.Y, v1.Z-v2.Z)
}

//Add func
func (p1 *Point) Add(tuple interface{}) (Point, error) {
	if isVector(tuple) {
		tp := Tuple(tuple.(Vector))
		pTemp := Tuple(*p1)
		newPoint := Point(AddTuple(&pTemp, &tp))
		return newPoint, nil
	} else if isPoint(tuple) {
		return Point{}, errors.New("Cant add Point")
	} else {
		return Point{}, errors.New("wrong type")
	}

	/*if isVector(tuple) {
		tu1 := tuple.(Vector)
		return AddVectors(&tu1, &tu2)
	} else if isPoint(t1) && isVector(t2) {
		tu1 := Tuple(t1.(Point))
		tu2 := Tuple(t2.(Vector))
		res := Point(AddTuple(&tu1, &tu2))
		return res
	}
	return nil*/
}

//Subtract func
func (p1 *Point) Subtract(tuple interface{}) (interface{}, error) {
	if isVector(tuple) {
		tp := Tuple(tuple.(Vector))
		pTemp := Tuple(*p1)
		newPoint := Point(SubTuple(&pTemp, &tp))
		return newPoint, nil
	} else if isPoint(tuple) {
		tp := Tuple(tuple.(Point))
		pTemp := Tuple(*p1)
		newVector := Vector(SubTuple(&pTemp, &tp))
		return newVector, nil
	} else {
		return Point{}, errors.New("wrong type")
	}
}
