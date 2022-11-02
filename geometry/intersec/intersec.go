package intersec

import (
	"sort"

	"github.com/noahssarcastic/tddraytracer/geometry/ray"
)

// Allow all geometries to be intersected.
type Interface interface {
	// Get a list of intersections made with the given Ray.
	Intersect(ray ray.Ray) []Intersection
}

// Store distance of intersection and pointer to intersected object.
type Intersection struct {
	t      float64   // Distance from ray origin to intersection.
	object Interface // Pointer to intersected object.
}

func New(t float64, obj Interface) *Intersection {
	return &Intersection{t, obj}
}

// Get the distance from ray origin to intersection.
func (x Intersection) Distance() float64 {
	return x.t
}

// Get a pointer to intersected object.
func (x Intersection) Object() Interface {
	return x.object
}

// A sorted list of Intersections.
type HitList struct {
	xs []*Intersection
}

// Implement sort.Interface
func (a HitList) Len() int           { return len(a.xs) }
func (a HitList) Swap(i, j int)      { a.xs[i], a.xs[j] = a.xs[j], a.xs[i] }
func (a HitList) Less(i, j int) bool { return a.xs[i].t < a.xs[j].t }

func search(xs []*Intersection, new *Intersection) int {
	return sort.Search(len(xs), func(i int) bool {
		return xs[i].t >= new.t
	})
}

func insertAt(xs []*Intersection, i int, new *Intersection) []*Intersection {
	if i == len(xs) {
		return append(xs, new)
	}
	xs = append(xs[:i+1], xs[i:]...)
	xs[i] = new
	return xs
}

func insertSorted(xs []*Intersection, new *Intersection) []*Intersection {
	foundAt := search(xs, new)
	return insertAt(xs, foundAt, new)
}

// Collate intersection in sorted order.
func Combine(xs ...*Intersection) HitList {
	sortedList := make([]*Intersection, 0, len(xs))
	for _, el := range xs {
		sortedList = insertSorted(sortedList, el)
	}
	return HitList{sortedList}
}

// Get a pointer to the closest non-negative intersection.
func (hl HitList) Hit() *Intersection {
	xs := hl.xs
	i := sort.Search(len(hl.xs), func(i int) bool {
		return hl.xs[i].t >= 0
	})
	if i >= len(xs) {
		return nil
	}
	return xs[i]
}
