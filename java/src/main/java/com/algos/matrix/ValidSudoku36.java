/*
36. Valid Sudoku - Medium

- Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:
    - Each row must contain the digits 1-9 without repetition.
    - Each column must contain the digits 1-9 without repetition.
    - Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.

Note:
    - A Sudoku board (partially filled) could be valid but is not necessarily solvable.
    - Only the filled cells need to be validated according to the mentioned rules.

Example 1:

Input: board =
[["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: true

Example 2:

Input: board =
[["8","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: false
Explanation: Same as Example 1, except with the 5 in the top left corner being modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.



Constraints:
    board.length == 9
    board[i].length == 9
    board[i][j] is a digit 1-9 or '.'.
*/

package com.algos.matrix;

import java.util.HashSet;
import java.util.Set;

public class ValidSudoku36 {
    public static boolean isValidSudoku(char[][] board) {
        int m = board.length, n = board[0].length;
        boolean[][] rows = new boolean[m][n];
        boolean[][] cols = new boolean[m][n];
        boolean[][] boxes = new boolean[m][n];
        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < n; ++j) {
                if (board[i][j] == '.')
                    continue;
                // int num = Character.getNumericValue(board[i][j]);
                // --num;
                int num = board[i][j] - '1'; // so fucking dumb
                // [ 0 1 2 ]
                // [ 3 4 5 ]
                // [ 6 7 8 ]
                // i = 012, j = 012. To get the actual index, multiply i by number of rows in a
                // 3x3 matrix (ie 3)
                int k = ((i / 3) * 3 + (j / 3));
                // boolean[][] functions as a set of seen numbers
                if (rows[i][num] || cols[j][num] || boxes[k][num]) {
                    return false;
                }
                rows[i][num] = cols[j][num] = boxes[k][num] = true;
            }
        }
        return true;
    }
}
