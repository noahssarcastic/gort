package tuple

import "math"

const EPSILON float64 = 0.00001

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

func FloatEqual(a, b float64) bool {
	return math.Abs(a-b) < EPSILON
}

func Equal(a, b Tuple) bool {
	return FloatEqual(a.x, b.x) &&
		FloatEqual(a.y, b.y) &&
		FloatEqual(a.z, b.z) &&
		FloatEqual(a.w, b.w)
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

func (a Tuple) Subtract(b Tuple) Tuple {
	return Tuple{
		x: a.x - b.x,
		y: a.y - b.y,
		z: a.z - b.z,
		w: a.w - b.w,
	}
}

func Negate(t Tuple) Tuple {
	return Vector(0, 0, 0).Subtract(t)
}

func (t Tuple) Multiply(scalar float64) Tuple {
	return Tuple{
		x: t.x * scalar,
		y: t.y * scalar,
		z: t.z * scalar,
		w: t.w * scalar,
	}
}

func (t Tuple) Divide(scalar float64) Tuple {
	return Tuple{
		x: t.x / scalar,
		y: t.y / scalar,
		z: t.z / scalar,
		w: t.w / scalar,
	}
}

func Magnitude(t Tuple) float64 {
	return math.Sqrt(
		math.Pow(t.x, 2) +
			math.Pow(t.y, 2) +
			math.Pow(t.z, 2) +
			math.Pow(t.w, 2))
}

func Normalize(t Tuple) Tuple {
	return t.Divide(Magnitude(t))
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
