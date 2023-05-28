package matrix

import "math"

// Translate returns a Matrix representing (x,y,z) translation.
func Translate(x, y, z float64) Matrix {
	return Matrix{
		{1, 0, 0, x},
		{0, 1, 0, y},
		{0, 0, 1, z},
		{0, 0, 0, 1},
	}
}

// Scale returns a Matrix representing (x,y,z) resizing.
func Scale(x, y, z float64) Matrix {
	return Matrix{
		{x, 0, 0, 0},
		{0, y, 0, 0},
		{0, 0, z, 0},
		{0, 0, 0, 1},
	}
}

// RotateX returns a Matrix representing a rotation about the X-axis.
// The rotation is centered at the origin and in radians.
func RotateX(rads float64) Matrix {
	return Matrix{
		{1, 0, 0, 0},
		{0, math.Cos(rads), -math.Sin(rads), 0},
		{0, math.Sin(rads), math.Cos(rads), 0},
		{0, 0, 0, 1},
	}
}

// RotateY returns a Matrix representing a rotation about the Y-axis.
// The rotation is centered at the origin and in radians.
func RotateY(rads float64) Matrix {
	return Matrix{
		{math.Cos(rads), 0, math.Sin(rads), 0},
		{0, 1, 0, 0},
		{-math.Sin(rads), 0, math.Cos(rads), 0},
		{0, 0, 0, 1},
	}
}

// RotateZ returns a Matrix representing a rotation about the Z-axis.
// The rotation is centered at the origin and in radians.
func RotateZ(rads float64) Matrix {
	return Matrix{
		{math.Cos(rads), -math.Sin(rads), 0, 0},
		{math.Sin(rads), math.Cos(rads), 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
}

// Rotate returns a Matrix representing a a rotation about all three axes.
// The rotation is centered at the origin and in radians.
func Rotate(x, y, z float64) Matrix {
	return Chain(
		RotateX(x),
		RotateY(y),
		RotateZ(z),
	)
}

// Shear returns a matrix representing a shear.
// See https://en.wikipedia.org/wiki/Transformation_matrix#Shearing
func Shear(xy, xz, yx, yz, zx, zy float64) Matrix {
	return Matrix{
		{1, xy, xz, 0},
		{yx, 1, yz, 0},
		{zx, zy, 1, 0},
		{0, 0, 0, 1},
	}
}

// Chain takes any number of Matrices and returns a composite Matrix.
// Matrices are applied to a ray.Intersectable in the order they are passed.
func Chain(tforms ...Matrix) Matrix {
	final := I()
	for _, t := range tforms {
		final = Mult(t, final)
	}
	return final
}
