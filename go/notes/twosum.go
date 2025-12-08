package main

import (
	"fmt"
	"os"
)

func twoSum(nums []int, target int) (int, int) {
	m := make(map[int]int)
	for i, v := range nums {
		m[v] = i
	}

	for i := range nums {
		complement := target - nums[i]
		if j, ok := m[complement]; ok && j != i {
			return i, j
		}
	}
	return 0, 0
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	target := 5
	i, j := twoSum(nums, target)
	fmt.Fprintf(os.Stderr, "DEBUG[4]: twosum.go:29: i, j are %#v and %#v \n", i, j)

}
