package color

import "github.com/noahssarcastic/tddraytracer/utils"

type Color struct {
	r, g, b float64
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
	return Color{0, 0, 0}
}
