package main

import (
	"flag"
	"math"
	"os"

	"github.com/noahssarcastic/gort/pkg/color"
	"github.com/noahssarcastic/gort/pkg/image"
	"github.com/noahssarcastic/gort/pkg/matrix"
	"github.com/noahssarcastic/gort/pkg/ppm"
	"github.com/noahssarcastic/gort/pkg/tuple"
)

var out = flag.String("o", "test.ppm", "output image path")

func main() {
	flag.Parse()

	img := image.New(100, 100)
	for i := 1; i <= 12; i++ {
		pt := tuple.Point(1, 0, 0)
		tform := matrix.Chain(
			matrix.Scale(30, 0, 0),
			matrix.RotateZ(2.*math.Pi/12*float64(i)),
			matrix.Translate(
				float64(img.Width())/2,
				float64(img.Width())/2,
				0,
			),
		)
		pt = tform.Apply(pt)
		x := int(pt.X)
		y := int(pt.Y)
		img.Set(x, (img.Height()-1)-y, color.Red)
	}
	pm := image.ImageToPixelMap(*img)

	f, err := os.OpenFile(*out, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
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
