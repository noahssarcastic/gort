package canvas

import "github.com/noahssarcastic/tddraytracer/color"

type Canvas struct {
	w, h   int
	pixels []color.Color
}

func New(w, h int) Canvas {
	pixels := make([]color.Color, w*h)
	for i := range pixels {
		pixels[i] = color.White()
	}
	return Canvas{w, h, pixels}
}

func (canv *Canvas) Width() int {
	return canv.w
}

func (canv *Canvas) Height() int {
	return canv.h
}

func (canv *Canvas) GetPixel(x, y int) color.Color {
	return canv.pixels[x+y*canv.w]
}
