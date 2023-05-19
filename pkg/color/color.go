// Package color implements an RGB color value and provides common utilities.
package color

import "github.com/noahssarcastic/gort/pkg/util"

// A Color represents an RGB color with component values between [0,1].
// The zero value for Color is black.
type Color struct {
	r, g, b float64
}

// New returns a new Color literal.
func New(r, g, b float64) Color {
	return Color{r, g, b}
}

// Get the R component value of a Color.
func (c Color) R() float64 { return c.r }

// Get the G component value of a Color.
func (c Color) G() float64 { return c.g }

// Get the B component value of a Color.
func (c Color) B() float64 { return c.b }

// Equal returns true if the given Colors are equal.
func Equal(c1, c2 Color) bool {
	return util.FloatEqual(c1.r, c2.r) &&
		util.FloatEqual(c1.g, c2.g) &&
		util.FloatEqual(c1.b, c2.b)
}

// Add returns the component-wise sum of two Colors.
func Add(c1, c2 Color) Color {
	return Color{
		r: c1.r + c2.r,
		g: c1.g + c2.g,
		b: c1.b + c2.b,
	}
}

// Add returns the component-wise sum of two Colors.
func (c1 Color) Add(c2 Color) Color {
	return Add(c1, c2)
}

// Sub returns the component-wise difference of two Colors.
func Sub(c1, c2 Color) Color {
	return Color{
		r: c1.r - c2.r,
		g: c1.g - c2.g,
		b: c1.b - c2.b,
	}
}

// Sub returns the component-wise difference of two Colors.
func (c1 Color) Sub(c2 Color) Color {
	return Sub(c1, c2)
}

// Mult scales each component of a Color c by the given scalar.
func Mult(c Color, scalar float64) Color {
	return Color{
		r: c.r * scalar,
		g: c.g * scalar,
		b: c.b * scalar,
	}
}

// PiecewiseMult returns the component-wise product of two Colors.
func PiecewiseMult(c1 Color, c2 Color) Color {
	return Color{
		r: c1.r * c2.r,
		g: c1.g * c2.g,
		b: c1.b * c2.b,
	}
}

var (
	White = Color{1, 1, 1}
	Red   = Color{1, 0, 0}
	Green = Color{0, 1, 0}
	Blue  = Color{0, 0, 1}
	Black = Color{0, 0, 0}
)

func clampValue(val float64) float64 {
	if val < 0 {
		return 0
	} else if val > 1 {
		return 1
	} else {
		return val
	}
}

// Clamp returns a Color whose component values are clamped between [0,1].
//
// If a component value is less than 0, the component is set to 0.
// If it is greater than 1, the component is set to 1.
// Otherwise, the component is left unchanged.
func Clamp(c Color) Color {
	return Color{
		clampValue(c.r),
		clampValue(c.g),
		clampValue(c.b)}
}
