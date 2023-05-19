package main

import (
	"errors"
	"flag"
	"runtime/pprof"

	"github.com/noahssarcastic/gort/pkg/color"
	"github.com/noahssarcastic/gort/pkg/geo"
	"github.com/noahssarcastic/gort/pkg/image"
	"github.com/noahssarcastic/gort/pkg/light"
	"github.com/noahssarcastic/gort/pkg/material"
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

	w, h := 100, 100
	img := image.New(w, h)

	eye := tuple.Point(0, 0, -5)
	screenOrigin := tuple.Point(0, 0, 10)

	// assume eye and screen are one the same ground plane
	cellCenterPadding := 0.5
	x0 := screenOrigin.X() - float64(w)/2 + cellCenterPadding
	y0 := screenOrigin.Y() - float64(h)/2 + cellCenterPadding

	sphere := geo.NewSphere()
	sphere.SetMaterial(material.New(color.New(1, .2, 1), 0.1, 0.9, 0.9, 200))
	objects := []Object{
		sphere,
	}

	ptLight := light.NewPointLight(color.White, tuple.Point(-10, 10, -10))

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			cellSize := 7.0 / float64(w)
			screenCellCenter := tuple.Point(
				(x0+float64(x))*cellSize,
				(y0+float64(y))*cellSize,
				screenOrigin.Z(),
			)
			rayDir := tuple.Sub(screenCellCenter, eye)
			rayDir = tuple.Norm(rayDir)
			r := ray.New(eye, rayDir)
			xs := make([]ray.Intersect, 0)
			for _, obj := range objects {
				for _, x := range obj.Intersect(r) {
					xs = ray.InsertIntersect(xs, x)
				}
			}
			hit, err := ray.Hit(xs)
			if errors.Is(err, ray.ErrNoHits) {
				img.Set(x, y, color.Black)
			} else {
				pt := ray.Position(r, hit.Distance())
				normalVec := hit.Object().NormalAt(pt)
				eyeVec := tuple.Neg(r.Direction())
				pixel := light.Lighting(
					hit.Object().Material(),
					pt,
					ptLight,
					eyeVec,
					normalVec,
				)
				img.Set(x, h-1-y, pixel)
			}
		}
	}

	pm := image.ImageToPixelMap(*img)
	ppm.WritePPM(cfg.file, pm)
}
