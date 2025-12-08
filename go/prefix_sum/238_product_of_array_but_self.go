/*
238. Product of Array Except Self - Medium

Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
You must write an algorithm that runs in O(n) time and without using the division operation.

Input: nums = [1,2,3,4]
Output: [24,12,8,6]
*/

package main

import "fmt"

func calcProducts(nums []int) []int {
	res := make([]int, len(nums))
	prod := 1
	for i, n := range nums {
		prod *= n
		res[i] = prod
	}
	return res
}

func productExceptSelfBORKEN(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	// have one map that has products of everything in nums[:elt]
	// have another map that has products of everything in nums[elt+1:]
	prods := calcProducts(nums)
	fmt.Printf("%+v\n", prods)
	answer := make([]int, len(nums))
	for i := range nums {
		// a[i] = prods[:elt] + prods[elt+1:]
		var prodBefore, prodAfter int
		if i > 0 {
			prodBefore = prods[i-1]  // nums[0] + nums[1] + ... + nums[elt-1]
			// what about nums[n-1]? 50-50 = 0, so no worries
			prodAfter = prods[len(prods)-1] - prods[i]
		}
		fmt.Printf("prodBefore, prodAfter: %d, %d\n", prodBefore, prodAfter)
		answer[i] = prodBefore + prodAfter
	}
	return answer
}

func productExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	l := len(nums)
	prefix := make([]int, l)
	prefix[0] = 1
	suffix := make([]int, l)
	suffix[l-1] = 1
	answer := []int{}
	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}
	for i := len(nums)-2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i+1]
	}
	for i := range nums {
		// Prefix[i] = product of all elements to the LEFT of index i
		// Suffix[i] = product of all elements to the RIGHT of index i
		answer = append(answer, prefix[i]*suffix[i])
	}
	return answer
}


func main() {
	nums := []int{1,2,3,4}  // output := []int{24,12,8,6}
	res := productExceptSelf(nums)
	fmt.Printf("%+v == {24,12,8,6}\n", res)
}
