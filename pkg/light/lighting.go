package light

import (
	"math"

	"github.com/noahssarcastic/gort/pkg/color"
	"github.com/noahssarcastic/gort/pkg/material"
	"github.com/noahssarcastic/gort/pkg/tuple"
)

func Lighting(
	mat material.Material,
	pt tuple.Tuple,
	light *PointLight,
	eyeVec, normalVec tuple.Tuple,
) color.Color {
	effectiveColor := color.PiecewiseMult(mat.Color, light.Intensity())
	lightVec := tuple.Norm(light.Position().Sub(pt))
	// initialize as black
	var diffuse, specular, ambient color.Color
	if cosLightNormal := tuple.Dot(lightVec, normalVec); cosLightNormal >= 0 {
		diffuse = color.Mult(effectiveColor, mat.Diffuse*cosLightNormal)
		reflectVec := tuple.Reflect(tuple.Neg(lightVec), normalVec)
		if cosReflectEye := tuple.Dot(reflectVec, eyeVec); cosReflectEye > 0 {
			factor := math.Pow(cosReflectEye, mat.Shininess)
			specular = color.Mult(light.Intensity(), mat.Specular*factor)
		}
	}
	ambient = color.Mult(effectiveColor, mat.Ambient)
	return diffuse.Add(specular).Add(ambient)
}
