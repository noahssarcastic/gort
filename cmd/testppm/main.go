package main

import (
	"github.com/noahssarcastic/tddraytracer/canvas"
	"github.com/noahssarcastic/tddraytracer/color"
	"github.com/noahssarcastic/tddraytracer/ppm"
)

func main() {
	canv := canvas.New(100, 100)
	for y := 0; y < canv.Height(); y++ {
		for x := 0; x < canv.Width(); x++ {
			canv.SetPixel(
				x, y,
				color.New(
					float64(x)/float64(canv.Width()),
					0,
					float64(y)/float64(canv.Height())))
		}
	}
	ppm.CanvasToPixmap(canv).WritePPM()
}
