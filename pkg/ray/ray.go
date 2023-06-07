// Package ray implements Ray
package ray

import (
	"github.com/noahssarcastic/gort/pkg/matrix"
	"github.com/noahssarcastic/gort/pkg/tuple"
)

// Ray represents a 3D ray. Ray is immutable.
type Ray struct {
	origin    tuple.Tuple
	direction tuple.Tuple
}

// New creates a 3D ray originating from origin and pointing in direction.
func New(origin tuple.Tuple, direction tuple.Tuple) Ray {
	return Ray{
		origin,
		direction,
	}
}

// Origin returns the origin of a Ray.
func (ray Ray) Origin() tuple.Tuple {
	return ray.origin
}

// Direction returns the direction of a Ray.
func (ray Ray) Direction() tuple.Tuple {
	return ray.direction
}

// Position calculates the coordinates a distance t from the origin.
func Position(ray Ray, t float64) tuple.Tuple {
	return tuple.Add(ray.origin, tuple.Mult(ray.direction, t))
}

// Transform creates a new Ray which is the product of Matrix tform and Ray r.
func Transform(r Ray, tform matrix.Matrix) Ray {
	return Ray{
		origin:    tform.Apply(r.origin),
		direction: tform.Apply(r.direction),
	}
}
