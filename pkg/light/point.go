package light

import (
	"github.com/noahssarcastic/gort/pkg/color"
	"github.com/noahssarcastic/gort/pkg/tuple"
)

type PointLight struct {
	intensity color.Color
	pos       tuple.Tuple
}

func NewPointLight(intensity color.Color, pos tuple.Tuple) *PointLight {
	return &PointLight{intensity, pos}
}

func (light *PointLight) Intensity() color.Color {
	return light.intensity
}

func (light *PointLight) Position() tuple.Tuple {
	return light.pos
}
