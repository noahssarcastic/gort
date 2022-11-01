package sphere

import (
	"github.com/noahssarcastic/tddraytracer/math/tuple"
)

type Sphere struct {
	origin tuple.Tuple
	radius float64
}

func New() Sphere {
	return Sphere{tuple.Point(0, 0, 0), 1}
}

// func (sphere *Sphere) Intersect(ray ray.Ray) []float64 {

// }
