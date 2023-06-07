// Package tuple implements points and vectors as a 4-tuple.
//
// Using a common representation for both types allows for simpler computations
// Use of homogeneous coordinates.
// See https://en.wikipedia.org/wiki/Homogeneous_coordinates.
package tuple

import (
	"math"

	"github.com/noahssarcastic/gort/pkg/util"
)

// Tuple implements points and vectors using a common struct.
type Tuple struct {
	X, Y, Z float64
	w       float64
}

// W returns the w component of the Tuple. A w value of 0 represents a vector
// and a value of 1 represents a point.
func (t Tuple) W() float64 {
	return t.w
}

// Equal returns true if two Tuples are equal.
func Equal(a, b Tuple) bool {
	return util.FloatEqual(a.X, b.X) &&
		util.FloatEqual(a.Y, b.Y) &&
		util.FloatEqual(a.Z, b.Z) &&
		util.FloatEqual(a.w, b.w)
}

// New creates a new 4-tuple. It is recommended to use Vector and Point
// rather than using New directly.
func New(x, y, z, w float64) Tuple {
	return Tuple{
		X: x,
		Y: y,
		Z: z,
		w: w,
	}
}

// Point creates a new point at (x,y,z).
func Point(x, y, z float64) Tuple {
	return Tuple{
		X: x,
		Y: y,
		Z: z,
		w: 1,
	}
}

// IsPoint returns true if a tuple is a point.
func (t Tuple) IsPoint() bool {
	return util.FloatEqual(t.w, 1)
}

// Vector creates a new vector starting at the origin and ending at (x,y,z).
func Vector(x, y, z float64) Tuple {
	return Tuple{
		X: x,
		Y: y,
		Z: z,
		w: 0,
	}
}

// IsVector returns true if a given Tuple is a vector.
func (t Tuple) IsVector() bool {
	return util.FloatEqual(t.w, 0)
}

// Add calculates the component-wise sum of two Tuples.
func Add(a, b Tuple) Tuple {
	return Tuple{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
		w: a.w + b.w,
	}
}

// Add calculates the component-wise sum of two Tuples.
func (t1 Tuple) Add(t2 Tuple) Tuple {
	return Add(t1, t2)
}

// Sub calculates the component-wise difference of two Tuples.
func Sub(a, b Tuple) Tuple {
	return Tuple{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
		w: a.w - b.w,
	}
}

// Sub calculates the component-wise difference of two Tuples.
func (t1 Tuple) Sub(t2 Tuple) Tuple {
	return Sub(t1, t2)
}

// Neg calculates the component-wise negation of a Tuple.
func Neg(t Tuple) Tuple {
	return Sub(Vector(0, 0, 0), t)
}

// Mult calculates the component-wise product of a Tuple and a scalar.
func Mult(t Tuple, scalar float64) Tuple {
	return Tuple{
		X: t.X * scalar,
		Y: t.Y * scalar,
		Z: t.Z * scalar,
		w: t.w * scalar,
	}
}

// Div calculates the component-wise quotient of a Tuple and a scalar.
func Div(t Tuple, scalar float64) Tuple {
	return Tuple{
		X: t.X / scalar,
		Y: t.Y / scalar,
		Z: t.Z / scalar,
		w: t.w / scalar,
	}
}

// Mag calculates the magnitude of a vector. Passing a point to Mag is
// undefined result.
func Mag(vec Tuple) float64 {
	return math.Sqrt(
		math.Pow(vec.X, 2) +
			math.Pow(vec.Y, 2) +
			math.Pow(vec.Z, 2) +
			math.Pow(vec.w, 2))
}

// Norm returns a new vector which is the norm of vec. Passing a point to Norm
// is undefined.
func Norm(vec Tuple) Tuple {
	return Div(vec, Mag(vec))
}

// Dot returns the dot-product of two vectors. Passing a point to Dot
// is undefined.
func Dot(a, b Tuple) float64 {
	return a.X*b.X +
		a.Y*b.Y +
		a.Z*b.Z +
		a.w*b.w
}

// Cross returns the cross-product of two vectors. Passing a point to Cross
// is undefined.
func Cross(a, b Tuple) Tuple {
	return Vector(
		a.Y*b.Z-a.Z*b.Y,
		a.Z*b.X-a.X*b.Z,
		a.X*b.Y-a.Y*b.X,
	)
}

// Reflect returns the reflection of vector in along vector norm.
func Reflect(in, norm Tuple) Tuple {
	return in.Sub(Mult(norm, 2*Dot(in, norm)))
}
