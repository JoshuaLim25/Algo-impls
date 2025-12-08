package main

import "slices"

func longestConsecutive(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	slices.Sort(nums)
	longest, counter := 0, 1
	for i := 1; i < len(nums); i++ {
		// [1234]
		if nums[i-1] == nums[i]-1 { // ascending by +1
			counter++
		} else if nums[i] != nums[i-1] {
			// [12234]
			// this run was a flop, but could be one later
			counter = 1
		}
		longest = max(longest, counter)
	}
	return longest
}
