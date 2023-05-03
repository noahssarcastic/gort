package color

import (
	"fmt"
	"testing"
)

func TestColorEquality(t *testing.T) {
	var tests = []struct {
		c1, c2 Color
		want   bool
	}{
		{Color{.1, .2, .3}, Color{.1, .2, .3}, true},
		{Color{0, .2, .3}, Color{.1, .2, .3}, false},
		{Color{.1, 0, .3}, Color{.1, .2, .3}, false},
		{Color{.1, .2, 0}, Color{.1, .2, .3}, false},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%v,%v", tt.c1, tt.c2)
		t.Run(name, func(t *testing.T) {
			ans := Equal(tt.c1, tt.c2)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestAddColors(t *testing.T) {
	c1 := Color{0.1, 0.2, 0.3}
	c2 := Color{0.4, 0.5, 0.6}
	got := Add(c1, c2)
	want := Color{0.5, 0.7, 0.9}
	if !Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestSubtractColors(t *testing.T) {
	c1 := Color{0.4, 0.5, 0.6}
	c2 := Color{0.1, 0.2, 0.3}
	got := Sub(c1, c2)
	want := Color{0.3, 0.3, 0.3}
	if !Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestMultiplyColorByScalar(t *testing.T) {
	c := Color{0.1, 0.2, 0.3}
	scalar := 2.0
	got := Mult(c, scalar)
	want := Color{0.2, 0.4, 0.6}
	if !Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestPiecewiseMultiplyColors(t *testing.T) {
	c1 := Color{0.4, 0.5, 0.6}
	c2 := Color{0.1, 0.2, 0.3}
	got := PiecewiseMult(c1, c2)
	want := Color{0.04, 0.1, 0.18}
	if !Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestClampColor(t *testing.T) {
	var tests = []struct {
		c    Color
		want Color
	}{
		{Color{0.5, 1.3, 0}, Color{0.5, 1, 0}},
		{Color{0.5, 0.9, -5}, Color{0.5, 0.9, 0}},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%v", tt.c)
		t.Run(name, func(t *testing.T) {
			ans := Clamp(tt.c)
			if !Equal(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
