package main

import (
	"errors"
	"flag"
	"runtime/pprof"

	"github.com/noahssarcastic/gort/pkg/color"
	"github.com/noahssarcastic/gort/pkg/geo"
	"github.com/noahssarcastic/gort/pkg/image"
	"github.com/noahssarcastic/gort/pkg/matrix"
	"github.com/noahssarcastic/gort/pkg/ppm"
	"github.com/noahssarcastic/gort/pkg/ray"
	"github.com/noahssarcastic/gort/pkg/tuple"
)

type Object = ray.Intersectable

func main() {
	flag.Parse()
	initConfig()
	defer cleanUp()
	if cfg.profile != nil {
		pprof.StartCPUProfile(cfg.profile)
		defer pprof.StopCPUProfile()
	}

	w, h := 500, 500
	img := image.New(w, h)

	eye := tuple.Point(0, 0, -10)
	screenOrigin := tuple.Point(0, 0, 0)

	// assume eye and screen are one the same ground plane
	cellCenterPadding := 0.5
	x0 := screenOrigin.X() - float64(w)/2 + cellCenterPadding
	y0 := screenOrigin.Y() - float64(h)/2 + cellCenterPadding

	objects := []Object{
		geo.NewSphere(matrix.Chain(
			matrix.Scale(100, 100, 100),
		)),
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			screenCellCenter := tuple.Point(
				x0+float64(x),
				y0+float64(y),
				screenOrigin.Z(),
			)
			rayDir := tuple.Sub(screenCellCenter, eye)
			rayDir = tuple.Norm(rayDir)
			r := ray.New(screenCellCenter, rayDir)
			xs := make([]ray.Intersect, 0)
			for _, obj := range objects {
				for _, i := range obj.Intersect(r) {
					xs = ray.InsertIntersect(xs, i)
				}
			}
			_, err := ray.Hit(xs)
			if errors.Is(err, ray.ErrNoHits) {
				img.Set(x, y, color.Black)
			} else {
				img.Set(x, y, color.Red)
			}
		}
	}

	pm := image.ImageToPixelMap(*img)
	ppm.WritePPM(cfg.file, pm)
}
