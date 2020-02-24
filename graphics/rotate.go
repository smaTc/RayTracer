package graphics

import "math"

func rotateX(radian float32) Matrix {
	var rotMatX Matrix
	rad64 := float64(radian)
	rotMatX.GenerateMatrixWithDimension(4, 4)

	rotMatX.Values = [][]float32{
		{1, 0, 0, 0},
		{0, float32(math.Cos(rad64)), -float32(math.Sin(rad64)), 0},
		{0, float32(math.Sin(rad64)), float32(math.Cos(rad64)), 0},
		{0, 0, 0, 1},
	}
	return rotMatX
}

func rotateY(radian float32) Matrix {
	var rotMatY Matrix
	rad64 := float64(radian)
	rotMatY.GenerateMatrixWithDimension(4, 4)

	rotMatY.Values = [][]float32{
		{float32(math.Cos(rad64)), 0, float32(math.Sin(rad64)), 0},
		{0, 1, 0, 0},
		{-float32(math.Sin(rad64)), 0, float32(math.Cos(rad64)), 0},
		{0, 0, 0, 1},
	}
	return rotMatY
}

func rotateZ(radian float64) Matrix {
	var rotMatZ Matrix
	//rad64 := float64(radian)
	rotMatZ.GenerateMatrixWithDimension(4, 4)

	rotMatZ.Values = [][]float32{
		{float32(math.Cos(radian)), -float32(math.Sin(radian)), 0, 0},
		{float32(math.Sin(radian)), float32(math.Cos(radian)), 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
	return rotMatZ
}
