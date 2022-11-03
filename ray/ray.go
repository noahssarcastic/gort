package ray

import "github.com/noahssarcastic/gort/math"

type Ray struct {
	origin    math.Tuple
	direction math.Tuple
}

func NewRay(origin math.Tuple, direction math.Tuple) Ray {
	return Ray{
		origin,
		direction,
	}
}

func (ray Ray) Origin() math.Tuple {
	return ray.origin
}

func (ray Ray) Direction() math.Tuple {
	return ray.direction
}

func Position(ray Ray, t float64) math.Tuple {
	return ray.origin.Add(ray.direction.Mult(t))
}

func (r Ray) Transform(mat math.Matrix) Ray {
	return Ray{
		origin:    mat.Apply(r.origin),
		direction: mat.Apply(r.direction),
	}
}
