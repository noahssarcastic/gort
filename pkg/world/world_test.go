package world

import (
	"testing"

	"github.com/noahssarcastic/gort/pkg/color"
	"github.com/noahssarcastic/gort/pkg/geo"
	"github.com/noahssarcastic/gort/pkg/light"
	"github.com/noahssarcastic/gort/pkg/material"
	"github.com/noahssarcastic/gort/pkg/matrix"
	"github.com/noahssarcastic/gort/pkg/ray"
	"github.com/noahssarcastic/gort/pkg/tuple"
	"github.com/noahssarcastic/gort/pkg/util"
)

func defaultWorld() *World {
	wrld := New()
	l := light.NewPointLight(color.White, tuple.Point(-10, 10, -10))
	wrld.Lights = append(wrld.Lights, *l)
	s1 := geo.NewSphere(
		matrix.I,
		material.New(
			color.New(0.8, 1, 0.6),
			0.1,
			0.7,
			0.2,
			200,
		),
	)
	s2 := geo.NewSphere(
		matrix.Scale(0.5, 0.5, 0.5),
		material.Default(),
	)
	wrld.Spheres = append(wrld.Spheres, s1)
	wrld.Spheres = append(wrld.Spheres, s2)
	return wrld
}

func TestIntersect(t *testing.T) {
	wrld := defaultWorld()
	r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
	xs := Intersect(wrld, r)
	if len(xs) != 4 {
		t.Errorf("expected 4 intersections; got %d", len(xs))
	}
	for i, want := range []float64{4, 4.5, 5.5, 6} {
		if got := xs[i].Distance; !util.FloatEqual(got, want) {
			t.Errorf("expected intersection %d to be %f; got %f",
				i+1, want, got)
		}
	}
}

func compEqual(a, b Computation) bool {
	return a.Sphere == b.Sphere &&
		tuple.Equal(a.Point, b.Point) &&
		tuple.Equal(a.EyeVec, b.EyeVec) &&
		tuple.Equal(a.NormalVec, b.NormalVec) &&
		a.Inside == b.Inside
}

func TestGetComputation_outside(t *testing.T) {
	r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
	sphere := geo.NewSphere(matrix.I, material.Default())
	i := geo.Intersection{
		Sphere:   &sphere,
		Distance: 4,
	}
	comp := GetComputation(i, r)
	want := Computation{
		Sphere:    &sphere,
		Point:     tuple.Point(0, 0, -1),
		EyeVec:    tuple.Vector(0, 0, -1),
		NormalVec: tuple.Vector(0, 0, -1),
		Inside:    false,
	}
	if !compEqual(comp, want) {
		t.Errorf("expected %v; got %v", want, comp)
	}
}

func TestGetComputation_inside(t *testing.T) {
	r := ray.New(tuple.Point(0, 0, 0), tuple.Vector(0, 0, 1))
	sphere := geo.NewSphere(matrix.I, material.Default())
	i := geo.Intersection{
		Sphere:   &sphere,
		Distance: 1,
	}
	comp := GetComputation(i, r)
	want := Computation{
		Sphere:    &sphere,
		Point:     tuple.Point(0, 0, 1),
		EyeVec:    tuple.Vector(0, 0, -1),
		NormalVec: tuple.Vector(0, 0, -1),
		Inside:    true,
	}
	if !compEqual(comp, want) {
		t.Errorf("expected %v; got %v", want, comp)
	}
}

func TestShadeHit_outside(t *testing.T) {
	wrld := defaultWorld()
	r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
	i := geo.Intersection{
		Sphere:   &wrld.Spheres[0],
		Distance: 4,
	}
	comp := GetComputation(i, r)
	c := ShadeHit(wrld, comp)
	want := color.New(0.38066, 0.47583, 0.2855)
	if !color.Equal(c, want) {
		t.Errorf("want %v; got %v", want, c)
	}
}

func TestShadeHit_inside(t *testing.T) {
	wrld := defaultWorld()
	wrld.Lights[0] = *light.NewPointLight(color.White, tuple.Point(0, 0.25, 0))
	r := ray.New(tuple.Point(0, 0, 0), tuple.Vector(0, 0, 1))
	i := geo.Intersection{
		Sphere:   &wrld.Spheres[1],
		Distance: 0.5,
	}
	comp := GetComputation(i, r)
	c := ShadeHit(wrld, comp)
	want := color.New(0.90498, 0.90498, 0.90498)
	if !color.Equal(c, want) {
		t.Errorf("want %v; got %v", want, c)
	}
}

func TestColorAt_miss(t *testing.T) {
	wrld := defaultWorld()
	r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 1, 0))
	c := ColorAt(wrld, r)
	want := color.Black
	if !color.Equal(c, want) {
		t.Errorf("want %v; got %v", want, c)
	}
}

func TestColorAt_hit(t *testing.T) {
	wrld := defaultWorld()
	r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
	c := ColorAt(wrld, r)
	want := color.New(0.38066, 0.47583, 0.2855)
	if !color.Equal(c, want) {
		t.Errorf("want %v; got %v", want, c)
	}
}

func TestColorAt_behind(t *testing.T) {
	wrld := defaultWorld()
	wrld.Spheres[0].Material.Ambient = 1
	wrld.Spheres[1].Material.Ambient = 1
	r := ray.New(tuple.Point(0, 0, 0.75), tuple.Vector(0, 0, -1))
	c := ColorAt(wrld, r)
	want := wrld.Spheres[1].Material.Color
	if !color.Equal(c, want) {
		t.Errorf("want %v; got %v", want, c)
	}
}
