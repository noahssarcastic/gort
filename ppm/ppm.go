package ppm

import (
	"fmt"
	"os"

	"github.com/noahssarcastic/tddraytracer/canvas"
	"github.com/noahssarcastic/tddraytracer/color"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const MAX_COLOR = 255
const MAGIC_NUMBER = "P3"

type Header struct {
	w, h int
}

type Pixel struct {
	r, g, b int
}

type Row []Pixel

type PixelData []Row

type Pixmap struct {
	head   Header
	pixels PixelData
}

func colorValueToPixelValue(val float64) int {
	return int(val * MAX_COLOR)
}

func colorToPixel(c color.Color) Pixel {
	clampedColor := c.Clamp()
	return Pixel{
		colorValueToPixelValue(clampedColor.R()),
		colorValueToPixelValue(clampedColor.G()),
		colorValueToPixelValue(clampedColor.B())}
}

func CanvasToPixmap(canv canvas.Canvas) Pixmap {
	w := canv.Width()
	h := canv.Height()

	head := Header{w, h}

	pixels := make(PixelData, h)
	for y := 0; y < h; y++ {
		row := make(Row, w)
		for x := 0; x < w; x++ {
			row[x] = colorToPixel(canv.GetPixel(x, y))
		}
		pixels[y] = row
	}

	return Pixmap{head, pixels}
}

func (pm Pixmap) getHeader() string {
	return fmt.Sprintf(
		"%v\n%v %v\n%v\n",
		MAGIC_NUMBER,
		pm.head.w, pm.head.h,
		MAX_COLOR)
}

func (pm Pixmap) WritePPM() {
	f, err := os.Create("pixmap.ppm")
	check(err)
	defer f.Close()

	_, err = f.WriteString(pm.getHeader())
	check(err)

	for _, row := range pm.pixels {
		for i, pixel := range row {
			if i != 0 {
				_, err = f.WriteString(" ")
				check(err)
			}
			_, err = f.WriteString(fmt.Sprintf("%v %v %v", pixel.r, pixel.g, pixel.b))
			check(err)
		}
		_, err = f.WriteString("\n")
		check(err)
	}
	// _, err = f.WriteString("\n")
	// check(err)
}
