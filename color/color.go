package color

import "github.com/noahssarcastic/tddraytracer/utils"

type Color struct {
	r, g, b float64
}

func New(r, g, b float64) Color {
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

func Equal(c1, c2 Color) bool {
	return utils.FloatEqual(c1.r, c2.r) &&
		utils.FloatEqual(c1.g, c2.g) &&
		utils.FloatEqual(c1.b, c2.b)
}

func (c1 Color) Add(c2 Color) Color {
	return Color{
		r: c1.r + c2.r,
		b: c1.b + c2.b,
		g: c1.g + c2.g,
	}
}

func (c1 Color) Subtract(c2 Color) Color {
	return Color{
		r: c1.r - c2.r,
		b: c1.b - c2.b,
		g: c1.g - c2.g,
	}
}

func (c Color) Multiply(scalar float64) Color {
	return Color{
		r: c.r * scalar,
		b: c.b * scalar,
		g: c.g * scalar,
	}
}

func PiecewiseMultiply(c1 Color, c2 Color) Color {
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
