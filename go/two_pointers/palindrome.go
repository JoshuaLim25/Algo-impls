// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.
// Given a string s, return true if it is a palindrome, or false otherwise.
//
// Example 1:
// Input: s = "A man, a plan, a canal: Panama"
// Output: true
// Explanation: "amanaplanacanalpanama" is a palindrome.
//
// Example 2:
// Input: s = "race a car"
// Output: false
// Explanation: "raceacar" is not a palindrome.
//
// Example 3:
// Input: s = " "
// Output: true
// Explanation: s is an empty string "" after removing non-alphanumeric characters.
// Since an empty string reads the same forward and backward, it is a palindrome.

package main

import (
	"fmt"
	"unicode"
)

func isValid(s string) bool {
	valid := true
	var tmp []rune
	for _, c := range s {
		if unicode.IsLetter(c) {
			tmp = append(tmp, unicode.ToLower(c))
		}
	}
	fmt.Printf("tmp: %+v\n", tmp)
	stripped := string(tmp)
	fmt.Printf("stripped: %+v\n", stripped)

	i, j := 0, len(stripped)-1
	for i != j && j >= 0 {
		if stripped[i] == stripped[j] {
			fmt.Printf("samesies: %v, %v\n", stripped[i], stripped[j])
			i++
			j--
		} else {
			valid = false
			break
		}
	}
	return valid
}


func main() {
	s1 := "A man, a plan, a canal: Panama"
	v1 := isValid(s1)
	fmt.Printf("string %v is valid: %v\n", s1, v1)

	s2 := "race a car"
	v2 := isValid(s2)
	fmt.Printf("string %v is valid: %v\n", s2, v2)


	s3 := " "
	v3 := isValid(s3)
	fmt.Printf("string %v is valid: %v\n", s3, v3)
}
