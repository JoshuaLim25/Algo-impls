/*
110. Balanced Binary Tree - Easy

Given a binary tree, determine if it is height-balanced, meaning that its left and right subtree heights differ by at most 1.

Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: true

Example 2:

Input: root = [1,2,2,3,3,null,null,4,4]
Output: false
*/

package com.algos.tree;

public class BalancedBinaryTree110 {
    public boolean isBalanced(TreeNode root) {
        if (root == null)
            return true;

        int leftHeight = height(root.left);
        int rightHeight = height(root.right);

        if (leftHeight == rightHeight || leftHeight == rightHeight + 1 || leftHeight == rightHeight - 1)
            return true;
        return false;
    }

    private int height(TreeNode root) {
        if (root == null)
            return 0;
        return 1 + Math.max(height(root.left), height(root.right));
    }
}
