package main

import (
	"fmt"
	"math"
	"os"
)

/*
Suppose an array of length n sorted in ascending order is rotated between 1 and n times. For example, the array nums = [0,1,2,4,5,6,7] might become:

    [4,5,6,7,0,1,2] if it was rotated 4 times.
    [0,1,2,4,5,6,7] if it was rotated 7 times.

Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].

Given the sorted rotated array nums of unique elements, return the minimum element of this array.

You must write an algorithm that runs in O(log n) time.

*/

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	// equal case unhandled, since input guaranteed unique
	// go thru.
	// target = nums[0]
	smallest := math.MaxInt
	for l <= r {
		mid := l + (r-l)/2
		// sorted: a[n-1] <= a[n] <= a[n+1]
		if nums[mid] < smallest {
			smallest = nums[mid]
		}

		// BAD:
		// if mid+1 <= len(nums)-1 && nums[mid] > nums[mid+1] {
		// 	smallest = nums[mid+1]
		// }
		// if mid-1 >= 0 && nums[mid] < nums[mid-1] {
		// 	smallest = nums[mid-1]
		// }

		if nums[mid] > nums[l] { // 2 3 4 5
			if nums[l] < smallest {
				smallest = nums[l]
			}
			l = mid + 1
		} else { // 3 4 5 0 2
			if nums[r] < smallest {
				smallest = nums[r]
			}
			r = mid - 1
		}
	}
	return smallest
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2} // 0
	// nums := []int{3, 4, 5, 1, 2} // 1
	// nums := []int{11, 13, 15, 17} // 11
	minVal := findMin(nums)
	fmt.Fprintf(os.Stderr, "DEBUG[1]: min_rotated_arr.go:40: nums=%+v\n", nums)
	fmt.Fprintf(os.Stderr, "DEBUG[2]: min_rotated_arr.go:37: minVal=%+v\n", minVal)
}
