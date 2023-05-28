package light

import (
	"fmt"
	"math"
	"testing"

	"github.com/noahssarcastic/gort/pkg/color"
	"github.com/noahssarcastic/gort/pkg/material"
	"github.com/noahssarcastic/gort/pkg/tuple"
)

func TestLightingFunction(t *testing.T) {
	material := material.Default()
	pos := tuple.Point(0, 0, 0)
	tests := []struct {
		eye    tuple.Tuple
		normal tuple.Tuple
		light  *PointLight
		want   color.Color
	}{
		{
			tuple.Vector(0, 0, -1),
			tuple.Vector(0, 0, -1),
			NewPointLight(color.White, tuple.Point(0, 0, -10)),
			color.New(1.9, 1.9, 1.9),
		},
		{
			tuple.Vector(0, math.Sqrt(2)/2, -math.Sqrt(2)/2),
			tuple.Vector(0, 0, -1),
			NewPointLight(color.White, tuple.Point(0, 0, -10)),
			color.New(1.0, 1.0, 1.0),
		},
		{
			tuple.Vector(0, 0, -1),
			tuple.Vector(0, 0, -1),
			NewPointLight(color.White, tuple.Point(0, 10, -10)),
			color.New(0.7364, 0.7364, 0.7364),
		},
		{
			tuple.Vector(0, -math.Sqrt(2)/2, -math.Sqrt(2)/2),
			tuple.Vector(0, 0, -1),
			NewPointLight(color.White, tuple.Point(0, 10, -10)),
			color.New(1.6364, 1.6364, 1.6364),
		},
		{
			tuple.Vector(0, 0, -1),
			tuple.Vector(0, 0, -1),
			NewPointLight(color.White, tuple.Point(0, 0, 10)),
			color.New(0.1, 0.1, 0.1),
		},
	}
	for i, tt := range tests {
		name := fmt.Sprint(i)
		t.Run(name, func(t *testing.T) {
			got := Lighting(material, pos, tt.light, tt.eye, tt.normal)
			if !color.Equal(tt.want, got) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
