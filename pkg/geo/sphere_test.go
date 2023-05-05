package geo

import (
	"testing"

	"github.com/noahssarcastic/gort/pkg/mat"
	"github.com/noahssarcastic/gort/pkg/ray"
	"github.com/noahssarcastic/gort/pkg/tuple"
	"github.com/noahssarcastic/gort/pkg/util"
)

func intersectsEqual(a, b []ray.Intersect) bool {
	if len(a) != len(b) {
		return false
	}
	for i, el := range a {
		if !util.FloatEqual(el.Distance(), b[i].Distance()) {
			return false
		}
		if el.Object() != b[i].Object() {
			return false
		}
	}
	return true
}

func TestIntersect(t *testing.T) {
	sphere := NewSphere(mat.I())
	tests := []struct {
		name  string
		start tuple.Tuple
		want  []ray.Intersect
	}{
		{"normal", tuple.Point(0, 0, -5), []ray.Intersect{
			ray.NewIntersect(4, sphere),
			ray.NewIntersect(6, sphere),
		}},
		{"tangent", tuple.Point(0, 1, -5), []ray.Intersect{
			ray.NewIntersect(5, sphere),
			ray.NewIntersect(5, sphere),
		}},
		{"miss", tuple.Point(0, 2, -5), []ray.Intersect{}},
		{"inside", tuple.Point(0, 0, 0), []ray.Intersect{
			ray.NewIntersect(-1, sphere),
			ray.NewIntersect(1, sphere),
		}},
		{"behind", tuple.Point(0, 0, 5), []ray.Intersect{
			ray.NewIntersect(-6, sphere),
			ray.NewIntersect(-4, sphere),
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

func TestIntersect_transformed(t *testing.T) {
	sphere := NewSphere(mat.I())
	tests := []struct {
		name  string
		tform mat.Matrix
		start tuple.Tuple
		want  []ray.Intersect
	}{
		{
			"scale",
			mat.Scale(2, 2, 2),
			tuple.Point(0, 0, -5),
			[]ray.Intersect{
				ray.NewIntersect(3, sphere),
				ray.NewIntersect(7, sphere),
			},
		},
		{
			"translate",
			mat.Translate(5, 0, 0),
			tuple.Point(0, 0, -5),
			[]ray.Intersect{},
		},
	}
	for _, tt := range tests {
		name := tt.name
		t.Run(name, func(t *testing.T) {
			sphere.SetTransform(tt.tform)
			ans := sphere.Intersect(ray.New(tt.start, tuple.Vector(0, 0, 1)))
			if !intersectsEqual(tt.want, ans) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
