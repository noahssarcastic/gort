package ray

import (
	"sort"
)

// Allow all geometries to be intersected.
type Object interface {
	// Get a list of intersections made with the given Ray.
	Intersect(ray Ray) []Intersection
}

// Store distance of intersection and pointer to intersected object.
type Intersection struct {
	t      float64 // Distance from ray origin to intersection.
	object Object  // Pointer to intersected object.
}

func NewIntersection(t float64, obj Object) *Intersection {
	return &Intersection{t, obj}
}

// Get the distance from ray origin to intersection.
func (x Intersection) Distance() float64 {
	return x.t
}

// Get a pointer to intersected object.
func (x Intersection) Object() Object {
	return x.object
}

func search(xs []Intersection, new *Intersection) int {
	return sort.Search(len(xs), func(i int) bool {
		return xs[i].t >= new.t
	})
}

func insertAt(xs []Intersection, i int, new *Intersection) []Intersection {
	if i == len(xs) {
		return append(xs, *new)
	}
	xs = append(xs[:i+1], xs[i:]...)
	xs[i] = *new
	return xs
}

func InsertIntersection(xs []Intersection, new *Intersection) []Intersection {
	foundAt := search(xs, new)
	return insertAt(xs, foundAt, new)
}

// Get a pointer to the closest non-negative intersection.
func Hit(xs []Intersection) *Intersection {
	i := sort.Search(len(xs), func(i int) bool {
		return xs[i].t >= 0
	})
	if i >= len(xs) {
		return nil
	}
	return &xs[i]
}
