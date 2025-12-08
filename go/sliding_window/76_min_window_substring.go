/*
Given two strings s and t of lengths m and n respectively, return the minimum window of s such that every character in t (including **duplicates**) is included in the window. If there is no such substring, return the empty string "".

The testcases will be generated such that the answer is unique.


m == s.length
n == t.length
1 <= m, n <= 105
*s and t consist of uppercase and lowercase English letters.*

Edge cases: here, it was lenth of m, n (n is a subset of n?), duplicate chars, upper and lowercase chars.
*/

package main

import (
	"fmt"
	"os"
)

func minWindow420(s, t string) string {
	if len(s) == 0 || len(t) == 0 || len(t) > len(s) { // t subset of s
		return ""
	}
	equals := []string{}
	acc := make([]rune, 0, len(t)) // not really a consideration, since ascii
	lookup := make(map[rune]int)   // char and count
	m := make(map[rune]int)
	repr := []rune(s)

	for _, rune := range t {
		lookup[rune]++
	}

	l := 0
	isEqual := func() bool {
		for k, v := range lookup {
			if m[k] != v {
				return false
			}
		}
		return true
	}
	for r, rune := range s {
		// as we range, we always add the rightmost value. Then we check if our map has the right frequency of the char.
		acc = append(acc, rune)
		m[rune]++
		// here, they could be equal (all of t in s), or not equal (either not filled enough or there's redundancies)

		if isEqual() {
			equals = append(equals, string(repr))
		} else {
			if len(m) != len(lookup) { // t not contained in s
				continue
			}
			fmt.Printf("maps not equal, m: %+v has more than needed. acc: %+v\n", m, string(acc))
			leftVal := repr[l]
			// if decrementing freq would keep all of t inside, do so
			fmt.Printf("for %d < %d && %d >= %d\n", l, r, m[leftVal]-1, lookup[leftVal])
			for l <= r && m[leftVal]-1 >= lookup[leftVal] {
				m[leftVal]--
				acc = acc[1:] // remove leftmost elt
				l++
			}
		}
	}
	minSubset := 0
	for i, candidate := range equals {
		if len(candidate) < len(equals[i]) {
			minSubset = i
		}
	}
	return equals[minSubset]
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	ans := minWindow420(s, t) // "BANC"
	fmt.Fprintf(os.Stderr, "DEBUG[6]: 76_min_window_substring.go:69: ans=%+v\n", ans)

}
