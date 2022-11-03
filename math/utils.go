package math

import stdmath "math"

const EPSILON float64 = 0.00001

func FloatEqual(a, b float64) bool {
	return stdmath.Abs(a-b) < EPSILON
}
