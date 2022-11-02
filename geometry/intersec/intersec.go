package intersec

import "github.com/noahssarcastic/tddraytracer/geometry/ray"

type Object interface {
	Intersect(ray ray.Ray) []Intersection
}

type Intersection struct {
	t      float64
	object Object
}

func New(t float64, obj Object) Intersection {
	return Intersection{t, obj}
}

func (x Intersection) Distance() float64 {
	return x.t
}

func (x Intersection) Object() Object {
	return x.object
}

// TODO: determine if this is needed
func Intersections(xs ...Intersection) []Intersection {
	return xs
}
