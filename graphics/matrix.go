package graphics

import (
	"math"
)

//Matrix struct
type Matrix struct {
	Values   [][]float32
	Identity [][]float32
}

//GetDimensions func
func (m *Matrix) GetDimensions() (int, int) {
	rows := len(m.Values)
	columns := len(m.Values[0])

	return rows, columns
}

//GenerateMatrixWithDimension func
func (m *Matrix) GenerateMatrixWithDimension(rows, columns int) {
	mat := make([][]float32, rows)
	for i := 0; i < len(mat); i++ {
		mat[i] = make([]float32, columns)
	}
	m.Values = mat
	m.generateIdentity()
}

func (m *Matrix) generateIdentity() {
	if m.Values == nil || m.Values[0] == nil || len(m.Values) != len(m.Values[0]) {
		return
	}
	id := make([][]float32, len(m.Values))
	for i := 0; i < len(m.Values); i++ {

		id[i] = make([]float32, len(m.Values[0]))

		id[i][i] = 1
	}
	m.Identity = id
}

//SetMatrix func
func (m *Matrix) SetMatrix(mat [][]float32) {
	m.Values = mat
	m.generateIdentity()
}

//Equals func
func (m *Matrix) Equals(values [][]float32) bool {
	if len(m.Values) != len(values) {
		return false
	}

	for i := 0; i < len(m.Values); i++ {
		for j := 0; j < len(m.Values[i]); j++ {
			if m.Values[i][j] != values[i][j] {
				return false
			}
		}
	}
	return true
}

func (m *Matrix) rowTimesColumn(row, column []float32) float32 {
	var result float32
	for i := 0; i < len(row); i++ {
		result += row[i] * column[i]
	}
	return result
}

//GetRow func
func (m *Matrix) GetRow(rw int) []float32 {
	return m.Values[rw]
}

//GetColumn func
func (m *Matrix) GetColumn(cl int) []float32 {
	var column []float32
	for _, row := range m.Values {
		column = append(column, row[cl])
	}
	return column
}

//Multiply func
func (m *Matrix) Multiply(mat *Matrix) Matrix {
	var prod Matrix
	prod.GenerateMatrixWithDimension(len(m.Values), len(mat.Values[0]))

	for i := 0; i < len(prod.Values); i++ {
		for j := 0; j < len(prod.Values[0]); j++ {
			prod.Values[i][j] = prod.rowTimesColumn(m.GetRow(i), mat.GetColumn(j))
		}
	}
	return prod
}

//Transpose func
func (m *Matrix) Transpose() (bool, Matrix) {
	if m.Values == nil || m.Values[0] == nil {
		return false, Matrix{}
	}

	//mat := make([][]float32, len(m.Values))
	var mat Matrix
	r, c := m.GetDimensions()
	mat.GenerateMatrixWithDimension(r, c)

	for i := 0; i < len(m.Values); i++ {
		//mat.Values[i] = make([]float32, len(m.Values[0]))
		mat.Values[i] = m.GetColumn(i)
	}
	//m.Values = mat
	return true, mat
}

//Determinant func
func (m *Matrix) Determinant() float32 {
	if len(m.Values) == 2 && len(m.Values[0]) == 2 {
		return (m.Values[0][0] * m.Values[1][1]) - (m.Values[0][1] * m.Values[1][0])
	}

	var det float32
	for i := 0; i < len(m.Values[0]); i++ {
		det += m.Values[0][i] * m.Cofactor(0, i)
	}
	return det
}

//Submatrix func
func (m *Matrix) Submatrix(jumpRow, jumpColumn int) Matrix {
	mat := *m
	rows, cols := m.GetDimensions()
	mat.GenerateMatrixWithDimension(rows-1, cols-1)

	for i, matRow := 0, 0; i < len(m.Values); i++ {
		if i == jumpRow {
			continue
		}
		for j, matCol := 0, 0; j < len(m.Values[0]); j++ {
			if j != jumpColumn {
				mat.Values[matRow][matCol] = m.Values[i][j]
				matCol++
			}
		}
		matRow++
	}

	return mat
}

//Minor func
func (m *Matrix) Minor(row, column int) float32 {
	mat := m.Submatrix(row, column)
	return mat.Determinant()
}

//Cofactor func
func (m *Matrix) Cofactor(row, column int) float32 {
	sign := math.Pow(-1, float64(row+column))
	//fmt.Println("Row+Columns:", row+column, "; sign:", sign)
	return float32(sign) * m.Minor(row, column)
}

//Inverse func
func (m *Matrix) Inverse() (bool, Matrix) {
	if m.Determinant() == 0 {
		return false, Matrix{}
	}

	var cofMat, invMat Matrix
	r, c := m.GetDimensions()
	cofMat.GenerateMatrixWithDimension(r, c)
	invMat.GenerateMatrixWithDimension(r, c)

	for i := 0; i < len(m.Values); i++ {
		for j := 0; j < len(m.Values[i]); j++ {
			cofMat.Values[i][j] = m.Cofactor(i, j)
		}
	}
	_, transpCofMat := cofMat.Transpose()

	for i := 0; i < len(m.Values); i++ {
		for j := 0; j < len(m.Values[i]); j++ {
			invMat.Values[i][j] = transpCofMat.Values[i][j] / m.Determinant()
		}
	}

	return true, invMat
}
