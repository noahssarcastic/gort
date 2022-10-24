package matrix

import (
	"fmt"
	"testing"
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
	want := 4.
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
