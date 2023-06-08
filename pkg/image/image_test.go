package image

import (
	"testing"

	"github.com/noahssarcastic/gort/pkg/color"
)

func TestWidth(t *testing.T) {
	w, h := 10, 15
	img := New(w, h)
	if want, got := w, img.Width(); want != got {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestHeight(t *testing.T) {
	w, h := 10, 15
	img := New(w, h)
	if want, got := h, img.Height(); want != got {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestNew(t *testing.T) {
	w, h := 10, 15
	img := New(w, h)
	for x := 0; x < w; x++ {
		for y := 0; y < w; y++ {
			if !color.Equal(img.Get(x, y), color.Black) {
				t.Errorf("pixel at %v,%v is not black", x, y)
			}
		}
	}
}

func TestSet(t *testing.T) {
	img := New(10, 10)
	img.Set(0, 0, color.Red)
	if want, got := color.Red, img.Get(0, 0); !color.Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}
