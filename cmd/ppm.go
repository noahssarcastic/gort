package cmd

import (
	"github.com/noahssarcastic/gort/image"
)

func RunPixmap() {
	canv := image.NewCanvas(100, 100)
	for y := 0; y < canv.Height(); y++ {
		for x := 0; x < canv.Width(); x++ {
			canv.SetPixel(
				x, y,
				image.NewColor(
					float64(x)/float64(canv.Width()),
					0,
					float64(y)/float64(canv.Height())))
		}
	}
	image.CanvasToPixmap(canv).WritePPM()
}
