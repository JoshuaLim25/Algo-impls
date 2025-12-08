/*
Given an integer array nums, handle multiple queries of the following type:

    Calculate the sum of the elements of nums between indices left and right inclusive where left <= right.

Implement the NumArray class:

    NumArray(int[] nums) Initializes the object with the integer array nums.
    int sumRange(int left, int right) Returns the sum of the elements of nums between indices left and right inclusive (i.e. nums[left] + nums[left + 1] + ... + nums[right]).


*/

package main

import (
	"fmt"
)

type NumArray []int

func NewNumArray(nums []int) *NumArray {
	n := NumArray(nums)
	return &n
}

func (this *NumArray) sumRange(left, right int) int {
	sum := 0
	m := make(map[int]int) // idx: accumulated sum
	// constraints
	if right >= len(*this) || left > right {
		return 0
	}
	for r := 0; r <= right; r++ {
		sum += (*this)[r]
		m[r] = sum
		// fmt.Fprintf(os.Stderr, "DEBUG[2]: 303_range_sum_query.go:34: sum=%+v\n", sum)
	}
	if left-1 >= 0 {
		// fmt.Printf("Gone thru til right. m[left-1] is %d\n", m[left-1])
		// fmt.Printf("sum before: %d\n", sum)
		sum -= m[left-1]
		// fmt.Printf("sum after: %d\n", sum)
	}
	return sum
}

type FasterNumArray struct {
	prefixSum []int
}

func NewFasterNumArray(nums []int) FasterNumArray {
	// start:  [1 2 3 4]
	// i := 1: [1 2+1 3 4]
	// i := 2: [1 3 3+3 4]
	// i := 3: [1 3 6 4+6]
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return FasterNumArray{prefixSum: nums}
}

func (n *FasterNumArray) SumRange(left, right int) int {
	if left > right {
		return 0
	}
	sum := 0
	if left <= right && left-1 >= 0 { // NOTE: second cond matters
		sum = n.prefixSum[right] - n.prefixSum[left-1] // bc this access
	} else if left-1 < 0 {
		sum = n.prefixSum[right]
	}
	return sum
}

/*

 */

func main() {
	// n := NewNumArray([]int{-2, 0, 3, -5, 2, -1})
	// n := NewNumArray([]int{-4, -5})
	// n := NewFasterNumArray([]int{-2, 0, 3, -5, 2, -1})
	n := NewFasterNumArray([]int{-4, -5})
	fmt.Printf("n: %+v\n", n)
	// l, r := 0, 0
	// l, r := 0, 1
	// l, r := 1, 0 // BAD
	l, r := 1, 1
	fmt.Printf("l: %d, r: %d\n", l, r)
	sum := n.SumRange(l, r)
	fmt.Printf("Final sum: %d\n", sum)
	// range2 := n.sumRange(2, 5)
	// range3 := n.sumRange(0, 5)

}
