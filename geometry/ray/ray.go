package ray

import "github.com/noahssarcastic/tddraytracer/math/tuple"

type Ray struct {
	origin    tuple.Tuple
	direction tuple.Tuple
}

func New(origin tuple.Tuple, direction tuple.Tuple) Ray {
	return Ray{
		origin,
		direction,
	}
}

func Position(ray Ray, t float64) tuple.Tuple {
	return ray.origin.Add(ray.direction.Multiply(t))
}
