// package sliding_window
package main

import "fmt"

/*
* Given two strings s and t of lengths m and n respectively, return the minimum window
* of s such that every character in t (including duplicates) is included in the window. If there is no such substring, return the empty string "".
* The testcases will be generated such that the answer is unique.
 */

func mapsEqual(m1, m2 map[rune]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k := range m1 {
		if m1[k] != m2[k] {
			return false
		}
	}
	return true
}

func advanceRight(cur, ref map[rune]int, r int) {
}

func minWindow(s, t string) string {
	if len(t) == 0 || len(t) > len(s) {
		return ""
	}

	ref := make(map[rune]int, len(t))
	for _, rune := range t {
		ref[rune]++
	}
	fmt.Println(ref)

	source := []rune(s)
	letters := make(map[rune]int, len(source)) // double check len hint
	// l, r := 0, 0
	r := 0
	// while not populated, keep filling by moving r ->. Once done stop
	for r < len(source) && !mapsEqual(letters, ref) {
		letters[source[r]]++
		r++
	}
	fmt.Printf("ref: %v, letters: %v", ref, letters)
	fmt.Printf("Maps equal, look at r: %v, %v\n", mapsEqual(letters, ref), r)
	// for _, rune := range source {
	// 	if count, ok := letters[rune]; !ok {
	// 		letters[rune]++
	// 	} else { // in there
	// 		for count != ref[rune] {
	// 			l++
	// 		}
	// 	}
	// }

	return ""
}

func main() {
	var (
		s string
		t string
	)
	s, t = "ADOBECODEBANC", "ABC" // "BANC"
	res := minWindow420(s, t)
	fmt.Printf("Smallest substring: %+v\n", res)
}
