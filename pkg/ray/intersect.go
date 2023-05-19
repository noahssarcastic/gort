package ray

import (
	"errors"
	"sort"

	"github.com/noahssarcastic/gort/pkg/material"
	"github.com/noahssarcastic/gort/pkg/tuple"
)

// Intersectable allows geometries to be intersected by rays.
type Intersectable interface {
	// Intersect returns an array of intersections made with ray.
	Intersect(ray Ray) []Intersect
	NormalAt(pt tuple.Tuple) tuple.Tuple
	Material() material.Material
}

// Intersect represents an intersection between a Ray and an Intersectable.
// Intersects are immutable.
type Intersect struct {
	t      float64
	objPtr Intersectable
}

// NewIntersect creates an Intersect that is t along some Ray.
func NewIntersect(t float64, objPtr Intersectable) Intersect {
	return Intersect{t, objPtr}
}

// Distance returns the distance from ray origin to intersection.
func (x Intersect) Distance() float64 {
	return x.t
}

// Object returns a pointer to the intersected object.
func (x Intersect) Object() Intersectable {
	return x.objPtr
}

// search calculates the insertion index an Intersect.
func search(xs []Intersect, new Intersect) int {
	return sort.Search(len(xs), func(i int) bool {
		return xs[i].t >= new.t
	})
}

func insertAt(xs []Intersect, i int, new Intersect) []Intersect {
	if i == len(xs) {
		return append(xs, new)
	}
	xs = append(xs[:i+1], xs[i:]...)
	xs[i] = new
	return xs
}

// InsertIntersect adds an Intersect to a []Intersect in sorted order.
func InsertIntersect(xs []Intersect, new Intersect) []Intersect {
	foundAt := search(xs, new)
	return insertAt(xs, foundAt, new)
}

var ErrNoHits = errors.New("no non-negative intersections found")

// Hit returns the closest non-negative intersection.
func Hit(xs []Intersect) (Intersect, error) {
	i := sort.Search(len(xs), func(i int) bool {
		return xs[i].t >= 0
	})
	if i >= len(xs) {
		return Intersect{}, ErrNoHits
	}
	return xs[i], nil
}
