/*
Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.
*/

package main

import "fmt"

// [ 1 4 5 8 7 ]
func permute(nums []int) [][]int {
	//choices, constraints, goal
	res := [][]int{}
	perms := []int{}
	used := make([]bool, len(nums)) // [f f f f]
	var backtrack func()
	backtrack = func() {
		if len(perms) == len(nums) {
			// res = append(res, append([]int{}, perms...))  // must do this
			tmp := make([]int, len(perms))
			copy(tmp, perms)
			res = append(res, tmp)
			return
		}
		for i := range nums {
			//make choice
			if !used[i] {
				used[i] = true
				perms = append(perms, nums[i])
				backtrack()
				// undo choice
				used[i] = false
				perms = perms[:len(perms)-1]
			}
		}
	}
	backtrack()
	return res
}

func main() {
	nums := []int{1,2,3}
	res := permute(nums)
	fmt.Printf("%+v should == [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]\n", res)
}
