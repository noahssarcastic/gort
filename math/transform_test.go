package math

import (
	"fmt"
	stdmath "math"
	"testing"
)

func TestTranslate_point(t *testing.T) {
	tform := Translate(5, -3, 2)
	pt := Point(-3, 4, 5)
	want := Point(2, 1, 7)
	got := tform.Apply(pt)
	if !TupleEqual(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestTranslate_inv(t *testing.T) {
	tform := Translate(5, -3, 2)
	pt := Point(-3, 4, 5)
	want := Point(-8, 7, 3)
	got := Inv(tform).Apply(pt)
	if !TupleEqual(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestTranslate_vector(t *testing.T) {
	tform := Translate(5, -3, 2)
	vec := Vector(-3, 4, 5)
	want := vec
	got := tform.Apply(vec)
	if !TupleEqual(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestScale_point(t *testing.T) {
	pt := Point(-4, 6, 8)
	tform := Scale(2, 3, 4)
	want := Point(-8, 18, 32)
	got := tform.Apply(pt)
	if !TupleEqual(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestScale_inv(t *testing.T) {
	vec := Vector(-4, 6, 8)
	tform := Scale(2, 3, 4)
	want := Vector(-2, 2, 2)
	got := Inv(tform).Apply(vec)
	if !TupleEqual(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestScale_vector(t *testing.T) {
	vec := Vector(-4, 6, 8)
	tform := Scale(2, 3, 4)
	want := Vector(-8, 18, 32)
	got := tform.Apply(vec)
	if !TupleEqual(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestScale_reflect(t *testing.T) {
	pt := Point(2, 3, 4)
	tform := Scale(-1, 1, 1)
	want := Point(-2, 3, 4)
	got := tform.Apply(pt)
	if !TupleEqual(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestRotate_x(t *testing.T) {
	tests := []struct {
		pt   Tuple
		rads float64
		want Tuple
	}{
		{Point(0, 1, 0), stdmath.Pi / 4, Point(0, stdmath.Sqrt(2)/2, stdmath.Sqrt(2)/2)},
		{Point(0, 1, 0), stdmath.Pi / 2, Point(0, 0, 1)},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%v,%v", tt.pt, tt.rads)
		t.Run(name, func(t *testing.T) {
			ans := RotateX(tt.rads).Apply(tt.pt)
			if !TupleEqual(ans, tt.want) {
				t.Errorf("got %v; want %v", ans, tt.want)
			}
		})
	}
}

func TestRotate_inv(t *testing.T) {
	pt := Point(0, 1, 0)
	tform := RotateX(stdmath.Pi / 4)
	want := Point(0, stdmath.Sqrt(2)/2, -stdmath.Sqrt(2)/2)
	got := Inv(tform).Apply(pt)
	if !TupleEqual(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestRotate_y(t *testing.T) {
	tests := []struct {
		pt   Tuple
		rads float64
		want Tuple
	}{
		{Point(0, 0, 1), stdmath.Pi / 4, Point(stdmath.Sqrt(2)/2, 0, stdmath.Sqrt(2)/2)},
		{Point(0, 0, 1), stdmath.Pi / 2, Point(1, 0, 0)},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%v,%v", tt.pt, tt.rads)
		t.Run(name, func(t *testing.T) {
			ans := RotateY(tt.rads).Apply(tt.pt)
			if !TupleEqual(ans, tt.want) {
				t.Errorf("got %v; want %v", ans, tt.want)
			}
		})
	}
}

func TestRotate_z(t *testing.T) {
	tests := []struct {
		pt   Tuple
		rads float64
		want Tuple
	}{
		{Point(0, 1, 0), stdmath.Pi / 4, Point(-stdmath.Sqrt(2)/2, stdmath.Sqrt(2)/2, 0)},
		{Point(0, 1, 0), stdmath.Pi / 2, Point(-1, 0, 0)},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%v,%v", tt.pt, tt.rads)
		t.Run(name, func(t *testing.T) {
			ans := RotateZ(tt.rads).Apply(tt.pt)
			if !TupleEqual(ans, tt.want) {
				t.Errorf("got %v; want %v", ans, tt.want)
			}
		})
	}
}

func TestShear_xy(t *testing.T) {
	pt := Point(2, 3, 4)
	tform := Shear(1, 0, 0, 0, 0, 0)
	want := Point(5, 3, 4)
	got := tform.Apply(pt)
	if !TupleEqual(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestShear_xz(t *testing.T) {
	pt := Point(2, 3, 4)
	tform := Shear(0, 1, 0, 0, 0, 0)
	want := Point(6, 3, 4)
	got := tform.Apply(pt)
	if !TupleEqual(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestShear_yx(t *testing.T) {
	pt := Point(2, 3, 4)
	tform := Shear(0, 0, 1, 0, 0, 0)
	want := Point(2, 5, 4)
	got := tform.Apply(pt)
	if !TupleEqual(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestShear_yz(t *testing.T) {
	pt := Point(2, 3, 4)
	tform := Shear(0, 0, 0, 1, 0, 0)
	want := Point(2, 7, 4)
	got := tform.Apply(pt)
	if !TupleEqual(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestShear_zx(t *testing.T) {
	pt := Point(2, 3, 4)
	tform := Shear(0, 0, 0, 0, 1, 0)
	want := Point(2, 3, 6)
	got := tform.Apply(pt)
	if !TupleEqual(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestShear_zy(t *testing.T) {
	pt := Point(2, 3, 4)
	tform := Shear(0, 0, 0, 0, 0, 1)
	want := Point(2, 3, 7)
	got := tform.Apply(pt)
	if !TupleEqual(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestChainTransforms(t *testing.T) {
	pt := Point(1, 0, 1)
	tform := Chain(
		RotateX(stdmath.Pi/2),
		Scale(5, 5, 5),
		Translate(10, 5, 7),
	)
	want := Point(15, 0, 7)
	got := tform.Apply(pt)
	if !TupleEqual(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}
