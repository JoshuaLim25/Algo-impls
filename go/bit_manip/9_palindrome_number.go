/*
9. Palindrome Number
Easy
Topics
premium lock iconCompanies
Hint

Given an integer x, return true if x is a

, and false otherwise.

e.g., 121 |-> true, -121 |-> false
*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func isPalindromeSTRING(x int) bool {
	// repr := string(x)   // BAD:
	repr := strconv.Itoa(x)
	fmt.Fprintf(os.Stderr, "DEBUG[1]: 9_palindrome_number.go:18: repr=%+v\n", repr)
	l, r := 0, len(repr)-1
	for l <= r {
		if repr[l] != repr[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func isPalindrome(x int) bool {
	num := x
	// idea is to store theones place number and keep dividing by 10 to "shift" it over (visually like >>)
	res := 0
	// say num == 212
	for num > 0 {
		// how do you get the ones place?
		remainder := num % 10
		num /= 10
		// ORDER MATTERS
		res *= 10
		res += remainder
	}
	return res == x
}

func main() {
	isPalindromeSTRING(121)
}
