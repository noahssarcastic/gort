package math

import (
	stdmath "math"
)

type Tuple struct {
	x, y, z, w float64
}

func (t Tuple) X() float64 {
	return t.x
}

func (t Tuple) Y() float64 {
	return t.y
}

func (t Tuple) Z() float64 {
	return t.z
}

func (t Tuple) W() float64 {
	return t.w
}

func TupleEqual(a, b Tuple) bool {
	return FloatEqual(a.x, b.x) &&
		FloatEqual(a.y, b.y) &&
		FloatEqual(a.z, b.z) &&
		FloatEqual(a.w, b.w)
}

func NewTuple(x, y, z, w float64) Tuple {
	return Tuple{
		x: x,
		y: y,
		z: z,
		w: w,
	}
}

func Point(x, y, z float64) Tuple {
	return Tuple{
		x: x,
		y: y,
		z: z,
		w: 1,
	}
}

func (t Tuple) IsPoint() bool {
	return FloatEqual(t.w, 1)
}

func Vector(x, y, z float64) Tuple {
	return Tuple{
		x: x,
		y: y,
		z: z,
		w: 0,
	}
}

func (t Tuple) IsVector() bool {
	return FloatEqual(t.w, 0)
}

func (a Tuple) Add(b Tuple) Tuple {
	return Tuple{
		x: a.x + b.x,
		y: a.y + b.y,
		z: a.z + b.z,
		w: a.w + b.w,
	}
}

func (a Tuple) Sub(b Tuple) Tuple {
	return Tuple{
		x: a.x - b.x,
		y: a.y - b.y,
		z: a.z - b.z,
		w: a.w - b.w,
	}
}

func Neg(t Tuple) Tuple {
	return Vector(0, 0, 0).Sub(t)
}

func (t Tuple) Mult(scalar float64) Tuple {
	return Tuple{
		x: t.x * scalar,
		y: t.y * scalar,
		z: t.z * scalar,
		w: t.w * scalar,
	}
}

func (t Tuple) Div(scalar float64) Tuple {
	return Tuple{
		x: t.x / scalar,
		y: t.y / scalar,
		z: t.z / scalar,
		w: t.w / scalar,
	}
}

func Mag(t Tuple) float64 {
	return stdmath.Sqrt(
		stdmath.Pow(t.x, 2) +
			stdmath.Pow(t.y, 2) +
			stdmath.Pow(t.z, 2) +
			stdmath.Pow(t.w, 2))
}

func Norm(t Tuple) Tuple {
	return t.Div(Mag(t))
}

func Dot(a, b Tuple) float64 {
	return a.x*b.x +
		a.y*b.y +
		a.z*b.z +
		a.w*b.w
}

func Cross(a, b Tuple) Tuple {
	return Vector(
		a.y*b.z-a.z*b.y,
		a.z*b.x-a.x*b.z,
		a.x*b.y-a.y*b.x,
	)
}
