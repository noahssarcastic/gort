package math

import (
	"strconv"
)

// Represents a square matrix
type Matrix [][]float64

func NewMatrix(dim int) Matrix {
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
func MatrixEqual(a, b Matrix) bool {
	if a.Dim() != b.Dim() {
		return false
	}
	for y := 0; y < a.Dim(); y++ {
		for x := 0; x < a.Dim(); x++ {
			if !FloatEqual(a.Get(y, x), b.Get(y, x)) {
				return false
			}
		}
	}
	return true
}

// Multiply two 4x4 matrices
func Mult(a, b Matrix) Matrix {
	newMat := NewMatrix(4)
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
func (mat Matrix) Apply(t Tuple) Tuple {
	return NewTuple(
		mat[0][0]*t.X()+mat[0][1]*t.Y()+mat[0][2]*t.Z()+mat[0][3]*t.W(),
		mat[1][0]*t.X()+mat[1][1]*t.Y()+mat[1][2]*t.Z()+mat[1][3]*t.W(),
		mat[2][0]*t.X()+mat[2][1]*t.Y()+mat[2][2]*t.Z()+mat[2][3]*t.W(),
		mat[3][0]*t.X()+mat[3][1]*t.Y()+mat[3][2]*t.Z()+mat[3][3]*t.W(),
	)
}

// Get the identity matrix
func I(dim int) Matrix {
	mat := NewMatrix(dim)
	for i := 0; i < dim; i++ {
		mat.Set(i, i, 1)
	}
	return mat
}

// Get the transpose of the matrix
func (mat Matrix) T() Matrix {
	trans := NewMatrix(mat.Dim())
	for y, row := range mat {
		for x, el := range row {
			trans.Set(x, y, el)
		}
	}
	return trans
}

// Get the determinant of a 2x2 matrix
func det2(mat Matrix) float64 {
	return mat[0][0]*mat[1][1] - mat[0][1]*mat[1][0]
}

// Get the submatrix of a matrix
func (mat Matrix) sub(r, c int) Matrix {
	subMat := NewMatrix(mat.Dim() - 1)
	for y, row := range mat {
		if y < r {
			copy(subMat[y][:c], row[:c])
			copy(subMat[y][c:], row[c+1:])
		} else if y == r {
			continue
		} else {
			copy(subMat[y-1], row)
			copy(subMat[y-1][c:], row[c+1:])
		}
	}
	return subMat
}

// Get the minor (determinant of sub matrix at (row,column) for a 3x3 matrix)
func (mat Matrix) minor(r, c int) float64 {
	return Det(mat.sub(r, c))
}

func (mat Matrix) cofactor(r, c int) float64 {
	minor := mat.minor(r, c)
	if (r+c)%2 != 0 {
		minor *= -1
	}
	return minor
}

// Get the determinant of a 4x4 matrix
func Det(mat Matrix) (det float64) {
	if mat.Dim() == 2 {
		return det2(mat)
	}

	det = 0
	for i := 0; i < mat.Dim(); i++ {
		det += mat.Get(0, i) * mat.cofactor(0, i)
	}
	return det
}

// Check if matrix is invertible
func (mat Matrix) IsInvertible() bool {
	return !FloatEqual(Det(mat), 0)
}

// Get the inverse of a matrix
func Inv(mat Matrix) Matrix {
	if !mat.IsInvertible() {
		panic("matrix is not invertible")
	}
	det := Det(mat)
	invMat := NewMatrix(mat.Dim())
	for r, row := range mat {
		for c := range row {
			// note col,row order below accomplishes a transpose
			invMat[c][r] = mat.cofactor(r, c) / det
		}
	}
	return invMat
}
