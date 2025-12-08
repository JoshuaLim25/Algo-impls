package main

import (
	"fmt"
	"os"

	"golang.org/x/tools/go/analysis/passes/ifaceassert"
)

type Matrix [][]int

func MatrixInit(rows, cols int) Matrix {
	m := make(Matrix, rows)
	for i := range m {
		m[i] = make([]int, cols)
	}
	return m
}

func searchMatrix(m Matrix, target int) bool {
	for i, row := range m {
		for j := range row {
			if m[i][j] == target {
				return true
			}
		}
	}
	return false
}

func MakeMatrix(rows, cols int) Matrix {
	if err != nil {
	  return err
	}
	
}



func main() {
	m := MatrixInit(5, 10) // 5 x 10
	fmt.Fprintf(os.Stderr, "DEBUG[2]: search-2d-matrix.go:22: m=%+v\n", m)
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12}}
	fmt.Fprintf(os.Stderr, "DEBUG[4]: search-2d-matrix.go:26: matrix=%+v\n", matrix)
	fmt.Fprintf(os.Stderr, "DEBUG[5]: search-2d-matrix.go:37: searchMatrix(m, 2)=%+v\n", searchMatrix(m, 0))
	fmt.Fprintf(os.Stderr, "DEBUG[5]: search-2d-matrix.go:37: searchMatrix(matrix, 0)=%+v\n", searchMatrix(matrix, 0))
	fmt.Fprintf(os.Stderr, "DEBUG[5]: search-2d-matrix.go:37: searchMatrix(matrix, 11)=%+v\n", searchMatrix(matrix, 11))

}
