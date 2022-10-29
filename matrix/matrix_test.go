package matrix

import (
	"fmt"
	"testing"

	"github.com/noahssarcastic/tddraytracer/tuple"
	"github.com/noahssarcastic/tddraytracer/utils"
)

func TestAccess(t *testing.T) {
	mat := Matrix{
		{1, 2, 3},
		{4, 5, 6},
		{4, 5, 6},
	}
	want := 4.
	got := mat[1][0]
	if want != got {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestGet(t *testing.T) {
	mat := Matrix{
		{1, 2, 3},
		{4, 5, 6},
		{4, 5, 6},
	}
	want := 2.
	got := mat.Get(0, 1)
	if want != got {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestEqual(t *testing.T) {
	var tests = []struct {
		a, b Matrix
		want bool
	}{
		{
			Matrix{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			Matrix{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			true,
		},
		{
			Matrix{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			Matrix{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
			false,
		},
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

func TestMultiply_square(t *testing.T) {
	a := Matrix{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	}
	b := Matrix{
		{-2, 1, 2, 3},
		{3, 2, 1, -1},
		{4, 3, 6, 5},
		{1, 2, 7, 8},
	}
	want := Matrix{
		{20, 22, 50, 48},
		{44, 54, 114, 108},
		{40, 58, 110, 102},
		{16, 26, 46, 42},
	}
	got := Mult(a, b)
	if !Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestMultiply_identity(t *testing.T) {
	mat := Matrix{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	}
	id := I(4)
	want := Matrix{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	}
	got := Mult(mat, id)
	if !Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestMultiply_tuple(t *testing.T) {
	mat := Matrix{
		{1, 2, 3, 4},
		{2, 4, 4, 2},
		{8, 6, 4, 1},
		{0, 0, 0, 1},
	}
	tup := tuple.New(1, 2, 3, 1)
	want := tuple.New(18, 24, 33, 1)
	got := mat.MultTuple(tup)
	if !tuple.Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestMatrixT(t *testing.T) {
	mat := Matrix{
		{1, 2, 3},
		{4, 5, 6},
		{4, 5, 6},
	}
	want := Matrix{
		{1, 4, 4},
		{2, 5, 5},
		{3, 6, 6},
	}
	got := mat.T()
	if !Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestMatrixT_identity(t *testing.T) {
	mat := I(4)
	want := mat
	got := mat.T()
	if !Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestDet2(t *testing.T) {
	mat := Matrix{
		{1, 5},
		{-3, 2},
	}
	want := 17.
	got := det2(mat)
	if !utils.FloatEqual(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestSub(t *testing.T) {
	mat3 := Matrix{
		{1, 5, 0},
		{-3, 2, 7},
		{0, 6, -3},
	}
	var tests = []struct {
		mat  Matrix
		r, c int
		want Matrix
	}{
		{
			mat3, 0, 0, Matrix{
				{2, 7},
				{6, -3},
			},
		},
		{
			mat3, 1, 1, Matrix{
				{1, 0},
				{0, -3},
			},
		},
		{
			mat3, 2, 2, Matrix{
				{1, 5},
				{-3, 2},
			},
		},
		{
			Matrix{
				{-6, 1, 1, 6},
				{-8, 5, 8, 6},
				{-1, 0, 8, 2},
				{-7, 1, -1, 1},
			},
			2, 1,
			Matrix{
				{-6, 1, 6},
				{-8, 8, 6},
				{-7, -1, 1},
			},
		},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("dim%v,(%v,%v)", tt.mat.Dim(), tt.r, tt.c)
		t.Run(name, func(t *testing.T) {
			ans := tt.mat.sub(tt.r, tt.c)
			if !Equal(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestMinor(t *testing.T) {
	mat := Matrix{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	}
	want := 25.
	got := mat.minor(1, 0)
	if !utils.FloatEqual(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestCofactor(t *testing.T) {
	mat3 := Matrix{
		{1, 2, 6},
		{-5, 8, -4},
		{2, 6, 4},
	}
	mat4 := Matrix{
		{-2, -8, 3, 5},
		{-3, 1, 7, 3},
		{1, 2, -9, 6},
		{-6, 7, 7, -9},
	}
	tests := []struct {
		mat  Matrix
		r, c int
		want float64
	}{
		{mat3, 0, 0, 56},
		{mat3, 0, 1, 12},
		{mat3, 0, 2, -46},
		{mat4, 0, 0, 690},
		{mat4, 0, 1, 447},
		{mat4, 0, 2, 210},
		{mat4, 0, 3, 51},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("dim%v,(%v,%v)", tt.mat.Dim(), tt.r, tt.c)
		t.Run(name, func(t *testing.T) {
			ans := tt.mat.cofactor(tt.r, tt.c)
			if !utils.FloatEqual(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestDet(t *testing.T) {
	tests := []struct {
		mat  Matrix
		want float64
	}{
		{
			Matrix{
				{1, 2, 6},
				{-5, 8, -4},
				{2, 6, 4},
			},
			-196,
		},
		{
			Matrix{
				{-2, -8, 3, 5},
				{-3, 1, 7, 3},
				{1, 2, -9, 6},
				{-6, 7, 7, -9},
			},
			-4071,
		},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("dim%v", tt.mat.Dim())
		t.Run(name, func(t *testing.T) {
			ans := Det(tt.mat)
			if !utils.FloatEqual(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
