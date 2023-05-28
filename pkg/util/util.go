package util

import "math"

// Epsilon is the global machine epsilon for floating point arithmetic.
// See https://en.wikipedia.org/wiki/Machine_epsilon.
const Epsilon float64 = 0.00001

// FloatEqual returns true if two float64s are approximately equal.
func FloatEqual(a, b float64) bool {
	return math.Abs(a-b) < Epsilon
}
