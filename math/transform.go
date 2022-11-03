package math

import (
	stdmath "math"
)

func Translate(x, y, z float64) Matrix {
	return Matrix{
		{1, 0, 0, x},
		{0, 1, 0, y},
		{0, 0, 1, z},
		{0, 0, 0, 1},
	}
}

func Scale(x, y, z float64) Matrix {
	return Matrix{
		{x, 0, 0, 0},
		{0, y, 0, 0},
		{0, 0, z, 0},
		{0, 0, 0, 1},
	}
}

func RotateX(rads float64) Matrix {
	return Matrix{
		{1, 0, 0, 0},
		{0, stdmath.Cos(rads), -stdmath.Sin(rads), 0},
		{0, stdmath.Sin(rads), stdmath.Cos(rads), 0},
		{0, 0, 0, 1},
	}
}

func RotateY(rads float64) Matrix {
	return Matrix{
		{stdmath.Cos(rads), 0, stdmath.Sin(rads), 0},
		{0, 1, 0, 0},
		{-stdmath.Sin(rads), 0, stdmath.Cos(rads), 0},
		{0, 0, 0, 1},
	}
}

func RotateZ(rads float64) Matrix {
	return Matrix{
		{stdmath.Cos(rads), -stdmath.Sin(rads), 0, 0},
		{stdmath.Sin(rads), stdmath.Cos(rads), 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
}

func Rotate(x, y, z float64) Matrix {
	return Chain(
		RotateX(x),
		RotateY(y),
		RotateZ(z),
	)
}

func Shear(xy, xz, yx, yz, zx, zy float64) Matrix {
	return Matrix{
		{1, xy, xz, 0},
		{yx, 1, yz, 0},
		{zx, zy, 1, 0},
		{0, 0, 0, 1},
	}
}

func Chain(tforms ...Matrix) Matrix {
	final := I(4)
	for _, t := range tforms {
		final = Mult(t, final)
	}
	return final
}
