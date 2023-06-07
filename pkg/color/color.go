// Package color implements an RGB color value and provides common utilities.
package color

import "github.com/noahssarcastic/gort/pkg/util"

// A Color represents an RGB color with component values between [0,1].
// The zero value for Color is black.
type Color struct {
	R, G, B float64
}

// New returns a new Color literal.
func New(r, g, b float64) Color {
	return Color{r, g, b}
}

// Equal returns true if the given Colors are equal.
func Equal(c1, c2 Color) bool {
	return util.FloatEqual(c1.R, c2.R) &&
		util.FloatEqual(c1.G, c2.G) &&
		util.FloatEqual(c1.B, c2.B)
}

// Add returns the component-wise sum of two Colors.
func Add(c1, c2 Color) Color {
	return Color{
		R: c1.R + c2.R,
		G: c1.G + c2.G,
		B: c1.B + c2.B,
	}
}

// Add returns the component-wise sum of two Colors.
func (c1 Color) Add(c2 Color) Color {
	return Add(c1, c2)
}

// Sub returns the component-wise difference of two Colors.
func Sub(c1, c2 Color) Color {
	return Color{
		R: c1.R - c2.R,
		G: c1.G - c2.G,
		B: c1.B - c2.B,
	}
}

// Sub returns the component-wise difference of two Colors.
func (c1 Color) Sub(c2 Color) Color {
	return Sub(c1, c2)
}

// Mult scales each component of a Color c by the given scalar.
func Mult(c Color, scalar float64) Color {
	return Color{
		R: c.R * scalar,
		G: c.G * scalar,
		B: c.B * scalar,
	}
}

// PiecewiseMult returns the component-wise product of two Colors.
func PiecewiseMult(c1 Color, c2 Color) Color {
	return Color{
		R: c1.R * c2.R,
		G: c1.G * c2.G,
		B: c1.B * c2.B,
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
		clampValue(c.R),
		clampValue(c.G),
		clampValue(c.B)}
}
