package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}
	if p.Val < root.Val && q.Val < root.Val { // both belong in left subtree
		lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		lowestCommonAncestor(root.Right, p, q)
	}

	return root
}
