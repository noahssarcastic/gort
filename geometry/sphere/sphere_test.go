package sphere

import (
	"testing"

	"github.com/noahssarcastic/tddraytracer/geometry/ray"
	"github.com/noahssarcastic/tddraytracer/math/tuple"
	"github.com/noahssarcastic/tddraytracer/math/utils"
)

func intersectEqual(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, el := range a {
		if !utils.FloatEqual(el, b[i]) {
			return false
		}
	}
	return true
}

func TestIntersect_inner(t *testing.T) {
	sphere := New()
	tests := []struct {
		name  string
		start tuple.Tuple
		want  []float64
	}{
		{"normal", tuple.Point(0, 0, -5), []float64{4, 6}},
		{"tangent", tuple.Point(0, 1, -5), []float64{5, 5}},
		{"miss", tuple.Point(0, 2, -5), []float64{}},
		{"inside", tuple.Point(0, 0, 0), []float64{-1, 1}},
		{"behind", tuple.Point(0, 0, 5), []float64{-6, -4}},
	}
	for _, tt := range tests {
		name := tt.name
		t.Run(name, func(t *testing.T) {
			ans := sphere.Intersect(ray.New(tt.start, tuple.Vector(0, 0, 1)))
			if !intersectEqual(tt.want, ans) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
