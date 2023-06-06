// Package image provides an interface for manipulating 2D pixel data.
package image

// TODO: Convert pixels to 2D array

import (
	"github.com/noahssarcastic/gort/pkg/color"
	"github.com/noahssarcastic/gort/pkg/ppm"
)

// Image stores a 2D array of color.Color values.
type Image struct {
	w, h   int
	pixels []color.Color
}

// New returns a pointer to a blank Image.
// A blank Image is initialized with all white pixels.
func New(w, h int) *Image {
	pixels := make([]color.Color, w*h)
	for i := range pixels {
		pixels[i] = color.White
	}
	return &Image{w, h, pixels}
}

// Width gets the width in pixels of a Image.
func (img *Image) Width() int {
	return img.w
}

// Height gets the height in pixels of a Image.
func (img *Image) Height() int {
	return img.h
}

// Get returns the color.Color value of a pixel at a given coordinate.
func (img *Image) Get(x, y int) color.Color {
	return img.pixels[x+y*img.w]
}

// Get manipulates the color.Color value of a pixel at a given coordinate.
func (img *Image) Set(x, y int, c color.Color) {
	img.pixels[x+y*img.w] = c
}

// ImageToPixelMap converts a Image to a ppm.PixelMap.
func ImageToPixelMap(img Image) *ppm.PixelMap {
	w := img.Width()
	h := img.Height()
	pm := ppm.New(w, h)
	var c color.Color
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			c = color.Clamp(img.Get(x, y))
			r, g, b := int(c.R()*ppm.MaxColor),
				int(c.G()*ppm.MaxColor),
				int(c.B()*ppm.MaxColor)
			pm.Set(x, y, r, g, b)
		}
	}
	return pm
}
