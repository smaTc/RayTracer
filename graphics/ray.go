package graphics

//Ray struct
type Ray struct {
	Origin    Point
	Direction Vector
}

func NewRay(o Point, d Vector) Ray {
	return Ray{Origin: o, Direction: d}
}

func (r *Ray) SetRayProps(o Point, d Vector) {
	r.Origin = o
	r.Direction = d
}

func (r *Ray) Position(t float32) Point {
	vec := r.Direction.Scale(t)
	newPoint, _ := r.Origin.Add(vec)
	return newPoint
}
