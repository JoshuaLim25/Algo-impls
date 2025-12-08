/*
Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).
*/

package main


func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(l, r *TreeNode) bool {
	if l == nil && r == nil { return true }
	if l == nil || r == nil { return false }
	return l.Val == r.Val && isMirror(l.Left, r.Right) && isMirror(l.Right, r.Left)
}
