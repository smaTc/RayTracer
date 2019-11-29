package graphics

//Equals checks if two Tupels,Points,Vectors are equal
func Equals(t1, t2 interface{}) bool {
	if isVector(t1) && isVector(t2) {
		v1 := t1.(Vector)
		v2 := t2.(Vector)
		return v1.Equals(v2)
	} else if isPoint(t1) && isPoint(t2) {
		p1 := t1.(Point)
		p2 := t2.(Point)
		return p1.Equals(p2)
	}
	return false
}

//Add Function
func Add(t1, t2 interface{}) interface{} {
	if isVector(t1) && isVector(t2) {
		tu1 := t1.(Vector)
		tu2 := t2.(Vector)
		return AddVectors(&tu1, &tu2)
	} else if isPoint(t1) && isVector(t2) {
		tu1 := Tuple(t1.(Point))
		tu2 := Tuple(t2.(Vector))
		res := Point(AddTuple(&tu1, &tu2))
		return res
	} else if isPoint(t2) && isVector(t1) {
		tu1 := Tuple(t2.(Point))
		tu2 := Tuple(t1.(Vector))
		res := Point(AddTuple(&tu1, &tu2))
		return res
	}
	return nil
}

//Sub Function
func Sub(t1, t2 interface{}) interface{} {
	if isVector(t1) && isVector(t2) {
		tu1 := t1.(Vector)
		tu2 := t2.(Vector)
		return SubVector(&tu1, &tu2)
	} else if isPoint(t1) && isVector(t2) {
		tu1 := Tuple(t1.(Point))
		tu2 := Tuple(t2.(Vector))
		res := Point(SubTuple(&tu1, &tu2))
		return res
	} else if isPoint(t2) && isVector(t1) {
		tu1 := Tuple(t2.(Point))
		tu2 := Tuple(t1.(Vector))
		res := Point(SubTuple(&tu1, &tu2))
		return res
	}
	return nil
}

//
/*Tuple Functions*/

//AddTuple Function
func AddTuple(t1, t2 *Tuple) Tuple {
	return NewTuple(t1.X+t2.X, t1.Y+t2.Y, t1.Z+t2.Z, t1.W+t2.W)
}

//SubTuple Function
func SubTuple(t1, t2 *Tuple) Tuple {
	return NewTuple(t1.X-t2.X, t1.Y-t2.Y, t1.Z-t2.Z, t1.W-t2.W)
}
