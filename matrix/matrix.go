package matrix

import (
	"strconv"

	"github.com/noahssarcastic/tddraytracer/tuple"
	"github.com/noahssarcastic/tddraytracer/utils"
)

// Represents a square matrix
type Matrix [][]float64

func New(dim int) Matrix {
	mat := make(Matrix, dim)
	for i := range mat {
		mat[i] = make([]float64, dim)
	}
	return mat
}

func (mat Matrix) String() string {
	s := []byte{'['}
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
	return string(append(s, ']'))
}

// Get the dimension of the matrix
func (mat Matrix) Dim() int {
	return len(mat[0])
}

// Get the value at (row, column)
func (mat Matrix) Get(r, c int) float64 {
	return mat[r][c]
}

// Set the value at (row, column)
func (mat Matrix) Set(r, c int, val float64) {
	mat[r][c] = val
}

// Check if two matrices are equal
func Equal(a, b Matrix) bool {
	if a.Dim() != b.Dim() {
		return false
	}
	for y := 0; y < a.Dim(); y++ {
		for x := 0; x < a.Dim(); x++ {
			if !utils.FloatEqual(a.Get(y, x), b.Get(y, x)) {
				return false
			}
		}
	}
	return true
}

// Multiply two 4x4 matrices
func Mult(a, b Matrix) Matrix {
	newMat := New(4)
	for y, row := range newMat {
		for x := range row {
			row[x] = a.Get(y, 0)*b.Get(0, x) +
				a.Get(y, 1)*b.Get(1, x) +
				a.Get(y, 2)*b.Get(2, x) +
				a.Get(y, 3)*b.Get(3, x)
		}
	}
	return newMat
}

// Multiply a matrix by a tuple
func (mat Matrix) MultTuple(t tuple.Tuple) tuple.Tuple {
	return tuple.New(
		mat[0][0]*t.X()+mat[0][1]*t.Y()+mat[0][2]*t.Z()+mat[0][3]*t.W(),
		mat[1][0]*t.X()+mat[1][1]*t.Y()+mat[1][2]*t.Z()+mat[1][3]*t.W(),
		mat[2][0]*t.X()+mat[2][1]*t.Y()+mat[2][2]*t.Z()+mat[2][3]*t.W(),
		mat[3][0]*t.X()+mat[3][1]*t.Y()+mat[3][2]*t.Z()+mat[3][3]*t.W(),
	)
}

// Get the identity matrix
func I(dim int) Matrix {
	mat := New(dim)
	for i := 0; i < dim; i++ {
		mat.Set(i, i, 1)
	}
	return mat
}

// Get the transpose of the matrix
func (mat Matrix) T() Matrix {
	trans := New(mat.Dim())
	for y, row := range mat {
		for x, el := range row {
			trans.Set(x, y, el)
		}
	}
	return trans
}

// // Get the determinant of a 2x2 matrix
// func det2(mat Matrix) float64 {

// }

// // Get the determinant of a 3x3 matrix
// func det3(mat Matrix) float64 {

// }

// // Get the determinant of a 4x4 matrix
// func (mat Matrix) Det() float64 {

// }
