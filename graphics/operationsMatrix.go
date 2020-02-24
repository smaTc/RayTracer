package graphics

//Translation func
func Translation(x, y, z float32) Matrix {
	var mat Matrix
	mat.GenerateMatrixWithDimension(4, 4)
	mat.Values = [][]float32{
		{1, 0, 0, x},
		{0, 1, 0, y},
		{0, 0, 1, z},
		{0, 0, 0, 1},
	}
	return mat
}

//Scaling func
func Scaling(x, y, z float32) Matrix {
	var mat Matrix
	mat.GenerateMatrixWithDimension(4, 4)
	mat.Values = [][]float32{
		{x, 0, 0, 0},
		{0, y, 0, 0},
		{0, 0, z, 0},
		{0, 0, 0, 1},
	}
	return mat
}

//Shearing func
func Shearing(xy, xz, yx, yz, zx, zy float32) Matrix {
	var mat Matrix
	mat.GenerateMatrixWithDimension(4, 4)
	mat.Values = [][]float32{
		{1, xy, xz, 0},
		{yx, 1, yz, 0},
		{zx, zy, 1, 0},
		{0, 0, 0, 1},
	}
	return mat
}
