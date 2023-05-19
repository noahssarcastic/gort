package geo

import (
	"math"

	"github.com/noahssarcastic/gort/pkg/material"
	"github.com/noahssarcastic/gort/pkg/matrix"
	"github.com/noahssarcastic/gort/pkg/ray"
	"github.com/noahssarcastic/gort/pkg/tuple"
)

// A Sphere represents a unit sphere at the origin.
type Sphere struct {
	center    tuple.Tuple
	radius    float64
	transform matrix.Matrix
	material  material.Material
}

// NewSphere creates a sphere with a given transformation matrix tform.
func NewSphere() *Sphere {
	return &Sphere{tuple.Point(0, 0, 0), 1, matrix.I(), material.Default()}
}

// SetTransform overwrites the Sphere's transformation matrix.
func (sphere *Sphere) SetTransform(mat matrix.Matrix) {
	sphere.transform = mat
}

// SetMaterial overwrites the Sphere's material.
func (sphere *Sphere) SetMaterial(mat material.Material) {
	sphere.material = mat
}

// Intersect takes a ray and returns an array of intersections. Intersections
// can be in both the positive and negative direction of the ray.
func (sphere *Sphere) Intersect(r ray.Ray) []ray.Intersect {
	r = ray.Transform(r, matrix.Inv(sphere.transform))
	sphereToRay := tuple.Sub(r.Origin(), sphere.center)
	a := tuple.Dot(r.Direction(), r.Direction())
	b := 2 * tuple.Dot(r.Direction(), sphereToRay)
	c := tuple.Dot(sphereToRay, sphereToRay) - 1
	discriminant := math.Pow(b, 2) - 4*a*c
	if discriminant < 0 {
		return []ray.Intersect{}
	}
	return []ray.Intersect{
		ray.NewIntersect((-b-math.Sqrt(discriminant))/(2*a), sphere),
		ray.NewIntersect((-b+math.Sqrt(discriminant))/(2*a), sphere),
	}
}

// NormalAt returns the normal vector at a given point along the Sphere.
// The returned vector is normalized. Passing a point not on the surface of the
// Sphere is undefined.
func (sphere *Sphere) NormalAt(pt tuple.Tuple) tuple.Tuple {
	ptObjSpace := matrix.Inv(sphere.transform).Apply(pt)
	normalObjSpace := tuple.Sub(ptObjSpace, tuple.Point(0, 0, 0))
	normalWrldSpace := matrix.Inv(sphere.transform).T().Apply(normalObjSpace)
	// reset w component if mangled by transpose
	return tuple.Norm(tuple.Vector(
		normalWrldSpace.X(), normalWrldSpace.Y(), normalWrldSpace.Z()))
}

func (sphere *Sphere) Material() material.Material {
	return sphere.material
}
