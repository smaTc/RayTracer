package graphics

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
