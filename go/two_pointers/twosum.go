package main

import (
	"fmt"
	"slices"
)

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.
*/

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)  // nums: idx
	for i := range nums {
		diff := target - nums[i]
		if _, ok := m[diff]; ok {
			return []int{i, m[diff]}
		}
		m[nums[i]] = i
	}
	return nil
}


func main() {
	// nums := []int{2,7,11,15}
	nums := []int{3,2,4}
	// target := 9
	target := 6
	twoSum(nums, target)
}
