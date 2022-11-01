package main

import (
	"math"

	"github.com/noahssarcastic/tddraytracer/image/canvas"
	"github.com/noahssarcastic/tddraytracer/image/color"
	"github.com/noahssarcastic/tddraytracer/image/ppm"
	"github.com/noahssarcastic/tddraytracer/math/transform"
	"github.com/noahssarcastic/tddraytracer/math/tuple"
)

func main() {
	canv := canvas.New(100, 100)
	for i := 1; i <= 12; i++ {
		pt := tuple.Point(1, 0, 0)
		tform := transform.Chain(
			transform.Scale(30, 0, 0),
			transform.RotateZ(2.*math.Pi/12*float64(i)),
			transform.Translate(float64(canv.Width())/2, float64(canv.Width())/2, 0),
		)
		pt = tform.Apply(pt)
		x := int(pt.X())
		y := int(pt.Y())
		canv.SetPixel(x, (canv.Height()-1)-y, color.Red())
	}
	ppm.CanvasToPixmap(canv).WritePPM()
}
