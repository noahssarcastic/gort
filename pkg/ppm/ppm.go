// Package ppm implements the Plain PPM (Portable Pixel Map) image format.
// See https://netpbm.sourceforge.net/doc/ppm.html.
package ppm

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	MaxColor = 255 // maximum valid RGB color value.
)

const (
	magicNumber = "P3" // magic number for Plain PPM.
	maxLineLen  = 70   // maximum valid line length for PPM
)

type pixel struct {
	r, g, b int
}

// A PixelMap represents a single image.
type PixelMap struct {
	pixels [][]pixel
}

// NewPixelMap initializes an all-black w x h image.
func New(w, h int) *PixelMap {
	pixels := make([][]pixel, 0, h)
	for i := 0; i < h; i++ {
		pixels = append(pixels, make([]pixel, w))
	}
	return &PixelMap{pixels}
}

// Width returns the image's width in pixels.
func (pm *PixelMap) Width() int {
	return len(pm.pixels[0])
}

// Height returns the image's height in pixels.
func (pm *PixelMap) Height() int {
	return len(pm.pixels)
}

// Set changes the value of the pixel at (x,y) to the given RGB value.
func (pm *PixelMap) Set(x, y int, r, g, b int) {
	pm.pixels[y][x] = pixel{r, g, b}
}

func (pm *PixelMap) header() string {
	return fmt.Sprintf(
		"%v\n%v %v\n%v\n",
		magicNumber,
		pm.Width(), pm.Height(),
		MaxColor)
}

// WritePPM writes the given PixelMap to the given [io.Writer].
func WritePPM(w io.Writer, pm *PixelMap) error {
	_, err := w.Write([]byte(pm.header()))
	if err != nil {
		return fmt.Errorf("could not write header: %w", err)
	}
	for j, row := range pm.pixels {
		err := writeRow(w, row)
		if err != nil {
			return fmt.Errorf("could not write row %d: %w", j, err)
		}
	}
	return nil
}

func writeRow(w io.Writer, row []pixel) (err error) {
	buff := make([]byte, 0, maxLineLen)
	for _, pixel := range row {
		buff, err = writePixel(w, buff, pixel)
		if err != nil {
			return err
		}
	}
	err = flush(w, buff)
	if err != nil {
		return err
	}
	return nil
}

func writePixel(w io.Writer, buff []byte, p pixel) (_ []byte, err error) {
	for _, val := range []int{p.r, p.g, p.b} {
		str := strconv.Itoa(val)
		if len(buff)+len(str) > maxLineLen {
			err = flush(w, buff)
			if err != nil {
				return buff, err
			}
			buff = make([]byte, 0, maxLineLen)
		}
		buff = append(buff, []byte(str)...)
		buff = append(buff, ' ') // add padding
	}
	return buff, nil
}

func flush(w io.Writer, buff []byte) error {
	str := string(buff)
	str = strings.TrimSpace(str)
	str += "\n"
	_, err := w.Write([]byte(str))
	if err != nil {
		return err
	}
	return nil
}
