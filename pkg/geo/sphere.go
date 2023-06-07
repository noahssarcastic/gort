package geo

import (
	"errors"
	"math"
	"sort"

	"github.com/noahssarcastic/gort/pkg/material"
	"github.com/noahssarcastic/gort/pkg/matrix"
	"github.com/noahssarcastic/gort/pkg/ray"
	"github.com/noahssarcastic/gort/pkg/tuple"
)

// A Sphere represents a unit sphere at the origin.
type Sphere struct {
	center    tuple.Tuple
	radius    float64
	Transform matrix.Matrix
	Material  material.Material
}

// NewSphere creates a sphere with a given transformation matrix tform.
func NewSphere(tform matrix.Matrix, mat material.Material) Sphere {
	return Sphere{tuple.Point(0, 0, 0), 1, tform, mat}
}

func DefaultSphere() Sphere {
	return Sphere{tuple.Point(0, 0, 0), 1, matrix.I, material.Default()}
}

// Intersect takes a ray and returns an array of intersections. Intersections
// can be in both the positive and negative direction of the ray.
func Intersect(sphere *Sphere, r ray.Ray) []Intersection {
	r = ray.Transform(r, matrix.Inv(sphere.Transform))
	sphereToRay := tuple.Sub(r.Origin(), sphere.center)
	a := tuple.Dot(r.Direction(), r.Direction())
	b := 2 * tuple.Dot(r.Direction(), sphereToRay)
	c := tuple.Dot(sphereToRay, sphereToRay) - 1
	discriminant := math.Pow(b, 2) - 4*a*c
	if discriminant < 0 {
		return []Intersection{}
	}
	return []Intersection{
		{sphere, (-b - math.Sqrt(discriminant)) / (2 * a)},
		{sphere, (-b + math.Sqrt(discriminant)) / (2 * a)},
	}
}

// NormalAt returns the normal vector at a given point along the Sphere.
// The returned vector is normalized. Passing a point not on the surface of the
// Sphere is undefined.
func NormalAt(sphere *Sphere, pt tuple.Tuple) tuple.Tuple {
	ptObjSpace := matrix.Inv(sphere.Transform).Apply(pt)
	normalObjSpace := tuple.Sub(ptObjSpace, tuple.Point(0, 0, 0))
	normalWrldSpace := matrix.Inv(sphere.Transform).T().Apply(normalObjSpace)
	// reset w component if mangled by transpose
	return tuple.Norm(tuple.Vector(
		normalWrldSpace.X(), normalWrldSpace.Y(), normalWrldSpace.Z()))
}

type Intersection struct {
	Sphere   *Sphere
	Distance float64
}

func search(xs []Intersection, new Intersection) int {
	return sort.Search(len(xs), func(i int) bool {
		return xs[i].Distance >= new.Distance
	})
}

// TODO: test this better, might be broken
func insertAt(xs []Intersection, i int, new Intersection) []Intersection {
	if i == len(xs) {
		return append(xs, new)
	}
	xs = append(xs[:i+1], xs[i:]...)
	xs[i] = new
	return xs
}

func InsertIntersection(xs []Intersection, new Intersection) []Intersection {
	foundAt := search(xs, new)
	return insertAt(xs, foundAt, new)
}

var ErrNoHits = errors.New("no non-negative intersections found")

// Hit returns the closest non-negative intersection.
func Hit(xs []Intersection) (Intersection, error) {
	i := sort.Search(len(xs), func(i int) bool {
		return xs[i].Distance >= 0
	})
	if i >= len(xs) {
		return Intersection{}, ErrNoHits
	}
	return xs[i], nil
}
