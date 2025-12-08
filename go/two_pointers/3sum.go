/*
5. 3Sum
Medium

Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

**Notice that the solution set must not contain duplicate triplets.**


Example 1:
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]

Explanation:
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.
*/

package main

import (
	"fmt"
	"os"
	"slices"
)

// NOTE: on intuition
func threeSum(nums []int) [][]int {
	slices.Sort(nums)  // NOTE: this must happen
	res := [][]int{}
	i := 0
	for i < len(nums)-2 {  // stops 3 elts before and marks left boundary
		for i > 0 && i < len(nums)-2 && nums[i] == nums[i-1] {
			i++  // deduplicates
		}
		j, k := i + 1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum < 0 {
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			} else if sum > 0 {
				k--
				for k >= j && nums[k] == nums[k+1] {
					k--
				}
			} else {  // == 0
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				// deduplicate
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				k-- // move j,k since one possible combo found
				// deduplicate
				for k >= j && nums[k] == nums[k+1] {
					k--
				}
			}
		}
		i++
	}
	return res
}


func badSum(nums []int) [][]int {
	res := [][]int{}
	m := make(map[int]struct{})
	for i := range nums {
		for j := i+1; j < len(nums); j++ {
			for k := j +1; k < len(nums); k++ {
				seen := false
				for i := range m {
					if _, ok := m[i]; ok {
						seen = true
					}
				}
				sum := nums[i] + nums[j] + nums[k]
				fmt.Fprintf(os.Stderr, "DEBUG[3]: 3sum.go:41: sum=%+v\n", sum)
				if sum == 0 && !seen {
					res = append(res, []int{nums[i], nums[j], nums[k]})
					m[nums[i]] = struct{}{}
					m[nums[j]] = struct{}{}
					m[nums[k]] = struct{}{}
				}
			}
		}
	}
	return res
}

func main() {
	input := []int{-1,0,1,2,-1,-4}
	ans := threeSum(input)
	fmt.Fprintf(os.Stderr, "DEBUG[1]: 3sum.go:44: ans=%+v\n", ans)
}
