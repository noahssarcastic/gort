package ppm

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPixelMap_width(t *testing.T) {
	w, h := 10, 10
	pm := New(w, h)
	want, got := w, pm.Width()
	if want != got {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestPixelMap_height(t *testing.T) {
	w, h := 10, 10
	pm := New(w, h)
	want, got := h, pm.Height()
	if want != got {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestWritePPM(t *testing.T) {
	var buff bytes.Buffer
	w, h := 10, 10
	pm := New(w, h)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			pm.Set(x, y, 0, 0, 0)
		}
	}
	WritePPM(&buff, pm)
	str := buff.String()
	fmt.Println(str)
}
