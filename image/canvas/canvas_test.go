package canvas

import (
	"testing"

	"github.com/noahssarcastic/tddraytracer/image/color"
)

func TestNew(t *testing.T) {
	w := 10
	h := 15
	canv := New(w, h)

	var want, got int
	want = w
	if got = canv.Width(); want != got {
		t.Errorf("want %v; got %v", want, got)
	}

	want = h
	if got = canv.Height(); want != got {
		t.Errorf("want %v; got %v", want, got)
	}

	for x := 0; x < w; x++ {
		for y := 0; y < w; y++ {
			if !color.Equal(canv.GetPixel(x, y), color.White()) {
				t.Errorf("pixel at %v,%v is not white", x, y)
			}
		}
	}
}

func TestSetPixel(t *testing.T) {
	canv := New(10, 10)
	want := color.Red()
	canv.SetPixel(0, 0, color.Red())
	if got := canv.GetPixel(0, 0); !color.Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}
