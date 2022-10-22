package utils

import "math"

const EPSILON float64 = 0.00001

func FloatEqual(a, b float64) bool {
	return math.Abs(a-b) < EPSILON
}
