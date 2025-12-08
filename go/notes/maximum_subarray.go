package main

import (
	"fmt"
	"math"
	"os"
)

func Sum(data[]int, start, end int) int {
	var sum int
	for beg := start; beg <= end; beg++ {
		sum += data[beg]
	}
	return sum
}

func MaxSubarray(nums []int) int {
	maxSum := math.MinInt
	l := 0
	var lastSum int
	for r := range nums {
		sum := Sum(nums, l, r)
		fmt.Printf("i: %v, j: %v, sum: %v, maxSum: %v\n", r, l, sum, maxSum)
		if sum > maxSum {
			maxSum = sum
		}
		if maxSum < lastSum {
			for r-l > 0 {
				l++
			}
		}
	}
	return maxSum
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Fprintf(os.Stderr, "DEBUG[6]: maximum_subarray.go:40: nums=%+v\n", nums)
	sum := MaxSubarray(nums)
	fmt.Fprintf(os.Stderr, "DEBUG[4]: maximum_subarray.go:42: sum=%+v\n", sum)
	// nums2 := []int{5, 4, -1, 7, 8}
	// sum2 := MaxSubarray(nums2)
	// fmt.Fprintf(os.Stderr, "DEBUG[5]: maximum_subarray.go:44: sum2=%+v\n", sum2)
}
