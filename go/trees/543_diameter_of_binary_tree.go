/*
543. Diameter of Binary Tree - Easy

Given the root of a binary tree, return the *length of the diameter of the tree*.
The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.

The length of a path between two nodes is represented by the number of edges between them.
*/

package main

// idea: get depth of left, right subtreees respectively. Add these together and return
// BAD: doesn't work because the diameter may **or may not** pass thru the root.
func diameterOfBinaryTreeBROKEN(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := depth(root.Left)
	rightHeight := depth(root.Right)
	return leftHeight + rightHeight
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(depth(root.Left), depth(root.Right))
}

func diameterOfBinaryTree(root *TreeNode) int {
	diameter := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)

		// Update diameter
		diameter = max(diameter, left + right)

		// return height of current subtree
		return 1 + max(left, right)
	}
	dfs(root)
	return diameter
}

