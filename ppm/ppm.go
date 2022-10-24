package ppm

import (
	"fmt"
	"os"
	"strconv"

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

func insertString(line []byte, s string) []byte {
	if len(line) == 0 {
		return append(line, []byte(s)...)
	}

	// check if enough room to add string
	if len(line)+len(s) >= cap(line) {
		// scale up array and copy
		newSlice := make([]byte, cap(line)+70)[0:0]
		line = append(newSlice, line...)

		// start new line
		line = append(line, '\n')
	} else {
		// add padding
		line = append(line, ' ')
	}
	return append(line, []byte(s)...)
}

func (pm Pixmap) WritePPM() {
	f, err := os.Create("pixmap.ppm")
	check(err)
	defer f.Close()

	_, err = f.WriteString(pm.getHeader())
	check(err)

	for _, row := range pm.pixels {
		line := make([]byte, 70)[0:0]
		for _, pixel := range row {
			line = insertString(line, strconv.Itoa(pixel.r))
			line = insertString(line, strconv.Itoa(pixel.g))
			line = insertString(line, strconv.Itoa(pixel.b))
		}
		_, err = f.Write(append(line, '\n'))
		check(err)
	}
}
