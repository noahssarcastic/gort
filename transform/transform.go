package transform

import (
	"math"

	"github.com/noahssarcastic/tddraytracer/matrix"
)

func Translate(x, y, z float64) matrix.Matrix {
	return matrix.Matrix{
		{1, 0, 0, x},
		{0, 1, 0, y},
		{0, 0, 1, z},
		{0, 0, 0, 1},
	}
}

func Scale(x, y, z float64) matrix.Matrix {
	return matrix.Matrix{
		{x, 0, 0, 0},
		{0, y, 0, 0},
		{0, 0, z, 0},
		{0, 0, 0, 1},
	}
}

func rotateX(rads float64) matrix.Matrix {
	return matrix.Matrix{
		{1, 0, 0, 0},
		{0, math.Cos(rads), -math.Sin(rads), 0},
		{0, math.Sin(rads), math.Cos(rads), 0},
		{0, 0, 0, 1},
	}
}

func rotateY(rads float64) matrix.Matrix {
	return matrix.Matrix{
		{math.Cos(rads), 0, math.Sin(rads), 0},
		{0, 1, 0, 0},
		{-math.Sin(rads), 0, math.Cos(rads), 0},
		{0, 0, 0, 1},
	}
}

func rotateZ(rads float64) matrix.Matrix {
	return matrix.Matrix{
		{math.Cos(rads), -math.Sin(rads), 0, 0},
		{math.Sin(rads), math.Cos(rads), 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
}

func Rotate(x, y, z float64) matrix.Matrix {
	panic("not implemented")
}

func Shear(xy, xz, yx, yz, zx, zy float64) matrix.Matrix {
	return matrix.Matrix{
		{1, xy, xz, 0},
		{yx, 1, yz, 0},
		{zx, zy, 1, 0},
		{0, 0, 0, 1},
	}
}

func ChainTransforms(tforms []matrix.Matrix) matrix.Matrix {
	panic("not implemented")
}
