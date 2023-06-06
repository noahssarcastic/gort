package matrix

import (
	"fmt"
	"testing"

	"github.com/noahssarcastic/gort/pkg/tuple"
	"github.com/noahssarcastic/gort/pkg/util"
)

func TestMatrixAccess(t *testing.T) {
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

func TestMatrixGet(t *testing.T) {
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

func TestMatrixEqual(t *testing.T) {
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

func TestMatrixMult(t *testing.T) {
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

func TestMult_identity(t *testing.T) {
	mat := Matrix{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	}
	id := I
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

func TestMatrixApply(t *testing.T) {
	mat := Matrix{
		{1, 2, 3, 4},
		{2, 4, 4, 2},
		{8, 6, 4, 1},
		{0, 0, 0, 1},
	}
	tup := tuple.New(1, 2, 3, 1)
	want := tuple.New(18, 24, 33, 1)
	got := mat.Apply(tup)
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
	mat := I
	want := mat
	got := mat.T()
	if !Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestMatrixDet2(t *testing.T) {
	mat := Matrix{
		{1, 5},
		{-3, 2},
	}
	want := 17.
	got := det2(mat)
	if !util.FloatEqual(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestMatrixSub(t *testing.T) {
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

func TestMatrixMinor(t *testing.T) {
	mat := Matrix{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	}
	want := 25.
	got := mat.minor(1, 0)
	if !util.FloatEqual(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestMatrixCofactor(t *testing.T) {
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
			if !util.FloatEqual(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestMatrixDet(t *testing.T) {
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
			if !util.FloatEqual(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestMatrixIsInvertible(t *testing.T) {
	mat := Matrix{
		{6, 4, 4, 4},
		{5, 5, 7, 6},
		{4, -9, 3, -7},
		{9, 1, 7, -6},
	}
	want := true
	got := mat.IsInvertible()
	if want != got {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestMatrixIsInvertible_not(t *testing.T) {
	mat := Matrix{
		{-4, 2, -2, -3},
		{9, 6, 2, 6},
		{0, -5, 1, -5},
		{0, 0, 0, 0},
	}
	want := false
	got := mat.IsInvertible()
	if want != got {
		t.Errorf("want %v; got %v", want, got)
	}
}

func TestMatrixInv(t *testing.T) {
	tests := []struct {
		mat  Matrix
		want Matrix
	}{
		{
			Matrix{
				{-5, 2, 6, -8},
				{1, -5, 1, 8},
				{7, 7, -6, -7},
				{1, -3, 7, 4},
			},
			Matrix{
				{0.21805, 0.45113, 0.24060, -0.04511},
				{-0.80827, -1.45677, -0.44361, 0.52068},
				{-0.07895, -0.22368, -0.05263, 0.19737},
				{-0.52256, -0.81391, -0.30075, 0.30639},
			},
		},
		{
			Matrix{
				{8, -5, 9, 2},
				{7, 5, 6, 1},
				{-6, 0, 9, 6},
				{-3, 0, -9, -4},
			},
			Matrix{
				{-0.15385, -0.15385, -0.28205, -0.53846},
				{-0.07692, 0.12308, 0.02564, 0.03077},
				{0.35897, 0.35897, 0.43590, 0.92308},
				{-0.69231, -0.69231, -0.76923, -1.92308},
			},
		},
		{
			Matrix{
				{9, 3, 0, 9},
				{-5, -2, -6, -3},
				{-4, 9, 6, 4},
				{-7, 6, 6, 2},
			},
			Matrix{
				{-0.04074, -0.07778, 0.14444, -0.22222},
				{-0.07778, 0.03333, 0.36667, -0.33333},
				{-0.02901, -0.14630, -0.10926, 0.12963},
				{0.17778, 0.06667, -0.26667, 0.33333},
			},
		},
	}
	for i, tt := range tests {
		name := fmt.Sprintf("%v", i)
		t.Run(name, func(t *testing.T) {
			ans := Inv(tt.mat)
			if !Equal(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestMatrixMult_undo(t *testing.T) {
	a := Matrix{
		{8, -5, 9, 2},
		{7, 5, 6, 1},
		{-6, 0, 9, 6},
		{-3, 0, -9, -4},
	}
	b := Matrix{
		{9, 3, 0, 9},
		{-5, -2, -6, -3},
		{-4, 9, 6, 4},
		{-7, 6, 6, 2},
	}
	want := a
	got := Mult(Mult(a, b), Inv(b))
	if !Equal(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}
}
