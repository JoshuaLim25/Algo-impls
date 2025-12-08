/*
 *
200. Number of Islands - Medium

Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

Example 1:

Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1
 */


public class NumberOfIslands200 {
    public static int number(char[][] grid) {
        // go through O(n^2) row by col, try to match against '1'
        int m = grid.length, n = grid[0].length;
        int count = 0;
        if (m == 0 || n == 0) return count;


        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < n; ++j) {
                if (grid[i][j] == '1') {
                    count++;
                    search(grid, i, j);
                }
            }
        }

        return count;

    }

    private static void search(char[][] grid, int i, int j) {
        if (i < 0 || i >= grid.length || j < 0 || j >= grid[0].length || grid[i][j] == '0') {
            return;
        }
        grid[i][j] = '0';
        search(grid, i+1, j);
        search(grid, i-1, j);
        search(grid, i, j+1);
        search(grid, i, j-1);
    }

    public static void main(String[] args) {
        char[][] grid = new char[][]{
               {'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}
        };
        System.out.println(String.format("Should be 1: %s", number(grid)));
    }

}
