/*
79. Word Search - Medium

Given an m x n grid of characters board and a string word, return true if word exists in the grid.

The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.


Example 1:

Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
Output: true
*/

package com.algos.backtracking;

public class WordSearch79 {
    char[][] board;
    boolean[][] visited;
    String word;
    int curChar = 0;

    public boolean exist(char[][] board, String word) {
        int m = board.length, n = board[0].length;
        this.board = board;
        this.word = word;
        this.visited = new boolean[m][n];

        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < n; ++j) {
                if (search(i, j, curChar))
                    return true;
            }
        }

        return false;
    }

    private boolean search(int i, int j, int curChar) {
        int m = board.length, n = board[0].length;
        if (curChar == word.length())
            return true;
        if (i < 0 || i >= m || j < 0 || j >= n || visited[i][j] || board[i][j] != word.charAt(curChar))
            return false;

        visited[i][j] = true;

        if (search(i + 1, j, curChar + 1) ||
                search(i - 1, j, curChar + 1) ||
                search(i, j + 1, curChar + 1) ||
                search(i, j - 1, curChar + 1)) {
            return true;
        }

        visited[i][j] = false;
        return false;
    }
}


/*
exist()
└─ search(0,0,0)
   └─ search(0,1,1)
      └─ search(0,2,2)
         └─ search(1,2,3)
            └─ search(2,2,4)
               └─ search(2,1,5)
                  └─ search(...,6) → returns TRUE
               → returns TRUE
            → returns TRUE
         → returns TRUE
      → returns TRUE
   → returns TRUE
→ returns TRUE
*/
