package transform

import (
	"fmt"
	"math"
	"testing"

	"github.com/noahssarcastic/tddraytracer/matrix"
	"github.com/noahssarcastic/tddraytracer/tuple"
)

func TestTranslate_point(t *testing.T) {
	tform := Translate(5, -3, 2)
	pt := tuple.Point(-3, 4, 5)
	want := tuple.Point(2, 1, 7)
	got := tform.MultTuple(pt)
	if !tuple.Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestTranslate_inv(t *testing.T) {
	tform := Translate(5, -3, 2)
	pt := tuple.Point(-3, 4, 5)
	want := tuple.Point(-8, 7, 3)
	got := matrix.Inv(tform).MultTuple(pt)
	if !tuple.Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestTranslate_vector(t *testing.T) {
	tform := Translate(5, -3, 2)
	vec := tuple.Vector(-3, 4, 5)
	want := vec
	got := tform.MultTuple(vec)
	if !tuple.Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestScale_point(t *testing.T) {
	pt := tuple.Point(-4, 6, 8)
	tform := Scale(2, 3, 4)
	want := tuple.Point(-8, 18, 32)
	got := tform.MultTuple(pt)
	if !tuple.Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestScale_inv(t *testing.T) {
	vec := tuple.Vector(-4, 6, 8)
	tform := Scale(2, 3, 4)
	want := tuple.Vector(-2, 2, 2)
	got := matrix.Inv(tform).MultTuple(vec)
	if !tuple.Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestScale_vector(t *testing.T) {
	vec := tuple.Vector(-4, 6, 8)
	tform := Scale(2, 3, 4)
	want := tuple.Vector(-8, 18, 32)
	got := tform.MultTuple(vec)
	if !tuple.Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestScale_reflect(t *testing.T) {
	pt := tuple.Point(2, 3, 4)
	tform := Scale(-1, 1, 1)
	want := tuple.Point(-2, 3, 4)
	got := tform.MultTuple(pt)
	if !tuple.Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestRotate_x(t *testing.T) {
	tests := []struct {
		pt   tuple.Tuple
		rads float64
		want tuple.Tuple
	}{
		{tuple.Point(0, 1, 0), math.Pi / 4, tuple.Point(0, math.Sqrt(2)/2, math.Sqrt(2)/2)},
		{tuple.Point(0, 1, 0), math.Pi / 2, tuple.Point(0, 0, 1)},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%v,%v", tt.pt, tt.rads)
		t.Run(name, func(t *testing.T) {
			ans := rotateX(tt.rads).MultTuple(tt.pt)
			if !tuple.Equal(ans, tt.want) {
				t.Errorf("got %v; want %v", ans, tt.want)
			}
		})
	}
}

func TestRotate_inv(t *testing.T) {
	pt := tuple.Point(0, 1, 0)
	tform := rotateX(math.Pi / 4)
	want := tuple.Point(0, math.Sqrt(2)/2, -math.Sqrt(2)/2)
	got := matrix.Inv(tform).MultTuple(pt)
	if !tuple.Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestRotate_y(t *testing.T) {
	tests := []struct {
		pt   tuple.Tuple
		rads float64
		want tuple.Tuple
	}{
		{tuple.Point(0, 0, 1), math.Pi / 4, tuple.Point(math.Sqrt(2)/2, 0, math.Sqrt(2)/2)},
		{tuple.Point(0, 0, 1), math.Pi / 2, tuple.Point(1, 0, 0)},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%v,%v", tt.pt, tt.rads)
		t.Run(name, func(t *testing.T) {
			ans := rotateY(tt.rads).MultTuple(tt.pt)
			if !tuple.Equal(ans, tt.want) {
				t.Errorf("got %v; want %v", ans, tt.want)
			}
		})
	}
}

func TestRotate_z(t *testing.T) {
	tests := []struct {
		pt   tuple.Tuple
		rads float64
		want tuple.Tuple
	}{
		{tuple.Point(0, 1, 0), math.Pi / 4, tuple.Point(-math.Sqrt(2)/2, math.Sqrt(2)/2, 0)},
		{tuple.Point(0, 1, 0), math.Pi / 2, tuple.Point(-1, 0, 0)},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%v,%v", tt.pt, tt.rads)
		t.Run(name, func(t *testing.T) {
			ans := rotateZ(tt.rads).MultTuple(tt.pt)
			if !tuple.Equal(ans, tt.want) {
				t.Errorf("got %v; want %v", ans, tt.want)
			}
		})
	}
}

func TestShear_xy(t *testing.T) {
	pt := tuple.Point(2, 3, 4)
	tform := Shear(1, 0, 0, 0, 0, 0)
	want := tuple.Point(5, 3, 4)
	got := tform.MultTuple(pt)
	if !tuple.Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestShear_xz(t *testing.T) {
	pt := tuple.Point(2, 3, 4)
	tform := Shear(0, 1, 0, 0, 0, 0)
	want := tuple.Point(6, 3, 4)
	got := tform.MultTuple(pt)
	if !tuple.Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestShear_yx(t *testing.T) {
	pt := tuple.Point(2, 3, 4)
	tform := Shear(0, 0, 1, 0, 0, 0)
	want := tuple.Point(2, 5, 4)
	got := tform.MultTuple(pt)
	if !tuple.Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestShear_yz(t *testing.T) {
	pt := tuple.Point(2, 3, 4)
	tform := Shear(0, 0, 0, 1, 0, 0)
	want := tuple.Point(2, 7, 4)
	got := tform.MultTuple(pt)
	if !tuple.Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestShear_zx(t *testing.T) {
	pt := tuple.Point(2, 3, 4)
	tform := Shear(0, 0, 0, 0, 1, 0)
	want := tuple.Point(2, 3, 6)
	got := tform.MultTuple(pt)
	if !tuple.Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestShear_zy(t *testing.T) {
	pt := tuple.Point(2, 3, 4)
	tform := Shear(0, 0, 0, 0, 0, 1)
	want := tuple.Point(2, 3, 7)
	got := tform.MultTuple(pt)
	if !tuple.Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestChainTransforms(t *testing.T) {
	// panic("not implemented")
}
