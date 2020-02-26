package graphics

//Ray struct
type Ray struct {
	Origin        Point
	Direction     Vector
	Intersections []Intersection
}

//NewRay func
func NewRay(o Point, d Vector) Ray {
	return Ray{Origin: o, Direction: d, Intersections: make([]Intersection, 0)}
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
func (r *Ray) Intersects(s Sphere) []Intersection {
	intersecs := Intersect(*r, s)
	if len(intersecs) != 0 {
		r.Intersections = append(r.Intersections, intersecs...)
	}

	return intersecs
}
