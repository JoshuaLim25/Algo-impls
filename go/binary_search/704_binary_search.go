/*
Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.

You must write an algorithm with O(log n) runtime complexity.
*/

package main

func search(nums []int, target int) int {
	// assumed sorted
	if len(nums) == 0 {
		return -1
	}

	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {  // ONE LAST TIME: **lesser** values affect the **left** pointer
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}

func main() {

}
