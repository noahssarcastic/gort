package geo

import (
	"fmt"
	"math"
	"testing"

	"github.com/noahssarcastic/gort/pkg/matrix"
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
	sphere := NewSphere()
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
	sphere := NewSphere()
	tests := []struct {
		name  string
		tform matrix.Matrix
		start tuple.Tuple
		want  []ray.Intersect
	}{
		{
			"scale",
			matrix.Scale(2, 2, 2),
			tuple.Point(0, 0, -5),
			[]ray.Intersect{
				ray.NewIntersect(3, sphere),
				ray.NewIntersect(7, sphere),
			},
		},
		{
			"translate",
			matrix.Translate(5, 0, 0),
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

func TestNormalAt(t *testing.T) {
	sphere := NewSphere()
	tests := []struct {
		pt   tuple.Tuple
		want tuple.Tuple
	}{
		{tuple.Point(1, 0, 0), tuple.Vector(1, 0, 0)},
		{tuple.Point(0, 1, 0), tuple.Vector(0, 1, 0)},
		{tuple.Point(0, 0, 1), tuple.Vector(0, 0, 1)},
		{
			tuple.Point(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3),
			tuple.Vector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3),
		},
	}
	for i, tt := range tests {
		name := fmt.Sprint(i)
		t.Run(name, func(t *testing.T) {
			ans := sphere.NormalAt(tt.pt)
			if !tuple.Equal(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestNormalAt_normalized(t *testing.T) {
	sphere := NewSphere()
	pt := tuple.Point(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3)
	normal := sphere.NormalAt(pt)
	want := tuple.Norm(normal)
	got := normal
	if !tuple.Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestNormalAt_transformed(t *testing.T) {
	tests := []struct {
		pt    tuple.Tuple
		tform matrix.Matrix
		want  tuple.Tuple
	}{
		{
			tuple.Point(0, 1.70711, -0.70711),
			matrix.Translate(0, 1, 0),
			tuple.Vector(0, 0.70711, -0.70711),
		},
		{
			tuple.Point(0, math.Sqrt(2)/2, -math.Sqrt(2)/2),
			matrix.Chain(
				matrix.RotateZ(math.Pi/5),
				matrix.Scale(1, 0.5, 1),
			),
			tuple.Vector(0, 0.97014, -0.24254),
		},
	}
	for i, tt := range tests {
		name := fmt.Sprint(i)
		t.Run(name, func(t *testing.T) {
			sphere := NewSphere()
			sphere.SetTransform(tt.tform)
			ans := sphere.NormalAt(tt.pt)
			if !tuple.Equal(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
