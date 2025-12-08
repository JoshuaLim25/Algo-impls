/*
169. Majority Element
Easy

Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.
*/

package main


// O(n), with O(n) space
func majorityElement(nums []int) int {
	var maj int
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	for num, count := range freq {
		if count > len(nums)/2 {
			maj = num
		}
	}

	return maj
}

// O(nlgn), with O(1) space
func majorityElementSort(nums []int) int {
	slices.Sort(nums)
	return nums[len(nums)/2]
}
