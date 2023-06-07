package world

import (
	"fmt"

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
	panic(fmt.Errorf("not implemented"))
}

func GetComputation(i geo.Intersection, r ray.Ray) Computation {
	panic(fmt.Errorf("not implemented"))
}

func ShadeHit(wrld *World, comp Computation) color.Color {
	panic(fmt.Errorf("not implemented"))
}

func ColorAt(wrld *World, r ray.Ray) color.Color {
	panic(fmt.Errorf("not implemented"))
}
