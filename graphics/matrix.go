package graphics

import "fmt"

//Matrix struct
type Matrix struct {
	Values   [][]float64
	Identity [][]float64
}

//GenerateMatrixWithDimension func
func (m *Matrix) GenerateMatrixWithDimension(rows, columns int) {
	mat := make([][]float64, rows)
	for i := 0; i < len(mat); i++ {
		mat[i] = make([]float64, columns)
		//id[i] = make([]float64, columns)
		//id[i][i] = 1
	}
	m.Values = mat
	//m.Identity = id
	m.generateIdentity()
}

func (m *Matrix) generateIdentity() {
	fmt.Println("generating Identity...")
	if m.Values == nil || len(m.Values) != len(m.Values[0]) {
		return
	}
	fmt.Println("current Matrix:", m.Values)
	id := make([][]float64, len(m.Values))
	fmt.Println("Rows generated:", id)

	fmt.Println("generating Columns")
	for i := 0; i < len(m.Values); i++ {

		id[i] = make([]float64, len(m.Values[0]))

		id[i][i] = 1

		fmt.Println(id, "; iteration:", i)
	}

	m.Identity = id
}

//SetMatrix func
func (m *Matrix) SetMatrix(mat [][]float64) {
	m.Values = mat
	m.generateIdentity()
}

//Equals func
func (m *Matrix) Equals(values [][]float64) bool {
	//fmt.Println("Comparing", m.Values, "and", mat.Values)
	if len(m.Values) != len(values) {
		fmt.Println("length wrong")
		return false
	}

	for i := 0; i < len(m.Values); i++ {
		for j := 0; j < len(m.Values[i]); j++ {
			if m.Values[i][j] != values[i][j] {
				//fmt.Println("Values not matching")
				//fmt.Println(m.Values[i][j], "!=", mat.Values[i][j])
				return false
			}
		}
	}
	return true
}

func (m *Matrix) rowTimesColumn(row, column []float64) float64 {
	var result float64
	for i := 0; i < len(row); i++ {
		result += row[i] * column[i]
	}
	return result
}

//GetRow func
func (m *Matrix) GetRow(rw int) []float64 {
	return m.Values[rw]
}

//GetColumn func
func (m *Matrix) GetColumn(cl int) []float64 {
	var column []float64
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
	fmt.Println(prod)
	return prod
}
