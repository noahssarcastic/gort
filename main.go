package main

import (
	"fmt"

	"github.com/noahssarcastic/gort/geometry"
	"github.com/noahssarcastic/gort/math"
	"github.com/noahssarcastic/gort/ray"
)

func main() {
	width := 5
	height := 5
	// canv := image.NewCanvas(width, height)

	eye := math.Point(0, 0, -5)
	screenOrigin := math.Point(0, 0, 0)

	// assume eye and screen are one the same ground plane
	cellCenterPadding := 0.5
	startingX := screenOrigin.X() - float64(width)/2 + cellCenterPadding
	startingY := screenOrigin.Y() - float64(height)/2 + cellCenterPadding

	sphere := geometry.NewSphere()
	sphere.SetTransform(
		math.Chain(
			math.Scale(1, 0.5, 1),
		),
	)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			screenCellCenter := math.Point(
				startingX+float64(x),
				startingY+float64(y),
				screenOrigin.Z(),
			)
			rayDir := screenCellCenter.Sub(eye)
			rayDir = math.Norm(rayDir)
			r := ray.NewRay(screenCellCenter, rayDir)
			fmt.Println(r)
			// 	ray.Combine(
			// 		sphere.Intersect(r),
			// 	)
		}
	}
}
