package main

import (
	"github.com/noahssarcastic/gort/geometry"
	"github.com/noahssarcastic/gort/image"
	"github.com/noahssarcastic/gort/math"
	"github.com/noahssarcastic/gort/ray"
)

func main() {
	width := 500
	height := 500
	canv := image.NewCanvas(width, height)

	eye := math.Point(0, 0, -10)
	screenOrigin := math.Point(0, 0, 0)

	// assume eye and screen are one the same ground plane
	cellCenterPadding := 0.5
	startingX := screenOrigin.X() - float64(width)/2 + cellCenterPadding
	startingY := screenOrigin.Y() - float64(height)/2 + cellCenterPadding

	objects := []ray.Object{
		geometry.NewSphere(math.Chain(
			math.Scale(100, 50, 100),
		)),
	}

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
			xs := make([]ray.Intersection, 0)
			for _, obj := range objects {
				for _, i := range obj.Intersect(r) {
					xs = ray.InsertIntersection(xs, &i)
				}
			}
			hit := ray.Hit(xs)
			if hit != nil {
				canv.SetPixel(x, y, image.Red())
			} else {
				canv.SetPixel(x, y, image.Black())
			}
		}
	}

	image.CanvasToPixmap(canv).WritePPM()

}
