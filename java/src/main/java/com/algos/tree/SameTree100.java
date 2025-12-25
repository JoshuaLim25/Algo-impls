/*
100. Same Tree
Solved
Easy
Topics
conpanies iconCompanies

Given the roots of two binary trees p and q, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.



Example 1:

Input: p = [1,2,3], q = [1,2,3]
Output: true

Example 2:

Input: p = [1,2], q = [1,null,2]
Output: false

Example 3:

Input: p = [1,2,1], q = [1,1,2]
Output: false
*/

package com.algos.tree;

public class SameTree100 {
    public boolean isSameTree(TreeNode p, TreeNode q) {
        // either/or check
        if (p == null && q == null) return true;
        if (p == null || q == null) return false;  // one is, one isn't

        if (p.val == q.val) {
            if (isSameTree(p.left, q.left) && isSameTree(p.right, q.right)) return true;
        }
        return false;
    }
}
