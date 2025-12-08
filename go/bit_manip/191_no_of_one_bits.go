/*
Given a positive integer n, write a function that returns the number of one bits in its binary representation (also known as the Hamming weight).

E.g., 11 |-> 1011 == 3
*/

package main

import "fmt"

func numZeroes(n int) []int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = (i&1) + dp[i>>1]
	}
	return dp
}

func hammingWeight(n int) int {
	ones := 0
	for n > 0 {
		ones += n & 1
		n >>= 1
	}
	return ones
}

func main() {
	n := 11
	fmt.Println(hammingWeight(n))
}
