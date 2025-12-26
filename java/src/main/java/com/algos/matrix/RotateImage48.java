/*
48. Rotate Image - Medium

You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).

You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.

Example 1:
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [[7,4,1],[8,5,2],[9,6,3]]

Example 2:
Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]

Constraints:
    n == matrix.length == matrix[i].length
    1 <= n <= 20
    -1000 <= matrix[i][j] <= 1000
*/

/*
// [1 2 3]
// [4 5 6]
// [7 8 9]

// [7 4 1]
// [6 5 2]
// [9 8 3]
 */

package com.algos.matrix;

public class RotateImage48 {
    public static void rotate(int[][] matrix) {
        // just transpose
        int m = matrix.length, n = matrix[0].length;
        if (m == 0 || matrix == null) return;
        for (int i = 0; i < m; ++i) {
            // NOTE: NOT j=0. if you do that, it swaps twice and undoes the operation
            // NOTE: by setting j = i, the inner loop only processes elements on or above the main diagonal (the line from the top-left to bottom-right).
            // envision an upper triangle. j=i makes sure you stay in that region,
            // and so you never even touch the lower one, and the swap happens once only.
            for (int j = i; j < n; ++j) {
                int tmp = matrix[i][j];

                matrix[i][j] = matrix[j][i];
                matrix[j][i] = tmp;
            }
        }

        for (int i = 0; i < m; ++i) {  // foreach row
            for (int j = 0; j < n/2; ++j) { // up until row's halfway pt (#cols/2)
                int tmp = matrix[i][n-1-j];  // mat @ ith row and the corr. end of that row
                matrix[i][n-1-j] = matrix[i][j];
                matrix[i][j] = tmp;
            }
        }
    }
}

