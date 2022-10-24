package main

import (
	"github.com/noahssarcastic/tddraytracer/canvas"
	"github.com/noahssarcastic/tddraytracer/color"
	"github.com/noahssarcastic/tddraytracer/ppm"
)

func main() {
	canv := canvas.New(10, 2)
	for y := 0; y < canv.Height(); y++ {
		for x := 0; x < canv.Width(); x++ {
			canv.SetPixel(x, y, color.Red())
		}
	}
	ppm.CanvasToPixmap(canv).WritePPM()
}
