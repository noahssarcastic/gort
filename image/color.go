package image

import (
	"github.com/noahssarcastic/gort/math"
)

type Color struct {
	r, g, b float64
}

func NewColor(r, g, b float64) Color {
	return Color{r, g, b}
}

func (c Color) R() float64 {
	return c.r
}

func (c Color) G() float64 {
	return c.g
}

func (c Color) B() float64 {
	return c.b
}

func ColorEqual(c1, c2 Color) bool {
	return math.FloatEqual(c1.r, c2.r) &&
		math.FloatEqual(c1.g, c2.g) &&
		math.FloatEqual(c1.b, c2.b)
}

func (c1 Color) Add(c2 Color) Color {
	return Color{
		r: c1.r + c2.r,
		b: c1.b + c2.b,
		g: c1.g + c2.g,
	}
}

func (c1 Color) Sub(c2 Color) Color {
	return Color{
		r: c1.r - c2.r,
		b: c1.b - c2.b,
		g: c1.g - c2.g,
	}
}

func (c Color) Mult(scalar float64) Color {
	return Color{
		r: c.r * scalar,
		b: c.b * scalar,
		g: c.g * scalar,
	}
}

func PiecewiseMult(c1 Color, c2 Color) Color {
	return Color{
		r: c1.r * c2.r,
		b: c1.b * c2.b,
		g: c1.g * c2.g,
	}
}

func White() Color {
	return Color{1, 1, 1}
}

func Red() Color {
	return Color{1, 0, 0}
}

func Black() Color {
	return Color{0, 0, 0}
}

func clampValue(val float64) float64 {
	if val < 0 {
		return 0
	} else if val > 1 {
		return 1
	} else {
		return val
	}
}

func (c Color) Clamp() Color {
	return Color{
		clampValue(c.r),
		clampValue(c.g),
		clampValue(c.b)}
}
