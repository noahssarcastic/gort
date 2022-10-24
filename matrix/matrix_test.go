package matrix

import "testing"

func TestWidth(t *testing.T) {
	mat := Matrix{
		{1, 2, 3},
		{4, 5, 6},
	}
	want := 3
	got := mat.Width()
	if want != got {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestHeight(t *testing.T) {
	mat := Matrix{
		{1, 2, 3},
		{4, 5, 6},
	}
	want := 2
	got := mat.Height()
	if want != got {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestAccess(t *testing.T) {
	mat := Matrix{
		{1, 2, 3},
		{4, 5, 6},
	}
	want := 4.
	got := mat[1][0]
	if want != got {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestGet(t *testing.T) {
	mat := Matrix{
		{1, 2, 3},
		{4, 5, 6},
	}
	want := 4.
	got := mat.Get(1, 0)
	if want != got {
		t.Errorf("want %v; got %v", want, got)
	}
}
