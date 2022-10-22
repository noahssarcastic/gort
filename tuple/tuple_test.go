package tuple

import (
	"fmt"
	"testing"

	"github.com/noahssarcastic/tddraytracer/utils"
)

func TestEqual(t *testing.T) {
	var tests = []struct {
		a, b Tuple
		want bool
	}{
		{Vector(1, 2, 3), Point(1, 2, 3), false},
		{Vector(1, 2, 3), Vector(1, 2, 3), true},
		{Vector(1, 2, 3), Vector(1.1, 2, 3), false},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("%v,%v", tt.a, tt.b)
		t.Run(name, func(t *testing.T) {
			ans := Equal(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestPointConstructor(t *testing.T) {
	point := Point(1, 2, 3)
	want := true
	got := point.IsPoint()
	if want != got {
		t.Errorf("Point(1, 2, 3).IsPoint() = %v; want %v", got, want)
	}
}

func TestVectorConstructor(t *testing.T) {
	vec := Vector(5, 6, 7)
	want := true
	got := vec.IsVector()
	if want != got {
		t.Errorf("Vector(5, 6, 7).IsVector() = %v; want %v", got, want)
	}
}

func TestAddVectors(t *testing.T) {
	v1 := Vector(1, 2, 3)
	v2 := Vector(4, 5, 6)
	want := Vector(5, 7, 9)
	got := v1.Add(v2)
	if !Equal(want, got) {
		t.Errorf("Vector(1, 2, 3).Add(Vector(4, 5, 6)) = %v; want %v", got, want)
	}
}

func TestAddPointAndVector(t *testing.T) {
	pt := Point(1, 2, 3)
	vec := Vector(4, 5, 6)
	want := Point(5, 7, 9)
	got := pt.Add(vec)
	if !Equal(want, got) {
		t.Errorf("Point(1, 2, 3).Add(Vector(4, 5, 6)) = %v; want %v", got, want)
	}
}

func TestSubVectors(t *testing.T) {
	v1 := Vector(1, 2, 3)
	v2 := Vector(4, 5, 6)
	want := Vector(-3, -3, -3)
	got := v1.Subtract(v2)
	if !Equal(want, got) {
		t.Errorf("Vector(1, 2, 3).Subtract(Vector(4, 5, 6)) = %v; want %v", got, want)
	}
}

func TestSubtractPoints(t *testing.T) {
	p1 := Point(1, 2, 3)
	p2 := Point(4, 5, 6)
	want := Vector(-3, -3, -3)
	got := p1.Subtract(p2)
	if !Equal(want, got) {
		t.Errorf("Point(1, 2, 3).Subtract(Point(4, 5, 6)) = %v; want %v", got, want)
	}
}

func TestSubtractVectorFromPoint(t *testing.T) {
	pt := Point(1, 2, 3)
	vec := Vector(4, 5, 6)
	want := Point(-3, -3, -3)
	got := pt.Subtract(vec)
	if !Equal(want, got) {
		t.Errorf("Point(1, 2, 3).Subtract(Vector(4, 5, 6)) = %v; want %v", got, want)
	}
}

func TestNegateVector(t *testing.T) {
	vec := Vector(1, 2, 3)
	want := Vector(-1, -2, -3)
	got := Negate(vec)
	if !Equal(want, got) {
		t.Errorf("Negate(Vector(1, 2, 3)) = %v; want %v", got, want)
	}
}

func TestMultiplyVector(t *testing.T) {
	vec := Vector(1, 2, 3)
	scalar := 2.0
	want := Vector(2, 4, 6)
	got := vec.Multiply(scalar)
	if !Equal(want, got) {
		t.Errorf("Vector(1, 2, 3).Multiply(2) = %v; want %v", got, want)
	}
}

func TestDivideVector(t *testing.T) {
	vec := Vector(1, 2, 3)
	scalar := 2.0
	want := Vector(0.5, 1, 1.5)
	got := vec.Divide(scalar)
	if !Equal(want, got) {
		t.Errorf("Vector(1, 2, 3).Divide(2) = %v; want %v", got, want)
	}
}

func TestMagnitude(t *testing.T) {
	var tests = []struct {
		t    Tuple
		want float64
	}{
		{Vector(1, 0, 0), 1},
		{Vector(0, 1, 0), 1},
		{Vector(0, 0, 1), 1},
		{Vector(-4, -8, -8), 12},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("%v", tt.t)
		t.Run(name, func(t *testing.T) {
			ans := Magnitude(tt.t)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestNormalize(t *testing.T) {
	var tests = []struct {
		t    Tuple
		want Tuple
	}{
		{Vector(10, 0, 0), Vector(1, 0, 0)},
		{Vector(4, 8, 8), Vector(1./3., 2./3., 2./3.)},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("%v", tt.t)
		t.Run(name, func(t *testing.T) {
			ans := Normalize(tt.t)
			if !Equal(tt.want, ans) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestMagnitudeOfUnitVector(t *testing.T) {
	vec := Vector(1, 2, 3)
	unitVec := Normalize(vec)
	want := 1.
	got := Magnitude(unitVec)
	if !utils.FloatEqual(want, got) {
		t.Errorf("Magnitude(Normalize(Vector(1, 2, 3))) = %v; want %v", got, want)
	}
}

func TestDotProduct(t *testing.T) {
	v1 := Vector(1, 2, 3)
	v2 := Vector(2, 3, 4)
	want := 20.
	got := Dot(v1, v2)
	if !utils.FloatEqual(want, got) {
		t.Errorf("Magnitude(Normalize(Vector(1, 2, 3))) = %v; want %v", got, want)
	}
}

func TestCrossProduct(t *testing.T) {
	var tests = []struct {
		t1, t2 Tuple
		want   Tuple
	}{
		{Vector(1, 0, 0), Vector(0, 1, 0), Vector(0, 0, 1)},
		{Vector(0, 1, 0), Vector(1, 0, 0), Vector(0, 0, -1)},
		{Vector(1, 2, 3), Vector(2, 3, 4), Vector(-1, 2, -1)},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("%v,%v", tt.t1, tt.t2)
		t.Run(name, func(t *testing.T) {
			ans := Cross(tt.t1, tt.t2)
			if !Equal(tt.want, ans) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
