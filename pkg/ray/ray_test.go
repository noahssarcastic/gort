package ray

import (
	"fmt"
	"testing"

	"github.com/noahssarcastic/gort/pkg/mat"
	"github.com/noahssarcastic/gort/pkg/tuple"
)

func rayEqual(r1, r2 Ray) bool {
	return tuple.Equal(r1.origin, r2.origin) &&
		tuple.Equal(r1.direction, r2.direction)
}

func TestPosition(t *testing.T) {
	r := New(tuple.Point(2, 3, 4), tuple.Vector(1, 0, 0))
	tests := []struct {
		ray  Ray
		t    float64
		want tuple.Tuple
	}{
		{r, 0, tuple.Point(2, 3, 4)},
		{r, 1, tuple.Point(3, 3, 4)},
		{r, -1, tuple.Point(1, 3, 4)},
		{r, 2.5, tuple.Point(4.5, 3, 4)},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%v,%v", tt.ray, tt.t)
		t.Run(name, func(t *testing.T) {
			ans := Position(tt.ray, tt.t)
			if !tuple.Equal(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestTransform(t *testing.T) {
	tests := []struct {
		r    Ray
		mat  mat.Matrix
		want Ray
	}{
		{
			New(tuple.Point(1, 2, 3), tuple.Vector(0, 1, 0)),
			mat.Translate(3, 4, 5),
			New(tuple.Point(4, 6, 8), tuple.Vector(0, 1, 0)),
		},
		{
			New(tuple.Point(1, 2, 3), tuple.Vector(0, 1, 0)),
			mat.Scale(2, 3, 4),
			New(tuple.Point(2, 6, 12), tuple.Vector(0, 3, 0)),
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			ans := Transform(tt.r, tt.mat)
			if !rayEqual(tt.want, ans) {
				t.Errorf("want %v; got %v", tt.want, ans)
			}
		})
	}
}

func TestTransform_no_modify(t *testing.T) {
	r := New(tuple.Point(1, 2, 3), tuple.Vector(0, 1, 0))
	mat := mat.Translate(3, 4, 5)
	want := New(tuple.Point(1, 2, 3), tuple.Vector(0, 1, 0))
	Transform(r, mat)
	got := r
	if !rayEqual(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}
