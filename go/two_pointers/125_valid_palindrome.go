/*
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.
*/

package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	// EDGE CASES:
	// empty string, len == 1
	// upper/lower, non-ascii

	// 0. input validation
	//1. converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters

	if len(s) <= 1 { // " ", "2", "" all qualify
		return true
	}
	// "aaaAAAaAA" |-> "aaaaaaa"
	repr := []rune{}
	fmt.Printf("%+v\n", repr)
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r){
			repr = append(repr, unicode.ToLower(r))
		}
	}
	fmt.Printf("%s\n", string(repr))

	l, r := 0, len(repr)-1
	for l < r {
		if repr[l] == repr[r] {
			l++
			r--
		} else {
			return false
		}
	}
	return true
}

func main() {
	// s := "A man, a plan, a canal: Panama"
	s2 := "race a car"
	boolyean := isPalindrome(s2)
	fmt.Println(boolyean)
	
}
