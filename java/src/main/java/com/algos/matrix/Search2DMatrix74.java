/* 
74. Search a 2D matrix - Medium

You are given an m x n integer matrix matrix with the following two properties:
- Each row is sorted in non-decreasing order.
- The first integer of each row is greater than the last integer of the previous row.
Given an integer target, return true if target is in matrix or false otherwise.

You must write a solution in O(log(m * n)) time complexity
*/

public class Search2DMatrix74 {
    public static boolean search(int[][] matrix, int target) {
        // matrix is mxn, meaning each row, col has same len as each row, col
        // [1, 2] [ matrix[0][0], matrix[0][1] ]
        // [3, 4] [ matrix[1][0], matrix[1][1] ]
        // [4, 5] [ matrix[2][0], matrix[2][1] ]
        // the basic idea is to run two binary searches
        // the first is to find the candidate row (m) that might have the target
        // the second is to find the target in [0, n-1]
        int m = matrix.length, n = matrix[0].length;
        if (m == 0 || matrix == null) {
            return false;
        }
        // 1. Init
        int top = 0, bottom = m - 1; // find the right row
        while (top <= bottom) {
            int mid = top + (bottom - top) / 2;
            // what does it mean to find candidate row? Target lies in the bounds
            if (matrix[mid][0] < target && target < matrix[mid][n - 1]) {
                break;
            }
            // else, try to find next one. Is target above or below the midline?
            // Use property that first elt of any row m is strictly > mat[]
            // Case: below
            else if (matrix[mid][0] < target) {
                top = mid + 1;
            }
            // case: above
            else if (matrix[mid][0] > target) {
                bottom = mid - 1;
            }
        }

        int candidateRow = top + (bottom - top)/2;
        int l = 0, r = n - 1;
        while (l <= r) {
            int mid = l + (r - l) / 2;
            if (matrix[candidateRow][mid] == target) {
                return true;
            } else if (matrix[candidateRow][mid] < target) {
                l = mid + 1;
            } else if (matrix[candidateRow][mid] > target) {
                r = mid - 1;
            }
        }

        return false;
    }

    public static void main(String[] args) {
        int[][] matrix = {
                { 1, 3, 5, 7 },
                { 10, 11, 16, 20 },
                { 23, 30, 34, 60 }
        };
        int target = 3;
        System.out.println(Search2DMatrix74.search(matrix, target));
    }
}
