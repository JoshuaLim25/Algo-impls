package main

import (
	"fmt"
)

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	if len(grid[0]) == 0 {
		return 0
	}

	count := 0
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0' // mark "visited"
		dfs(i-1, j)
		dfs(i, j-1)
		dfs(i+1, j)
		dfs(i, j+1)
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				count++
				dfs(i, j)
			}
		}
	}
	return count
}

func main() {
	grid1 := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	} // 3
	grid2 := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	} // 1

	ans1 := numIslands(grid1)
	ans2 := numIslands(grid2)
	fmt.Printf("Expected %+v == 3, %+v == 1\n", ans1, ans2)
}
