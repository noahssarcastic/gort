package matrix

import (
	"strconv"

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
	return mat[c][r]
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
