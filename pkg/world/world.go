package world

import (
	"errors"

	"github.com/noahssarcastic/gort/pkg/color"
	"github.com/noahssarcastic/gort/pkg/geo"
	"github.com/noahssarcastic/gort/pkg/light"
	"github.com/noahssarcastic/gort/pkg/ray"
	"github.com/noahssarcastic/gort/pkg/tuple"
)

type World struct {
	Lights  []light.PointLight
	Spheres []geo.Sphere
}

type Computation struct {
	Sphere    *geo.Sphere
	Point     tuple.Tuple
	EyeVec    tuple.Tuple
	NormalVec tuple.Tuple
	Inside    bool
}

func New() *World {
	return &World{}
}

func Intersect(w *World, r ray.Ray) []geo.Intersection {
	xs := make([]geo.Intersection, 0)
	for s := range w.Spheres {
		is := geo.Intersect(&w.Spheres[s], r)
		for _, x := range is {
			xs = geo.InsertIntersection(xs, x)
		}
	}
	return xs
}

func GetComputation(i geo.Intersection, r ray.Ray) Computation {
	sphere := i.Sphere
	pt := ray.Position(r, i.Distance)
	eye := tuple.Neg(r.Direction())
	normal := geo.NormalAt(i.Sphere, pt)
	inside := tuple.Dot(eye, normal) < 0
	if inside {
		normal = tuple.Neg(normal)
	}
	return Computation{
		Sphere:    sphere,
		Point:     pt,
		EyeVec:    eye,
		NormalVec: normal,
		Inside:    inside,
	}
}

func ShadeHit(wrld *World, comp Computation) color.Color {
	var c color.Color
	for _, l := range wrld.Lights {
		c = c.Add(
			light.Lighting(
				comp.Sphere.Material,
				comp.Point,
				&l,
				comp.EyeVec,
				comp.NormalVec,
			),
		)
	}
	return c
}

func ColorAt(wrld *World, r ray.Ray) color.Color {
	xs := Intersect(wrld, r)
	i, err := geo.Hit(xs)
	if errors.Is(err, geo.ErrNoHits) {
		return color.Black
	}
	comp := GetComputation(i, r)
	return ShadeHit(wrld, comp)
}
