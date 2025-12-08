package main

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	if m == 0 || n == 0 {
		return nil
	}

	next := nextCoordinate(m, n) // method val

	traversed := make([]int, m*n)
	for i := range traversed {
		x, y := next()
		traversed[i] = matrix[x][y]
	}
	return traversed
}

func nextCoordinate(m, n int) func() (int, int) {
	// NOTE: higher y values mean going DOWN, higher x is going right
	top, bottom := 0, m-1
	left, right := 0, n-1
	x, y := -1, 0
	dx, dy := 1, 0  // start moving right
	return func() (int, int) {
		x += dx
		y += dy
		switch {
		// change direction. what does that mean?
		// 1. if you hit the right and going down, change where you came (top).
		// if hit bottom and going left, forget right col (right--). etc.
		// 2. change dx, dy appropriately
		case x > right: {
			top++
			dx, dy = 0, 1
		}
		case y > bottom: {
			right--
			dx, dy = -1, 0
		}
		case x > left: {
			bottom--
			dx, dy = 0, -1
		}
		case y > top: {
			left++
			dx, dy = 1, 0
		}
		}
		return x, y
	}
}
