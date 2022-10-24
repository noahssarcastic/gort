package matrix

import "fmt"

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
		if y == 0 {
			s = append(s, '[')
		} else {
			s = append(s, " ["...)
		}

		for x, el := range row {
			if x != 0 {
				s = append(s, ' ')
			}
			s = append(s, []byte(fmt.Sprintf("%f", el))...)
		}

		if y < len(mat)-1 {
			s = append(s, "]\n"...)
		} else {
			s = append(s, "]"...)
		}
	}
	s = append(s, "]"...)
	return string(s)
}

func (mat Matrix) Width() int {
	return len(mat[0])
}

func (mat Matrix) Height() int {
	return len(mat)
}

func (mat Matrix) Get(x, y int) float64 {
	return mat[x][y]
}
