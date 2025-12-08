package main

import "fmt"

func searchRotated(nums []int, target int) int {
	// {4,5,6,7,0,1,2}, target = 0
	// 1. find midpt
	// 2. check: is midpt == target, if so return
	// 3. check: are we in the "left side" of the rotation or the "right"
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		// if nums[mid] >= nums[l] { // distinct
		if nums[mid] > nums[l] { // on the left of the split
			if target > nums[mid] || target < nums[l] {
				l = mid + 1
			} else { // target > left, but < midpt, so l < target < mid: in the sweet spot
				r = mid - 1
			}
		} else { // on the right
			if target < nums[mid] || target > nums[r] {
				r = mid - 1
			} else { // in the zone already, stay
				l = mid + 1
			}
		}
	}

	return -1
}

func main() {
	data := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	fmt.Printf("data: %+v\n\tIndex: %d\n", data, searchRotated(data, target))
}
