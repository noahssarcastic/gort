package geo

import (
	"errors"
	"fmt"
	"math"
	"testing"

	"github.com/noahssarcastic/gort/pkg/color"
	"github.com/noahssarcastic/gort/pkg/material"
	"github.com/noahssarcastic/gort/pkg/matrix"
	"github.com/noahssarcastic/gort/pkg/ray"
	"github.com/noahssarcastic/gort/pkg/tuple"
	"github.com/noahssarcastic/gort/pkg/util"
)

func intersectsEqual(a, b []Intersection) bool {
	if len(a) != len(b) {
		return false
	}
	for i, el := range a {
		if !util.FloatEqual(el.Distance, b[i].Distance) {
			return false
		}
		if el.Sphere != b[i].Sphere {
			return false
		}
	}
	return true
}

func TestIntersect(t *testing.T) {
	sphere := NewSphere(matrix.I, material.New(color.White, 0.1, 0.9, 0.9, 200))
	tests := []struct {
		name  string
		start tuple.Tuple
		want  []Intersection
	}{
		{"normal", tuple.Point(0, 0, -5), []Intersection{
			{&sphere, 4},
			{&sphere, 6},
		}},
		{"tangent", tuple.Point(0, 1, -5), []Intersection{
			{&sphere, 5},
			{&sphere, 5},
		}},
		{"miss", tuple.Point(0, 2, -5), []Intersection{}},
		{"inside", tuple.Point(0, 0, 0), []Intersection{
			{&sphere, -1},
			{&sphere, 1},
		}},
		{"behind", tuple.Point(0, 0, 5), []Intersection{
			{&sphere, -6},
			{&sphere, -4},
		}},
	}
	for _, tt := range tests {
		name := tt.name
		t.Run(name, func(t *testing.T) {
			ans := Intersect(&sphere, ray.New(tt.start, tuple.Vector(0, 0, 1)))
			if !intersectsEqual(tt.want, ans) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestIntersect_transformed(t *testing.T) {
	sphere := NewSphere(
		matrix.Scale(2, 2, 2),
		material.New(color.White, 0.1, 0.9, 0.9, 200),
	)
	want := []Intersection{
		{&sphere, 3},
		{&sphere, 7},
	}
	ans := Intersect(&sphere, ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1)))
	if !intersectsEqual(want, ans) {
		t.Errorf("got %v, want %v", ans, want)
	}
}

func TestIntersect_miss(t *testing.T) {
	sphere := NewSphere(
		matrix.Translate(5, 0, 0),
		material.New(color.White, 0.1, 0.9, 0.9, 200),
	)
	ans := Intersect(&sphere, ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1)))
	if len(ans) > 0 {
		t.Errorf("got %v, want []Intersection{}", ans)
	}
}

func TestNormalAt(t *testing.T) {
	sphere := NewSphere(matrix.I, material.New(color.White, 0.1, 0.9, 0.9, 200))
	tests := []struct {
		pt   tuple.Tuple
		want tuple.Tuple
	}{
		{tuple.Point(1, 0, 0), tuple.Vector(1, 0, 0)},
		{tuple.Point(0, 1, 0), tuple.Vector(0, 1, 0)},
		{tuple.Point(0, 0, 1), tuple.Vector(0, 0, 1)},
		{
			tuple.Point(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3),
			tuple.Vector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3),
		},
	}
	for i, tt := range tests {
		name := fmt.Sprint(i)
		t.Run(name, func(t *testing.T) {
			ans := NormalAt(&sphere, tt.pt)
			if !tuple.Equal(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestNormalAt_normalized(t *testing.T) {
	sphere := NewSphere(matrix.I, material.New(color.White, 0.1, 0.9, 0.9, 200))
	pt := tuple.Point(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3)
	normal := NormalAt(&sphere, pt)
	want := tuple.Norm(normal)
	got := normal
	if !tuple.Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestNormalAt_transformed(t *testing.T) {
	tests := []struct {
		pt     tuple.Tuple
		sphere Sphere
		want   tuple.Tuple
	}{
		{
			tuple.Point(0, 1.70711, -0.70711),
			NewSphere(
				matrix.Translate(0, 1, 0),
				material.Default(),
			),
			tuple.Vector(0, 0.70711, -0.70711),
		},
		{
			tuple.Point(0, math.Sqrt(2)/2, -math.Sqrt(2)/2),
			NewSphere(
				matrix.Chain(
					matrix.RotateZ(math.Pi/5),
					matrix.Scale(1, 0.5, 1),
				),
				material.Default(),
			),
			tuple.Vector(0, 0.97014, -0.24254),
		},
	}
	for i, tt := range tests {
		name := fmt.Sprint(i)
		t.Run(name, func(t *testing.T) {
			ans := NormalAt(&tt.sphere, tt.pt)
			if !tuple.Equal(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func intersectEqual(i1, i2 Intersection) bool {
	if i1.Distance != i2.Distance {
		return false
	}
	if i1.Sphere != i2.Sphere {
		return false
	}
	return true
}

func TestHit_all_positive(t *testing.T) {
	sphere := DefaultSphere()
	i1 := Intersection{&sphere, 1}
	i2 := Intersection{&sphere, 2}
	xs := make([]Intersection, 0, 2)
	xs = InsertIntersection(xs, i1)
	xs = InsertIntersection(xs, i2)
	want := i1
	got, err := Hit(xs)
	if errors.Is(err, ErrNoHits) {
		t.Error(err)
	} else if !intersectEqual(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestHit_some_negative(t *testing.T) {
	sphere := DefaultSphere()
	i1 := Intersection{&sphere, -1}
	i2 := Intersection{&sphere, 1}
	xs := make([]Intersection, 0, 2)
	xs = InsertIntersection(xs, i1)
	xs = InsertIntersection(xs, i2)
	want := i2
	got, err := Hit(xs)
	if errors.Is(err, ErrNoHits) {
		t.Error(err)
	} else if !intersectEqual(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestHit_all_negative(t *testing.T) {
	sphere := DefaultSphere()
	i1 := Intersection{&sphere, -2}
	i2 := Intersection{&sphere, -1}
	xs := make([]Intersection, 0, 2)
	xs = InsertIntersection(xs, i1)
	xs = InsertIntersection(xs, i2)
	_, err := Hit(xs)
	if !errors.Is(err, ErrNoHits) {
		t.Errorf("want %v; got %v", ErrNoHits, err)
	}
}

func TestHit_unsorted(t *testing.T) {
	sphere := DefaultSphere()
	i1 := Intersection{&sphere, 5}
	i2 := Intersection{&sphere, 7}
	i3 := Intersection{&sphere, -3}
	i4 := Intersection{&sphere, 2}
	xs := make([]Intersection, 0, 2)
	xs = InsertIntersection(xs, i1)
	xs = InsertIntersection(xs, i2)
	xs = InsertIntersection(xs, i3)
	xs = InsertIntersection(xs, i4)
	want := i4
	got, err := Hit(xs)
	if errors.Is(err, ErrNoHits) {
		t.Error(err)
	} else if !intersectEqual(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}
