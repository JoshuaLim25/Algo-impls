/*
409. Longest Palindrome
Easy
Topics
premium lock iconCompanies

Given a string s which consists of lowercase or uppercase letters, return the length of the longest

 that can be built with those letters.

Letters are case sensitive, for example, "Aa" is not considered a palindrome.


*/

package main

func longestPalindrome(s string) int {
	m := make(map[rune]int) // freq
	for _, r := range s {
		m[r]++
	}

	var (
		longestPossible int
		oddsCount       int
	)
	for k, freq := range m {
		if freq%2 == 0 { // even means it can be added
			longestPossible += freq
		} else {
			longestPossible += freq - 1 // "ccc"
			oddsCount++
		}
	}
	if oddsCount != 0 {
		return longestPossible + 1
	}
	return longestPossible

}
