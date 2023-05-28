package material

import "github.com/noahssarcastic/gort/pkg/color"

type Material struct {
	color                      color.Color
	ambient, diffuse, specular float64
	shininess                  float64
}

func New(
	color color.Color,
	ambient, diffuse, specular float64,
	shininess float64,
) Material {
	return Material{color, ambient, diffuse, specular, shininess}
}

func Default() Material {
	return Material{color.White, 0.1, 0.9, 0.9, 200}
}

func (m Material) Color() color.Color {
	return m.color
}

func (m Material) Ambient() float64 {
	return m.ambient
}

func (m Material) Diffuse() float64 {
	return m.diffuse
}

func (m Material) Specular() float64 {
	return m.specular
}

func (m Material) Shininess() float64 {
	return m.shininess
}
