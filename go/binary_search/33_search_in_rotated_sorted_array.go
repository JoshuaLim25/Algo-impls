/*
33. Search in Rotated Sorted Array - Medium

There is an integer array nums **sorted in ascending order** (with *distinct* values).
Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity.
*/

package main

func search(nums []int, target int) int {
	// idea: use binary search, but consider the possibility of rotation
	// need to check which "side" of the rotation we're on: [4,5,6,7,0,1,2]
	// Left: the og right side, w/big nums - [4567] <=> [l..mid]
	// Right: the original left side, w/small nums - [012] <=> [mid..right]
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if target == nums[mid] {
			return mid
		}
		// [4,5,6,7,0,1,2], target = 0, l = 0, r = 3
		// If not distinct, equals "hugs the side you're on"
		if nums[l] <= nums[mid] {
			// on the left sorted side, so [l..mid]
			// is it in the range [456]?
			if target >= nums[l] && target < nums[mid] {
				r = mid - 1 // it's in this range
			} else {
				l = mid + 1 // not in the left range, so move left ptr
			}
		} else {
			// on the right side, so [mid..r]
			// is it in the range [7012]?
			if target > nums[mid] && target <= nums[r] {
				l = mid + 1 // in the range
			} else {
				r = mid - 1 // not in the right range, so move right ptr
			}
		}
	}
	return -1
}
