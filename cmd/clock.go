package cmd

import (
	stdmath "math"

	"github.com/noahssarcastic/gort/image"
	"github.com/noahssarcastic/gort/math"
)

func RunClock() {
	canv := image.NewCanvas(100, 100)
	for i := 1; i <= 12; i++ {
		pt := math.Point(1, 0, 0)
		tform := math.Chain(
			math.Scale(30, 0, 0),
			math.RotateZ(2.*stdmath.Pi/12*float64(i)),
			math.Translate(float64(canv.Width())/2, float64(canv.Width())/2, 0),
		)
		pt = tform.Apply(pt)
		x := int(pt.X())
		y := int(pt.Y())
		canv.SetPixel(x, (canv.Height()-1)-y, image.Red())
	}
	image.CanvasToPixmap(canv).WritePPM()
}
