/*
977. Squares of a Sorted Array - Easy

Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.

Example 1:
Input: nums = [-4,-1,0,3,10]
Output: [0,1,9,16,100]
Explanation: After squaring, the array becomes [16,1,0,9,100].
After sorting, it becomes [0,1,9,16,100].
*/

package main

func sortedSquares(nums []int) []int {
	// Idea is to have 2 pointers at either end of res
	res := make([]int, len(nums))  // [ 0 0 0 ]
	// k iterates thru the input array from reverse end
	for i, j, k := 0, len(res)-1, len(nums)-1; i <= j; k-- {
		leftSquare := nums[i] * nums[i]
		rightSquare := nums[j] * nums[j]
		if leftSquare > rightSquare {
			res[k] = leftSquare
			i++
		} else {
			res[k] = rightSquare
			j--
		}
	}
	return res
}
