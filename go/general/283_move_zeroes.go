/*
283. Move Zeroes - Easy

Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
Note that you must do this in-place without making a copy of the array.

Example 1:
Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]

*/

package main

func moveZeroesBROKEN(nums []int) {
	for i := range nums {
		cur := i
		if nums[i] == 0 && i != len(nums)-1 {
			for cur < len(nums) && cur+1 < len(nums) {
				nums[cur], nums[cur+1] = nums[cur+1], nums[cur]
				cur++
			}
		}
		// wrong, because it fucks up the corner case of 0 @ end
	}
}

func moveZeroes(nums []int) {
	// i: 0 [ 0 1 0 0 0 3 0 ], l: 0
	// i: 1 [ 1 0 0 0 0 3 0 ], l: 1
	// i: 5 [ 1 3 0 0 0 0 0 ], l: 2
	l := 0
	for r := range nums {
		if nums[r] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
	}
}
