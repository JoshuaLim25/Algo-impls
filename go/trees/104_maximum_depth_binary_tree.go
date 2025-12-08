/*
104. Maximum Depth of Binary Tree
Easy

Given the root of a binary tree, return its maximum depth.

A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BFS
func maxDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := []*TreeNode{root}
	depth := 0
	for len(q) > 0 {
		depth++
		for range len(q) {
			nextUp := q[0]
			q = q[1:]

			// process
			if nextUp.Left != nil {
				q = append(q, nextUp.Left)
			}
			if nextUp.Right != nil {
				q = append(q, nextUp.Right)
			}

		}
	}
	return depth
}

// DFS
func maxDepthDFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(root.Left, root.Right) + 1
}
