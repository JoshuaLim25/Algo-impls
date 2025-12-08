package main

import "math"

// a balanced tree has levels that dont differ by at most one level

func BROKENisBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right != nil {
		if root.Left.Left != nil {
			return false
		}
	}
	if root.Right == nil && root.Left != nil {
		if root.Right.Right != nil {
			return false
		}
	}
	BROKENisBalanced(root.Left)
	BROKENisBalanced(root.Right)
	return true
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftHeight := depth(root.Left)
	rightHeight := depth(root.Right)
	// BAD:
	// if math.Abs(float64(rightHeight-leftHeight)) > 1 {
	// 	return false
	// }
	// return true
	return math.Abs(float64(rightHeight-leftHeight)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}
func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(depth(root.Left), depth(root.Right))
}


