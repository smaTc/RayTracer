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

//Multiply func
func (r *Ray) Multiply(m Matrix) Ray {
	//no := r.Origin.Multiply(m)
	//no := m.Multiply(r.Origin).(Point)

	//nd := r.Direction.Multiply(m)
	//nd := m.Multiply(r.Direction).(Vector)
	//fmt.Println("no,nd:", no, nd)
	//return NewRay(no, nd)
	return Ray{}
}

//Transform func
func (r *Ray) Transform(m Matrix) Ray {
	no := m.Multiply(r.Origin).(Point)
	nd := m.Multiply(r.Direction).(Vector)

	return NewRay(no, nd)

}
