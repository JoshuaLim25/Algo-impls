/*
695. Max Area of Island - Medium

You are given an m x n binary matrix grid. An island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid are surrounded by water.

The area of an island is the number of cells with a value 1 in the island.

Return the maximum area of an island in grid. If there is no island, return 0.

Example 1:
Input: grid = [[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]
Output: 6
Explanation: The answer is not 11, because the island must be connected 4-directionally.

Example 2:

Input: grid = [[0,0,0,0,0,0,0,0]]
Output: 0



Constraints:

    m == grid.length
    n == grid[i].length
    1 <= m, n <= 50
    grid[i][j] is either 0 or 1.

 */

package com.algos.matrix;


public class MaxAreaOfIsland695 {
    int[][] grid;
    boolean[][] visited;
    public int maxArea(int[][] grid) {
        int m = grid.length, n = grid[0].length;
        if (m == 0 || n == 0 || grid == null) return 0;
        visited = new boolean[m][n];
        int area = 0;
        this.grid = grid;
        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < n; ++j) {
                if (grid[i][j] == 0 || visited[i][j]) continue;
                area = Math.max(area, search(i, j));
            }
        }
        return area;
    }

    private int search(int i, int j) {
        int m = grid.length, n = grid[0].length;
        if (i < 0 || i >= m || j < 0 || j >= n || visited[i][j] || grid[i][j] == 0) return 0;
        visited[i][j] = true;

        return 1 + search(i+1, j) + search(i-1, j) + search(i, j+1) + search(i, j-1);
    }
}



