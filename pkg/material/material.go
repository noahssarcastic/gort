package material

import "github.com/noahssarcastic/gort/pkg/color"

type Material struct {
	Color                      color.Color
	Ambient, Diffuse, Specular float64
	Shininess                  float64
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
