package image

type Canvas struct {
	w, h   int
	pixels []Color
}

func NewCanvas(w, h int) Canvas {
	pixels := make([]Color, w*h)
	for i := range pixels {
		pixels[i] = White()
	}
	return Canvas{w, h, pixels}
}

func (canv *Canvas) Width() int {
	return canv.w
}

func (canv *Canvas) Height() int {
	return canv.h
}

func (canv *Canvas) GetPixel(x, y int) Color {
	return canv.pixels[x+y*canv.w]
}

func (canv *Canvas) SetPixel(x, y int, c Color) {
	canv.pixels[x+y*canv.w] = c
}
