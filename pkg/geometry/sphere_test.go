package geometry

import (
	"testing"

	"github.com/noahssarcastic/gort/pkg/math"
	"github.com/noahssarcastic/gort/pkg/ray"
)

func intersectsEqual(a, b []ray.Intersection) bool {
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

func TestIntersect(t *testing.T) {
	sphere := NewSphere(math.I())
	tests := []struct {
		name  string
		start math.Tuple
		want  []ray.Intersection
	}{
		{"normal", math.Point(0, 0, -5), []ray.Intersection{
			*ray.NewIntersection(4, sphere),
			*ray.NewIntersection(6, sphere),
		}},
		{"tangent", math.Point(0, 1, -5), []ray.Intersection{
			*ray.NewIntersection(5, sphere),
			*ray.NewIntersection(5, sphere),
		}},
		{"miss", math.Point(0, 2, -5), []ray.Intersection{}},
		{"inside", math.Point(0, 0, 0), []ray.Intersection{
			*ray.NewIntersection(-1, sphere),
			*ray.NewIntersection(1, sphere),
		}},
		{"behind", math.Point(0, 0, 5), []ray.Intersection{
			*ray.NewIntersection(-6, sphere),
			*ray.NewIntersection(-4, sphere),
		}},
	}
	for _, tt := range tests {
		name := tt.name
		t.Run(name, func(t *testing.T) {
			ans := sphere.Intersect(ray.NewRay(tt.start, math.Vector(0, 0, 1)))
			if !intersectsEqual(tt.want, ans) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestIntersect_transformed(t *testing.T) {
	sphere := NewSphere(math.I())
	tests := []struct {
		name  string
		tform math.Matrix
		start math.Tuple
		want  []ray.Intersection
	}{
		{
			"scale",
			math.Scale(2, 2, 2),
			math.Point(0, 0, -5),
			[]ray.Intersection{
				*ray.NewIntersection(3, sphere),
				*ray.NewIntersection(7, sphere),
			},
		},
		{
			"translate",
			math.Translate(5, 0, 0),
			math.Point(0, 0, -5),
			[]ray.Intersection{},
		},
	}
	for _, tt := range tests {
		name := tt.name
		t.Run(name, func(t *testing.T) {
			sphere.SetTransform(tt.tform)
			ans := sphere.Intersect(ray.NewRay(tt.start, math.Vector(0, 0, 1)))
			if !intersectsEqual(tt.want, ans) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
