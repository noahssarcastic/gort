package matrix

import (
	"fmt"
	"strconv"

	"github.com/noahssarcastic/tddraytracer/tuple"
	"github.com/noahssarcastic/tddraytracer/utils"
)

type Matrix [][]float64

func New(w, h int) Matrix {
	mat := make(Matrix, h)
	for i := range mat {
		mat[i] = make([]float64, w)
	}
	return mat
}

func FromTuple(t tuple.Tuple) Matrix {
	return Matrix{
		{t.X()},
		{t.Y()},
		{t.Z()},
		{t.W()},
	}
}

func (mat Matrix) String() string {
	s := make([]byte, 0)
	s = append(s, '[')
	for y, row := range mat {
		if y != 0 {
			s = append(s, ' ')
		}
		s = append(s, '[')
		for x, el := range row {
			if x != 0 {
				s = append(s, ' ')
			}
			s = append(s, []byte(strconv.FormatFloat(el, 'f', -1, 64))...)
		}
		s = append(s, ']')
		if y < len(mat)-1 {
			s = append(s, '\n')
		}
	}
	s = append(s, ']')
	return string(s)
}

func (mat Matrix) Width() int {
	return len(mat[0])
}

func (mat Matrix) Height() int {
	return len(mat)
}

func (mat Matrix) Get(r, c int) float64 {
	return mat[r][c]
}

func (mat Matrix) Set(r, c int, val float64) {
	mat[r][c] = val
}

func (mat Matrix) IsMatrix() bool {
	rows := len(mat)
	cols := len(mat[0])
	for y := 1; y < rows; y++ {
		if len(mat[y]) != cols {
			return false
		}
	}
	return true
}

func Equal(a, b Matrix) bool {
	if !(a.Width() == b.Width()) || !(a.Height() == b.Height()) {
		return false
	}

	for y := 0; y < a.Height(); y++ {
		for x := 0; x < a.Width(); x++ {
			if !utils.FloatEqual(a.Get(y, x), b.Get(y, x)) {
				return false
			}
		}
	}
	return true
}

func Multiply(a, b Matrix) Matrix {
	if a.Width() != b.Height() {
		panic(fmt.Sprintf("cannot multiply %v and %v; invalid dimensions", a, b))
	}

	w := b.Width()
	h := a.Height()
	newMat := New(w, h)
	for y, row := range newMat {
		for x := range row {
			for i := 0; i < a.Width(); i++ {
				row[x] += a.Get(y, i) * b.Get(i, x)
			}
		}
	}
	return newMat
}

func (mat Matrix) Multiply(t tuple.Tuple) tuple.Tuple {
	colVector := Matrix{
		{t.X()},
		{t.Y()},
		{t.Z()},
		{t.W()},
	}
	ret := Multiply(mat, colVector)
	return tuple.New(
		ret.Get(0, 0),
		ret.Get(1, 0),
		ret.Get(2, 0),
		ret.Get(3, 0),
	)
}

func I(dim int) Matrix {
	mat := New(dim, dim)
	for i := 0; i < dim; i++ {
		mat.Set(i, i, 1)
	}
	return mat
}

func (mat Matrix) T() Matrix {
	trans := New(mat.Height(), mat.Width())
	for y, row := range mat {
		for x, el := range row {
			trans.Set(x, y, el)
		}
	}
	return trans
}
