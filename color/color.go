// Package color implements an RGB color value and provides common utilities.
package color

import (
	"github.com/noahssarcastic/gort/math"
)

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
	return math.FloatEqual(c1.r, c2.r) &&
		math.FloatEqual(c1.g, c2.g) &&
		math.FloatEqual(c1.b, c2.b)
}

// Add returns the component-wise sum of two Colors.
func Add(c1, c2 Color) Color {
	return Color{
		r: c1.r + c2.r,
		b: c1.b + c2.b,
		g: c1.g + c2.g,
	}
}

// Sub returns the component-wise difference of two Colors.
func Sub(c1, c2 Color) Color {
	return Color{
		r: c1.r - c2.r,
		b: c1.b - c2.b,
		g: c1.g - c2.g,
	}
}

// Mult scales each component of a Color c by the given scalar.
func Mult(c Color, scalar float64) Color {
	return Color{
		r: c.r * scalar,
		b: c.b * scalar,
		g: c.g * scalar,
	}
}

// PiecewiseMult returns the component-wise product of two Colors.
func PiecewiseMult(c1 Color, c2 Color) Color {
	return Color{
		r: c1.r * c2.r,
		b: c1.b * c2.b,
		g: c1.g * c2.g,
	}
}

func White() Color { return Color{1, 1, 1} }

func Red() Color { return Color{1, 0, 0} }

func Green() Color { return Color{0, 1, 0} }

func Blue() Color { return Color{0, 0, 1} }

func Black() Color { return Color{0, 0, 0} }

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
