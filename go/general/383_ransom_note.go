/*
383. Ransom Note
Easy

Given two strings `ransomNote` and `magazine`, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.

Each letter in magazine can only be used once in ransomNote.

Example 1:
Input: ransomNote = "a", magazine = "b"
Output: false

Example 2:

Input: ransomNote = "aa", magazine = "ab"
Output: false
*/


/*

-'a' gives you -97 (the negative ASCII value of 'a').
In practice, it's used to create a zero-based offset for lowercase letters:

*/

package main

// slice impl
func canConstruct(ransomNote, magazine string) bool {
	if len(ransomNote) > len(magazine) { // can be == or less (i.e., mag has more)
		return false
	}
	// NOTE: use of set is WRONG. ("aa", "aab") |-> true
	runeCounter := make([]int, 26)
	for _, rune := range magazine {
		runeCounter[rune-'a']++
	}
	for _, rune := range ransomNote {
		if runeCounter[rune-'a'] == 0 {
			return false
		}
		runeCounter[rune-'a']--
	}

	return true
}

// map impl
func canConstruct(ransomNote, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	m := make(map[rune]int)
	for _, rune := range magazine {
		m[rune]++
	}
	for _, rune := range ransomNote {
		if count, ok := m[rune]; count == 0 {
			return false
		} else if !ok {
			return false
		}
		m[rune]--
	}
	return true
}
