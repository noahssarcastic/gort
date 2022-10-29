package transform

import "github.com/noahssarcastic/tddraytracer/matrix"

func Translate(x, y, z float64) matrix.Matrix {
	return matrix.Matrix{
		{1, 0, 0, x},
		{0, 1, 0, y},
		{0, 0, 1, z},
		{0, 0, 0, 1},
	}
}

func Rotate(x, y, z float64) matrix.Matrix {
	panic("not implemented")
}

func Scale(x, y, z float64) matrix.Matrix {
	panic("not implemented")
}

func Shear(x, y, z float64) matrix.Matrix {
	panic("not implemented")
}

func ChainTransforms(tforms []matrix.Matrix) matrix.Matrix {
	panic("not implemented")
}
