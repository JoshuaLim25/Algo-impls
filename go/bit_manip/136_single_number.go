/*
136. Single Number - Easy
Given a *non-empty* array of integers `nums`, every element appears twice *except for one*. Find that single one.
You must implement a solution with a linear runtime complexity and use only constant extra space - O(n) TC, O(1) SC
*/

package main

func singleNumber(nums []int) int {
	// NOTE: O(1) space, so no set allowed; O(n) TC, so no sorting allowed
	// KEY is usign XOR. Why?
	// If you XOR the same number twice, you get back what you started with
	// Start w/0. 0 ^ 11 => 11 |-> 1011. 11 ^ 11 => 0: 1011 ^ 1011 => 0000
	final := 0
	for _, num := range nums {
		final ^= num
	}
	return final
}
