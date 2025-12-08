/*
338. Counting Bits - Easy

Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n), ans[i] is the number of 1's in the binary representation of i.
*/

package main

import "fmt"

// O(nlgn)
func countBitsSLOW(n int) []int {
	// NOTE: each i is [0,n]
	ans := make([]int, 0, n+1)
	for i := 1; i <= n; i++ {
		onesCount := 0
		bits := fmt.Sprintf("%b", i)
		for _, r := range bits {
			if r == '1' {
				onesCount += 1
			}
		}
		ans = append(ans, onesCount)
	}
	return ans
}

// O(nlgn)
func countBits(n int) []int {
	ans := make([]int, n+1)
	for i := 1; i <= n; i++ {
		countOnes := 0
		num := i
		for num > 0 {
			countOnes += num & 1
			num >>= 1
		}
		ans[i] = countOnes
	}
	return ans
}

// O(n)
func countBitsDP(n int) []int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = (i & 1) + dp[i>>1]
	}
	return dp
}
