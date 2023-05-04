package math

import (
	"fmt"
	"testing"
)

func TestTupleEqual(t *testing.T) {
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
			ans := TupleEqual(tt.a, tt.b)
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
	if !TupleEqual(want, got) {
		t.Errorf("Vector(1, 2, 3).Add(Vector(4, 5, 6)) = %v; want %v", got, want)
	}
}

func TestAddPointAndVector(t *testing.T) {
	pt := Point(1, 2, 3)
	vec := Vector(4, 5, 6)
	want := Point(5, 7, 9)
	got := pt.Add(vec)
	if !TupleEqual(want, got) {
		t.Errorf("Point(1, 2, 3).Add(Vector(4, 5, 6)) = %v; want %v", got, want)
	}
}

func TestSubVectors(t *testing.T) {
	v1 := Vector(1, 2, 3)
	v2 := Vector(4, 5, 6)
	want := Vector(-3, -3, -3)
	got := v1.Sub(v2)
	if !TupleEqual(want, got) {
		t.Errorf("Vector(1, 2, 3).Sub(Vector(4, 5, 6)) = %v; want %v", got, want)
	}
}

func TestSubPoints(t *testing.T) {
	p1 := Point(1, 2, 3)
	p2 := Point(4, 5, 6)
	want := Vector(-3, -3, -3)
	got := p1.Sub(p2)
	if !TupleEqual(want, got) {
		t.Errorf("Point(1, 2, 3).Sub(Point(4, 5, 6)) = %v; want %v", got, want)
	}
}

func TestSubVectorFromPoint(t *testing.T) {
	pt := Point(1, 2, 3)
	vec := Vector(4, 5, 6)
	want := Point(-3, -3, -3)
	got := pt.Sub(vec)
	if !TupleEqual(want, got) {
		t.Errorf("Point(1, 2, 3).Sub(Vector(4, 5, 6)) = %v; want %v", got, want)
	}
}

func TestNegVector(t *testing.T) {
	vec := Vector(1, 2, 3)
	want := Vector(-1, -2, -3)
	got := Neg(vec)
	if !TupleEqual(want, got) {
		t.Errorf("Neg(Vector(1, 2, 3)) = %v; want %v", got, want)
	}
}

func TestMultVector(t *testing.T) {
	vec := Vector(1, 2, 3)
	scalar := 2.0
	want := Vector(2, 4, 6)
	got := vec.Mult(scalar)
	if !TupleEqual(want, got) {
		t.Errorf("Vector(1, 2, 3).Mult(2) = %v; want %v", got, want)
	}
}

func TestDivideVector(t *testing.T) {
	vec := Vector(1, 2, 3)
	scalar := 2.0
	want := Vector(0.5, 1, 1.5)
	got := vec.Div(scalar)
	if !TupleEqual(want, got) {
		t.Errorf("Vector(1, 2, 3).Divide(2) = %v; want %v", got, want)
	}
}

func TestMag(t *testing.T) {
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
			ans := Mag(tt.t)
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
			ans := Norm(tt.t)
			if !TupleEqual(tt.want, ans) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestMagOfUnitVector(t *testing.T) {
	vec := Vector(1, 2, 3)
	unitVec := Norm(vec)
	want := 1.
	got := Mag(unitVec)
	if !FloatEqual(want, got) {
		t.Errorf("Mag(Normalize(Vector(1, 2, 3))) = %v; want %v", got, want)
	}
}

func TestDotProduct(t *testing.T) {
	v1 := Vector(1, 2, 3)
	v2 := Vector(2, 3, 4)
	want := 20.
	got := Dot(v1, v2)
	if !FloatEqual(want, got) {
		t.Errorf("Mag(Normalize(Vector(1, 2, 3))) = %v; want %v", got, want)
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
			if !TupleEqual(tt.want, ans) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
