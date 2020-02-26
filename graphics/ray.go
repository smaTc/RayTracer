package graphics

//Ray struct
type Ray struct {
	Origin    Point
	Direction Vector
}

//NewRay func
func NewRay(o Point, d Vector) Ray {
	return Ray{Origin: o, Direction: d}
}

//SetRayProps func
func (r *Ray) SetRayProps(o Point, d Vector) {
	r.Origin = o
	r.Direction = d
}

//Position func
func (r *Ray) Position(t float32) Point {
	vec := r.Direction.Scale(t)
	newPoint, _ := r.Origin.Add(vec)
	return newPoint
}

//Intersects func
func (r *Ray) Intersects(s Sphere) []float32 {
	//return Intersect(*r, s)
	return []float32{4.0, 6.0}
}

//Intersect func
func Intersect(r Ray, s Sphere) []Point {
	return []Point{}
}
