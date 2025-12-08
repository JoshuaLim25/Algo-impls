package main

import (
	"fmt"
	"os"
)

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	if m == 0 || n == 0 {
		return nil
	}
	traversed := make([]int, m*n)
	// next := nextCoordinateBAD(m, n)
	next := nextCoordinate(m, n)
	for i := range traversed {
		x, y := next()
		traversed[i] = matrix[x][y]
	}
	return traversed
}

func nextCoordinateBAD(m, n int) func() (int, int) {
	x, y := -1, 0
	dx, dy := 1, 0
	// NOTE: top++ means going down, bottom-- means going up
	top, bottom := 0, m-1
	left, right := 0, n-1
	return func() (int, int) {
		x+=dx
		y+=dy
		switch {
		// if we went one further, what would happen if we hit a wall?
		case x+dx > right: {
			top++
			dx, dy = 0, 1
		}
		case y+dy > bottom: {
			right--
			dx, dy = -1, 0
		}
		case x+dx < left: {
			bottom--
			dx, dy = 0, -1
		}
		case y+dy < top: {
			left++
			dx, dy = 1, 0
		}
		}
		return x, y
	}
}

func nextCoordinate(m, n int) func() (int, int) {
	top, bottom := 0, m-1
	left, right := 0, n-1
	// y++ means the rightward direction, x++ means downward
	x, y := 0, -1
	dx, dy := 0, 1
	return func() (int, int) {
		x += dx
		y += dy
		switch {
		case y+dy > right: {
			top++
			dx, dy = 1, 0
		}
		case x+dx > bottom: {
			right--
			dx, dy = 0, -1
		}
		case y+dy < left: {
			bottom--
			dx, dy = -1, 0
		}
		case x+dx < top: {
			left++
			dx, dy = 0, 1
		}
		}
		return x, y
	}
}

func main() {
	mat1 := [][]int{{1,2,3},{4,5,6},{7,8,9}} //  output: [1,2,3,6,9,8,7,4,5]
	res1 := spiralOrder(mat1)
	fmt.Fprintf(os.Stderr, "DEBUG[1]: 54_spiral_matrix.go:51: res1=%+v\n", res1)
}
