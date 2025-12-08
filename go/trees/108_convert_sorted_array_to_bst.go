/*
108. Convert Sorted Array to Binary Search Tree - Easy

Given an integer array nums where the elements are sorted in ascending order, sortedArrayToBST it to a
binary search tree.
*/

package main

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return &TreeNode{
		Val: nums[len(nums)/2],
		Left: sortedArrayToBST(nums[:len(nums)/2]),
		Right: sortedArrayToBST(nums[len(nums)/2+1:]),
	}
}

