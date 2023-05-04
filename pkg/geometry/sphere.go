package geometry

import (
	stdmath "math"

	"github.com/noahssarcastic/gort/pkg/math"
	"github.com/noahssarcastic/gort/pkg/ray"
)

type Sphere struct {
	origin    math.Tuple
	radius    float64
	transform math.Matrix
}

func NewSphere(mat math.Matrix) *Sphere {
	return &Sphere{math.Point(0, 0, 0), 1, mat}
}

func (sphere *Sphere) Origin() math.Tuple {
	return sphere.origin
}

func (sphere *Sphere) Radius() float64 {
	return sphere.radius
}

func (sphere *Sphere) SetTransform(mat math.Matrix) {
	sphere.transform = mat
}

func (sphere *Sphere) Intersect(r ray.Ray) []ray.Intersection {
	r = r.Transform(math.Inv(sphere.transform))
	sphereToRay := r.Origin().Sub(sphere.origin)
	a := math.Dot(r.Direction(), r.Direction())
	b := 2 * math.Dot(r.Direction(), sphereToRay)
	c := math.Dot(sphereToRay, sphereToRay) - 1
	discriminant := stdmath.Pow(b, 2) - 4*a*c
	if discriminant < 0 {
		return []ray.Intersection{}
	}
	return []ray.Intersection{
		*ray.NewIntersection((-b-stdmath.Sqrt(discriminant))/(2*a), sphere),
		*ray.NewIntersection((-b+stdmath.Sqrt(discriminant))/(2*a), sphere),
	}
}
