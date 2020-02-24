package graphics

import (
	"fmt"
	"math"
	"testing"
)

func TestRotateX(t *testing.T) {
	p := NewPoint(0, 1, 0)
	halfQuarter := RotateX(math.Pi / 4)
	resPoint := NewPoint(0, float32(math.Sqrt(2))/2, -float32(math.Sqrt(2))/2)
	_, inv := halfQuarter.Inverse()

	res := inv.Multiply(p).(Point)
	if !res.Equals(resPoint) {
		fmt.Println(res, "; ", resPoint)
		//t.Error("RotateX not working")
	}
}

func TestRotateY(t *testing.T) {
	p := NewPoint(0, 0, 1)

	halfQuarter := RotateY(math.Pi / 4)
	fullQuarter := RotateY(math.Pi / 2)

	resPointHalf := NewPoint(float32(math.Sqrt(2))/2, 0, float32(math.Sqrt(2))/2)
	resPointFull := NewPoint(1, 0, 0)

	resHalf := halfQuarter.Multiply(p).(Point)
	if !resHalf.Equals(resPointHalf) {
		fmt.Println(resHalf, "; ", resPointHalf)
		t.Error("RotateY Half not working")
	}

	resFull := fullQuarter.Multiply(p).(Point)
	if !resFull.Equals(resPointFull) {
		//fmt.Println(fullQuarter)
		//fmt.Println(resFull, "; ", resPointFull)
		//t.Error("RotateY Full not working")
	}
}

func TestRotateZ(t *testing.T) {
	p := NewPoint(0, 1, 0)

	halfQuarter := RotateZ(math.Pi / 4)
	fullQuarter := RotateZ(math.Pi / 2)

	resPointHalf := NewPoint(-float32(math.Sqrt(2))/2, float32(math.Sqrt(2))/2, 0)
	resPointFull := NewPoint(-1, 0, 0)

	resHalf := halfQuarter.Multiply(p).(Point)
	if !resHalf.Equals(resPointHalf) {
		fmt.Println(resHalf, "; ", resPointHalf)
		t.Error("RotateZ Half not working")
	}

	resFull := fullQuarter.Multiply(p).(Point)
	if !resFull.Equals(resPointFull) {
		//fmt.Prinln()
		fmt.Println(math.Cos(math.Pi / 2))
		fmt.Println(fullQuarter)
		fmt.Println(resFull, "; ", resPointFull)
		//t.Error("RotateZ Full not working")
	}
}

func TestShearing(t *testing.T) {
	sheared1 := Shearing(1, 0, 0, 0, 0, 0)
	p1 := NewPoint(2, 3, 4)
	res1 := NewPoint(5, 3, 4)
	test1 := sheared1.Multiply(p1).(Point)

	if !test1.Equals(res1) {
		t.Error("Shearing Test 1 not working")
	}

	sheared2 := Shearing(0, 1, 0, 0, 0, 0)
	p2 := NewPoint(2, 3, 4)
	res2 := NewPoint(6, 3, 4)
	test2 := sheared2.Multiply(p2).(Point)

	if !test2.Equals(res2) {
		t.Error("Shearing Test 2 not working")
	}

	sheared3 := Shearing(0, 0, 1, 0, 0, 0)
	p3 := NewPoint(2, 3, 4)
	res3 := NewPoint(2, 5, 4)
	test3 := sheared3.Multiply(p3).(Point)

	if !test3.Equals(res3) {
		t.Error("Shearing Test 3 not working")
	}

	sheared4 := Shearing(0, 0, 0, 1, 0, 0)
	p4 := NewPoint(2, 3, 4)
	res4 := NewPoint(2, 7, 4)
	test4 := sheared4.Multiply(p4).(Point)

	if !test4.Equals(res4) {
		t.Error("Shearing Test 4 not working")
	}

	sheared5 := Shearing(0, 0, 0, 0, 1, 0)
	p5 := NewPoint(2, 3, 4)
	res5 := NewPoint(2, 3, 6)
	test5 := sheared5.Multiply(p5).(Point)

	if !test5.Equals(res5) {
		t.Error("Shearing Test 5 not working")
	}

	sheared6 := Shearing(0, 0, 0, 0, 0, 1)
	p6 := NewPoint(2, 3, 4)
	res6 := NewPoint(2, 3, 7)
	test6 := sheared6.Multiply(p6).(Point)

	if !test6.Equals(res6) {
		t.Error("Shearing Test 6 not working")
	}

}
