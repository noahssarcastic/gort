package sphere

import (
	"math"

	"github.com/noahssarcastic/tddraytracer/geometry/ray"
	"github.com/noahssarcastic/tddraytracer/math/tuple"
)

type Sphere struct {
	origin tuple.Tuple
	radius float64
}

func New() Sphere {
	return Sphere{tuple.Point(0, 0, 0), 1}
}

func (sphere *Sphere) Origin() tuple.Tuple {
	return sphere.origin
}

func (sphere *Sphere) Radius() float64 {
	return sphere.radius
}

func (sphere *Sphere) Intersect(ray ray.Ray) []float64 {
	sphereToRay := ray.Origin().Subtract(sphere.origin)
	a := tuple.Dot(ray.Direction(), ray.Direction())
	b := 2 * tuple.Dot(ray.Direction(), sphereToRay)
	c := tuple.Dot(sphereToRay, sphereToRay) - 1
	discriminant := math.Pow(b, 2) - 4*a*c
	if discriminant < 0 {
		return []float64{}
	}
	return []float64{
		((-b - math.Sqrt(discriminant)) / (2 * a)),
		((-b + math.Sqrt(discriminant)) / (2 * a)),
	}
}
