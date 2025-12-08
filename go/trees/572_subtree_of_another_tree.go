/*
572. Subtree of Another Tree - Easy

Given the roots of two binary trees root and subRoot, return true if there is a subtree of root with the same structure and node values of subRoot and false otherwise.
A subtree of a binary tree tree is a tree that consists of a node in tree and all of this node's descendants. The tree tree could also be considered as a subtree of itself.


The first case is that s and t are two completely identical trees, the second case is that t is a subtree in the left subtree of s, and the third case is that t is a subtree in the right subtree of s. The first case, determining whether two numbers are completely identical, is the original question of question 100.
*/

package main

func isSameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	// either or: one is non-nil
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func isSubtree(root, subRoot *TreeNode) bool {
	if isSameTree(root, subRoot) {
		return true
	}
	if root == nil {
		return false
	}
	if isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot) {
		return true
	}
	return false
}
