package matrix

import (
	"fmt"
	"testing"

	"github.com/noahssarcastic/tddraytracer/tuple"
)

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
	want := 2.
	got := mat.Get(0, 1)
	if want != got {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestIsMatrix(t *testing.T) {
	var tests = []struct {
		mat  Matrix
		want bool
	}{
		{
			Matrix{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			true,
		},
		{
			Matrix{
				{1, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 1},
			},
			false,
		},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%v", tt.mat)
		t.Run(name, func(t *testing.T) {
			ans := tt.mat.IsMatrix()
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestEqual(t *testing.T) {
	var tests = []struct {
		a, b Matrix
		want bool
	}{
		{
			Matrix{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			Matrix{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			true,
		},
		{
			Matrix{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			Matrix{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			false,
		},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%v,%v", tt.a, tt.b)
		t.Run(name, func(t *testing.T) {
			ans := Equal(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestMultiply_square(t *testing.T) {
	a := Matrix{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	}
	b := Matrix{
		{-2, 1, 2, 3},
		{3, 2, 1, -1},
		{4, 3, 6, 5},
		{1, 2, 7, 8},
	}
	want := Matrix{
		{20, 22, 50, 48},
		{44, 54, 114, 108},
		{40, 58, 110, 102},
		{16, 26, 46, 42},
	}
	got := Multiply(a, b)
	if !Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestMultiply_identity(t *testing.T) {
	mat := Matrix{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	}
	id := I(4)
	want := Matrix{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	}
	got := Multiply(mat, id)
	if !Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestMultiply_col_vector(t *testing.T) {
	a := Matrix{
		{1, 2, 3, 4},
		{2, 4, 4, 2},
		{8, 6, 4, 1},
		{0, 0, 0, 1},
	}
	b := Matrix{
		{1},
		{2},
		{3},
		{1},
	}
	want := Matrix{
		{18},
		{24},
		{33},
		{1},
	}
	got := Multiply(a, b)
	if !Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestMultiply_tuple(t *testing.T) {
	mat := Matrix{
		{1, 2, 3, 4},
		{2, 4, 4, 2},
		{8, 6, 4, 1},
		{0, 0, 0, 1},
	}
	tup := tuple.New(1, 2, 3, 1)
	want := tuple.New(18, 24, 33, 1)
	got := mat.Multiply(tup)
	if !tuple.Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestMatrixT(t *testing.T) {
	mat := Matrix{
		{1, 2, 3},
		{4, 5, 6},
	}
	want := Matrix{
		{1, 4},
		{2, 5},
		{3, 6},
	}
	got := mat.T()
	if !Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestMatrixT_identity(t *testing.T) {
	mat := I(4)
	want := mat
	got := mat.T()
	if !Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}
