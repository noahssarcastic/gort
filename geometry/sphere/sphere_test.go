package sphere

import (
	"testing"

	"github.com/noahssarcastic/tddraytracer/geometry/intersec"
	"github.com/noahssarcastic/tddraytracer/geometry/ray"
	"github.com/noahssarcastic/tddraytracer/math/tuple"
	"github.com/noahssarcastic/tddraytracer/math/utils"
)

func intersectsEqual(a, b []intersec.Intersection) bool {
	if len(a) != len(b) {
		return false
	}
	for i, el := range a {
		if !utils.FloatEqual(el.Distance(), b[i].Distance()) {
			return false
		}
		if el.Object() != b[i].Object() {
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
		want  []intersec.Intersection
	}{
		{"normal", tuple.Point(0, 0, -5), []intersec.Intersection{
			intersec.New(4, &sphere),
			intersec.New(6, &sphere),
		}},
		{"tangent", tuple.Point(0, 1, -5), []intersec.Intersection{
			intersec.New(5, &sphere),
			intersec.New(5, &sphere),
		}},
		{"miss", tuple.Point(0, 2, -5), []intersec.Intersection{}},
		{"inside", tuple.Point(0, 0, 0), []intersec.Intersection{
			intersec.New(-1, &sphere),
			intersec.New(1, &sphere),
		}},
		{"behind", tuple.Point(0, 0, 5), []intersec.Intersection{
			intersec.New(-6, &sphere),
			intersec.New(-4, &sphere),
		}},
	}
	for _, tt := range tests {
		name := tt.name
		t.Run(name, func(t *testing.T) {
			ans := sphere.Intersect(ray.New(tt.start, tuple.Vector(0, 0, 1)))
			if !intersectsEqual(tt.want, ans) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
