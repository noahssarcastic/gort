package mat

import (
	"strconv"

	"github.com/noahssarcastic/gort/pkg/tuple"
	"github.com/noahssarcastic/gort/pkg/util"
)

// Matrix keeps a 2D array of float64s representing a square matrix.
type Matrix [][]float64

// New initializes an empty Matrix with dimension dim.
func New(dim int) Matrix {
	mat := make(Matrix, dim)
	for i := range mat {
		mat[i] = make([]float64, dim)
	}
	return mat
}

// String returns a human readable string representation of a Matrix.
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

// Dim gets the dimension of the matrix.
func (mat Matrix) Dim() int {
	return len(mat[0])
}

// Get accesses the value at (row, col).
func (mat Matrix) Get(row, col int) float64 {
	return mat[row][col]
}

// Set changes the value at (row, col) to val.
func (mat Matrix) Set(row, col int, val float64) {
	mat[row][col] = val
}

// Equal returns true if two Matrices are equal.
func Equal(a, b Matrix) bool {
	if a.Dim() != b.Dim() {
		return false
	}
	for y := 0; y < a.Dim(); y++ {
		for x := 0; x < a.Dim(); x++ {
			if !util.FloatEqual(a.Get(y, x), b.Get(y, x)) {
				return false
			}
		}
	}
	return true
}

// Mult takes two 4x4 Matrices and returns the product.
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

// Apply returns the product of a 4x4 Matrix and a tuple.Tuple.
func (mat Matrix) Apply(t tuple.Tuple) tuple.Tuple {
	return tuple.New(
		mat[0][0]*t.X()+mat[0][1]*t.Y()+mat[0][2]*t.Z()+mat[0][3]*t.W(),
		mat[1][0]*t.X()+mat[1][1]*t.Y()+mat[1][2]*t.Z()+mat[1][3]*t.W(),
		mat[2][0]*t.X()+mat[2][1]*t.Y()+mat[2][2]*t.Z()+mat[2][3]*t.W(),
		mat[3][0]*t.X()+mat[3][1]*t.Y()+mat[3][2]*t.Z()+mat[3][3]*t.W(),
	)
}

// I initializes and returns the 4x4 identity matrix.
func I() Matrix {
	dim := 4
	mat := New(dim)
	for i := 0; i < dim; i++ {
		mat.Set(i, i, 1)
	}
	return mat
}

// T initializes a new Matrix that is the transpose of mat.
func (mat Matrix) T() Matrix {
	trans := New(mat.Dim())
	for y, row := range mat {
		for x, el := range row {
			trans.Set(x, y, el)
		}
	}
	return trans
}

// det2 calculates the determinant of a 2x2 matrix.
func det2(mat Matrix) float64 {
	return mat[0][0]*mat[1][1] - mat[0][1]*mat[1][0]
}

// sub calculates the submatrix of a Matrix mat. The submatrix of a Matrix mat
// is a Matrix in which the row r and column c have been removed from mat.
func (mat Matrix) sub(r, c int) Matrix {
	subMat := New(mat.Dim() - 1)
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

// minor calculates the minor of a 3x3 matrix at (r,c).
func (mat Matrix) minor(r, c int) float64 {
	return Det(mat.sub(r, c))
}

// cofactor calculates the cofactor of a 3x3 matrix at (r,c).
func (mat Matrix) cofactor(r, c int) float64 {
	minor := mat.minor(r, c)
	if (r+c)%2 != 0 {
		minor *= -1
	}
	return minor
}

// Det calculates the determinant of a 4x4 matrix.
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

// IsInvertible returns true if a Matrix mat is invertible.
func (mat Matrix) IsInvertible() bool {
	return !util.FloatEqual(Det(mat), 0)
}

// Inv initializes a new Matrix which is the inverse of a Matrix mat.
func Inv(mat Matrix) Matrix {
	if !mat.IsInvertible() {
		panic("matrix is not invertible")
	}
	det := Det(mat)
	invMat := New(mat.Dim())
	for r, row := range mat {
		for c := range row {
			// (col,row) order accomplishes a transpose
			invMat[c][r] = mat.cofactor(r, c) / det
		}
	}
	return invMat
}
