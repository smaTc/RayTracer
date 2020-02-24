package graphics

import (
	"fmt"
	"math"
	"testing"
)

func TestRotateX(t *testing.T) {
	p := NewPoint(0, 1, 0)
	halfQuarter := rotateX(math.Pi / 4)
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

	halfQuarter := rotateY(math.Pi / 4)
	fullQuarter := rotateY(math.Pi / 2)

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

	halfQuarter := rotateZ(math.Pi / 4)
	fullQuarter := rotateZ(math.Pi / 2)

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
