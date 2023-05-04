package main

import (
	"flag"
	"os"

	"github.com/noahssarcastic/gort/pkg/color"
	"github.com/noahssarcastic/gort/pkg/geometry"
	"github.com/noahssarcastic/gort/pkg/image"
	"github.com/noahssarcastic/gort/pkg/math"
	"github.com/noahssarcastic/gort/pkg/ppm"
	"github.com/noahssarcastic/gort/pkg/ray"
)

var out = flag.String("o", "trace.ppm", "output image path")

func main() {
	flag.Parse()

	w, h := 500, 500
	img := image.New(w, h)

	eye := math.Point(0, 0, -10)
	screenOrigin := math.Point(0, 0, 0)

	// assume eye and screen are one the same ground plane
	cellCenterPadding := 0.5
	x0 := screenOrigin.X() - float64(w)/2 + cellCenterPadding
	y0 := screenOrigin.Y() - float64(h)/2 + cellCenterPadding

	objects := []ray.Object{
		geometry.NewSphere(math.Chain(
			math.Scale(100, 50, 100),
		)),
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			screenCellCenter := math.Point(
				x0+float64(x),
				y0+float64(y),
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
				img.Set(x, y, color.Red())
			} else {
				img.Set(x, y, color.Black())
			}
		}
	}

	pm := image.ImageToPixelMap(*img)
	f, err := os.OpenFile(*out,
		os.O_RDWR|os.O_CREATE|os.O_TRUNC,
		0755)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}()
	ppm.WritePPM(f, pm)
}
