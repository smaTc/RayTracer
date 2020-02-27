package graphics

import (
	"fmt"
	"math"
	"reflect"
	"sort"
)

//Intersection struct
type Intersection struct {
	T      float32
	Object interface{}
}

//Equals func
func (is *Intersection) Equals(is2 Intersection) bool {
	if is.T != is2.T || !reflect.DeepEqual(is.Object, is2.Object) {
		return false
	}
	//fmt.Println("types:", reflect.TypeOf(is.Object), reflect.TypeOf(is2.Object))

	return true
}

//NewIntersection func
func NewIntersection(t float32, o interface{}) Intersection {
	return Intersection{T: t, Object: o}
}

//Intersections func
func Intersections(i1 ...Intersection) []Intersection {
	intersections := make([]Intersection, 0)
	intersections = append(intersections, i1...)
	return intersections
}

//Intersect func
func Intersect(r Ray, s Sphere) []Intersection {
	_, rInv := s.Transform.Inverse()
	rTrans := r.Transform(rInv)
	discr, a, b := Discriminant(rTrans, s)
	if discr < 0 {
		return nil
	}

	t1 := (-b - float32(math.Sqrt(float64(discr)))) / (2 * a)
	t2 := (-b + float32(math.Sqrt(float64(discr)))) / (2 * a)
	fmt.Println("intersections:", t1, ",", t2)

	intersec1 := Intersection{T: t1, Object: s}
	intersec2 := Intersection{T: t2, Object: s}

	if t1 <= t2 {
		return []Intersection{intersec1, intersec2}
	}
	return []Intersection{intersec2, intersec1}
}

//Discriminant func
func Discriminant(r Ray, s Sphere) (float32, float32, float32) {
	sphereToRay, _ := r.Origin.Subtract(s.Origin)
	//phereToRay, _ := s.Origin.Subtract(r.Origin)

	sphereToRayVector := sphereToRay.(Vector)
	//fmt.Println("SpheretoRay", sphereToRayVector)
	a := Dot(&r.Direction, &r.Direction)
	b := 2 * Dot(&r.Direction, &sphereToRayVector)
	c := Dot(&sphereToRayVector, &sphereToRayVector) - 1

	discriminant := b*b - 4*a*c
	//fmt.Println("dab", discriminant, a, b)
	return discriminant, a, b
}

//Hit func
func Hit(intersecs []Intersection) Intersection {
	//fmt.Println("unsorted", intersecs)
	sort.Slice(intersecs, func(i, j int) bool {
		return intersecs[i].T < intersecs[j].T
	})
	//fmt.Println("sorted", intersecs)

	hitIndex := -1

	for i, is := range intersecs {
		if is.T >= 0 {
			hitIndex = i
			break
		}
	}

	if hitIndex > -1 {
		//fmt.Println("hit:", intersecs[hitIndex])
		return intersecs[hitIndex]
	}

	return Intersection{}
}
