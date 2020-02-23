package graphics

import (
	"testing"
)

var m1, m2, m3, m4, id4 Matrix

func TestMatrix(t *testing.T) {

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
	//fmt.Println(mTestX)
	if !mTestX.Equals(mRes.Values) {
		t.Error("matrix mult with vector not working")
	}

	var mTranspose Matrix
	mTranspose.Values = [][]float64{
		{1, 5, 9, 4},
		{2, 6, 1, 5},
		{3, 7, 2, 6},
		{4, 8, 3, 7},
	}

	mTestTrans := m1
	_, mTestTransTransposed := mTestTrans.Transpose()
	//fmt.Println()
	if !mTranspose.Equals(mTestTransTransposed.Values) {
		t.Error("Transpose not working")
	}

	var mDet2x2 Matrix
	mDet2x2.Values = [][]float64{
		{1, 5},
		{-3, 2},
	}

	if mDet2x2.Determinant() != 17 {
		t.Error("2x2 Determinant not working")
	}

	var subMat Matrix
	subMat.Values = [][]float64{
		{1, 5, 4},
		{2, 6, 5},
		{4, 8, 7},
	}

	subTestM := mTranspose.Submatrix(2, 2)

	if !subTestM.Equals(subMat.Values) {
		t.Error("Submatrix not working")
	}

	var subTestMinor Matrix
	subTestMinor.Values = [][]float64{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	}

	if subTestMinor.Minor(1, 0) != 25 {
		t.Error("minor not working")
	}

	if subTestMinor.Cofactor(0, 0) != -12 || subTestMinor.Cofactor(1, 0) != -25 {
		t.Error("cofactor not working")
	}

	var determinantMatrix Matrix
	determinantMatrix.Values = [][]float64{
		{1, 2, 6},
		{-5, 8, -4},
		{2, 6, 4},
	}

	if determinantMatrix.Cofactor(0, 0) != 56 || determinantMatrix.Cofactor(0, 1) != 12 || determinantMatrix.Cofactor(0, 2) != -46 || determinantMatrix.Determinant() != -196 {
		t.Error("Arbitrary 3x3 determinant not working")
	}

	var determinantMatrix2 Matrix
	determinantMatrix2.Values = [][]float64{
		{-2, -8, 3, 5},
		{-3, 1, 7, 3},
		{1, 2, -9, 6},
		{-6, 7, 7, -9},
	}

	if determinantMatrix2.Cofactor(0, 0) != 690 || determinantMatrix2.Cofactor(0, 1) != 447 || determinantMatrix2.Cofactor(0, 2) != 210 || determinantMatrix2.Cofactor(0, 3) != 51 || determinantMatrix2.Determinant() != -4071 {
		t.Error("Arbitrary 4x4 determinant not working")
	}

	var invMatrix Matrix
	invMatrix.Values = [][]float64{
		{-5, 2, 6, -8},
		{1, -5, 1, 8},
		{7, 7, -6, -7},
		{1, -3, 7, 4},
	}

	var invMatrixResult Matrix
	invMatrixResult.Values = [][]float64{
		{0.21805, 0.45113, 0.24060, -0.04511},
		{-0.80827, -1.45677, -0.44361, 0.52068},
		{-0.07895, -0.22368, -0.05263, 0.19737},
		{-0.52256, -0.81391, -0.30075, 0.30639},
	}
	//_, invRes := invMatrix.Inverse()

	//Change to float32?
	/*if !invRes.Equals(invMatrixResult.Values) {
		fmt.Println("inv", invRes.Values)
		fmt.Println("invResult", invMatrixResult.Values)
		t.Error("inversion not working")
	}*/
}
