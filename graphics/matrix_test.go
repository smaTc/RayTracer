package graphics

import (
	"fmt"
	"testing"
)

func TestMatrix(t *testing.T) {
	var m1, m2, m3, m4, id4 Matrix

	m1.SetMatrix(
		[][]float64{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 1, 2, 3},
			{4, 5, 6, 7},
		})

	m2.SetMatrix(
		[][]float64{
			{9, 8, 7, 6},
			{5, 4, 3, 2},
			{1, 9, 8, 7},
			{6, 5, 4, 3},
		})

	m3.SetMatrix(
		[][]float64{
			{46, 63, 53, 43},
			{130, 167, 141, 115},
			{106, 109, 94, 79},
			{109, 141, 119, 97},
		})

	m4.SetMatrix(
		[][]float64{
			{1},
			{2},
			{3},
			{4},
		})

	ident4Vals := [][]float64{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
	id4.Values = ident4Vals

	if !id4.Equals(m1.Identity) {
		t.Error("Identity not working")
	}

	col0 := m1.GetColumn(0)
	testCol := []float64{1, 5, 9, 4}

	for i := 0; i < len(col0); i++ {
		if col0[i] != testCol[i] {
			t.Error("Get Column not working")
		}
	}

	row0 := m2.GetRow(1)
	testRow := []float64{5, 4, 3, 2}

	for i := 0; i < len(row0); i++ {
		if row0[i] != testRow[i] {
			t.Error("Get Row not working")
		}
	}

	if m1.Equals(m2.Values) || !m1.Equals(m1.Values) {
		t.Error("Matrix equal not working")
	}

	mTest := m1.Multiply(&m2)
	//fmt.Println(mTest, m3)
	if !mTest.Equals(m3.Values) {
		t.Error("Matrix multiplication not working")
	}

	var mRes Matrix
	mRes.Values = [][]float64{
		{30},
		{70},
		{29},
		{60},
	}
	mTestX := m1.Multiply(&m4)
	fmt.Println(mTestX)
	if !mTestX.Equals(mRes.Values) {
		t.Error("matrix mult with vector not working")
	}
}
