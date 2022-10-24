package main

import (
	"fmt"

	"github.com/noahssarcastic/tddraytracer/matrix"
)

func main() {
	mat := matrix.Matrix{
		[]float64{1, 2, 3},
		[]float64{4, 5, 6},
	}
	fmt.Println(mat)
}
