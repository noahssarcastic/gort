package image

import (
	"testing"
)

func TestNew(t *testing.T) {
	w := 10
	h := 15
	canv := NewCanvas(w, h)

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
			if !ColorEqual(canv.GetPixel(x, y), White()) {
				t.Errorf("pixel at %v,%v is not white", x, y)
			}
		}
	}
}

func TestSetPixel(t *testing.T) {
	canv := NewCanvas(10, 10)
	want := Red()
	canv.SetPixel(0, 0, Red())
	if got := canv.GetPixel(0, 0); !ColorEqual(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}
