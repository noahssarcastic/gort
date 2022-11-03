package ray

import (
	"fmt"
	"testing"

	"github.com/noahssarcastic/gort/math"
)

func rayEqual(r1, r2 Ray) bool {
	return math.TupleEqual(r1.origin, r2.origin) &&
		math.TupleEqual(r1.direction, r2.direction)
}

func TestRayPosition(t *testing.T) {
	r := NewRay(math.Point(2, 3, 4), math.Vector(1, 0, 0))
	tests := []struct {
		ray  Ray
		t    float64
		want math.Tuple
	}{
		{r, 0, math.Point(2, 3, 4)},
		{r, 1, math.Point(3, 3, 4)},
		{r, -1, math.Point(1, 3, 4)},
		{r, 2.5, math.Point(4.5, 3, 4)},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%v,%v", tt.ray, tt.t)
		t.Run(name, func(t *testing.T) {
			ans := Position(tt.ray, tt.t)
			if !math.TupleEqual(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestRayTransform(t *testing.T) {
	tests := []struct {
		r    Ray
		mat  math.Matrix
		want Ray
	}{
		{
			NewRay(math.Point(1, 2, 3), math.Vector(0, 1, 0)),
			math.Translate(3, 4, 5),
			NewRay(math.Point(4, 6, 8), math.Vector(0, 1, 0)),
		},
		{
			NewRay(math.Point(1, 2, 3), math.Vector(0, 1, 0)),
			math.Scale(2, 3, 4),
			NewRay(math.Point(2, 6, 12), math.Vector(0, 3, 0)),
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			ans := tt.r.Transform(tt.mat)
			if !rayEqual(tt.want, ans) {
				t.Errorf("want %v; got %v", tt.want, ans)
			}
		})
	}
}

func TestRayTransform_no_modify(t *testing.T) {
	r := NewRay(math.Point(1, 2, 3), math.Vector(0, 1, 0))
	mat := math.Translate(3, 4, 5)
	want := NewRay(math.Point(1, 2, 3), math.Vector(0, 1, 0))
	r.Transform(mat)
	got := r
	if !rayEqual(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}
