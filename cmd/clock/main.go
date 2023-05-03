package main

import (
	stdmath "math"
	"os"

	"github.com/noahssarcastic/gort/color"
	"github.com/noahssarcastic/gort/image"
	"github.com/noahssarcastic/gort/math"
	"github.com/noahssarcastic/gort/ppm"
)

func main() {
	img := image.New(100, 100)
	for i := 1; i <= 12; i++ {
		pt := math.Point(1, 0, 0)
		tform := math.Chain(
			math.Scale(30, 0, 0),
			math.RotateZ(2.*stdmath.Pi/12*float64(i)),
			math.Translate(float64(img.Width())/2, float64(img.Width())/2, 0),
		)
		pt = tform.Apply(pt)
		x := int(pt.X())
		y := int(pt.Y())
		img.Set(x, (img.Height()-1)-y, color.Red())
	}
	pm := image.ImageToPixelMap(*img)

	f, err := os.OpenFile("clock.ppm", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
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
