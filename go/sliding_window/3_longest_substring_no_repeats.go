/* Given a string s, find the length of the longest without duplicate characters. */

package main

import (
	"fmt"
	"os"
)

func lengthOfLongestSubstring(s string) int {
	set := make(map[rune]struct{})
	var (
		curLen int
		l      int
	)
	repr := []rune(s)

	// "abcabcbb"
	for r, rune := range repr {
		curLen = max(curLen, len(set))
		// case 2: conflict
		for l <= r && set[rune] == struct{}{} { // present in set
			delete(set, repr[l])
			l++
		}
		// case 1: no conflict, not in set, just insert
		if _, ok := set[rune]; !ok {
			set[rune] = struct{}{}
		}
	}

	return curLen
}

func main() {
	// input1 := "abcabcbb" // 3
	// input2 := "bbbbb"  // 1
	input3 := "pwwkew" // 3

	length := lengthOfLongestSubstring(input3)
	fmt.Fprintf(os.Stderr, "DEBUG[3]: 3_longest_substring_no_repeats.go:18: length=%+v\n", length)
}
