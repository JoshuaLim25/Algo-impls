package main

/*
Given a circular integer array nums (i.e., the next element of nums[nums.length - 1] is nums[0]), return the next greater number for every element in nums.

The next greater number of a number x is the first greater number to its traversing-order next in the array, which means you could search circularly to find its next greater number. If it doesn't exist, return -1 for this number.

- how would we do for normal arry?
	- general strat is to keep a monotonically *decreasing* stack

- for circular, need to track curIdx and save it. Then use another pointer to traverse from nums[curIdx:]. If curIdx is >= 0, need to traverse a bit more, so nums[:curIdx]
-- say we had nums = [1 2 3 4]. len == 4, curIdx == 2, nums[curIdx] == 3
traverse := curIdx; traverse < len(nums); traverse++

*/

import "fmt"

func nextGreater(nums []int) []int {
	res := make([]int, len(nums))
	s := []int{} // NOTE: tracks indices of lesser elts
	for i := range res {
		res[i] = -1
	}

	for i := range len(nums) * 2 { // DIFF
		num := nums[i%len(nums)] // DIFF
		for len(s) > 0 && num > nums[s[len(s)-1]] {
			top := s[len(s)-1]
			s = s[:len(s)-1]
			res[top] = num
		}
		s = append(s, i%len(nums)) // DIFF
	}
	return res

}

// BAD:
func nextGreaterBAD(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	res := make([]int, len(nums))
	for i := range res {
		res[i] = -1
	}
	s := []int{} // tracks lesser val indices

	for i, num := range nums {
		l := len(s)
		for len(s) > 0 && num > nums[s[len(s)-1]] { // repeatedly pop
			top := s[len(s)-1] // idx of prev elt in nums that's lesser in value
			s = s[:len(s)-1]
			res[top] = num
		}
		// wrap around if nonzero i and no pops happened (length check)
		if i != 0 && l == len(s) {
			curIdx := i
			// check nums[:i]
			for range curIdx { // 0..i, not inclusive i
				if len(s) > 0 && num > nums[s[len(s)-1]] {
					top := s[len(s)-1] // idx of prev elt in nums that's lesser in value
					s = s[:len(s)-1]
					res[top] = num
				}
			}
		}
		s = append(s, i)
	}
	return res
}

func main() {
	// nums := []int{1, 2, 1}  // [2 -1 2]
	nums := []int{1, 2, 3, 4, 3} // [2,3,4,-1,4]
	ans := nextGreater(nums)
	fmt.Printf("ans=%+v\n", ans)
}
