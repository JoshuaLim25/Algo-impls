/*
14. Longest Common Prefix - Easy

Write a function to find the longest common prefix string amongst an array of strings.
If there is no common prefix, return an empty string "".

Example 1:
Input: strs = ["flower","flow","flight"]
Output: "fl"
*/

package main

import (
	"slices"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	var sb strings.Builder
	// lexigraphically sort strs: ["ab", "ef", "xy"]
	slices.Sort(strs)
	first, last := strs[0], strs[len(strs)-1]
	for i := range first {
		if i < len(last) && first[i] == last[i] {
			sb.WriteByte(first[i])
		} else {
			break
		}
	}
	return sb.String()
}
