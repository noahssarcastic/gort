package geometry

import (
	"testing"

	"github.com/noahssarcastic/gort/math"
)

func intersectsEqual(a, b []Intersection) bool {
	if len(a) != len(b) {
		return false
	}
	for i, el := range a {
		if !math.FloatEqual(el.Distance(), b[i].Distance()) {
			return false
		}
		if el.Object() != b[i].Object() {
			return false
		}
	}
	return true
}

func TestIntersect_inner(t *testing.T) {
	sphere := NewSphere()
	tests := []struct {
		name  string
		start math.Tuple
		want  []Intersection
	}{
		{"normal", math.Point(0, 0, -5), []Intersection{
			*NewIntersection(4, &sphere),
			*NewIntersection(6, &sphere),
		}},
		{"tangent", math.Point(0, 1, -5), []Intersection{
			*NewIntersection(5, &sphere),
			*NewIntersection(5, &sphere),
		}},
		{"miss", math.Point(0, 2, -5), []Intersection{}},
		{"inside", math.Point(0, 0, 0), []Intersection{
			*NewIntersection(-1, &sphere),
			*NewIntersection(1, &sphere),
		}},
		{"behind", math.Point(0, 0, 5), []Intersection{
			*NewIntersection(-6, &sphere),
			*NewIntersection(-4, &sphere),
		}},
	}
	for _, tt := range tests {
		name := tt.name
		t.Run(name, func(t *testing.T) {
			ans := sphere.Intersect(NewRay(tt.start, math.Vector(0, 0, 1)))
			if !intersectsEqual(tt.want, ans) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
