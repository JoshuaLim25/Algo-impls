/*
- You are given an array of words where each word consists of lowercase English letters.
- wordA is a predecessor of wordB if and only if we can insert exactly one letter anywhere in wordA without changing the order of the other characters to make it equal to wordB.
	- For example, "abc" is a predecessor of "abac", while "cba" is not a predecessor of "bcad".
- A word chain is a sequence of words [word1, word2, ..., wordk] with k >= 1, where word1 is a predecessor of word2, word2 is a predecessor of word3, and so on. A single word is trivially a word chain with k == 1.
- Return the length of the longest possible word chain with words chosen from the given list of words.

Input: words = ["a","b","ba","bca","bda","bdca"]
Output: 4
Explanation: One of the longest word chains is ["a","ba","bda","bdca"].
*/

package main

import (
	"fmt"
	"sort"
)

// NOTE: for every word, assume it to be a unique entity and assign a chain length of 1. This is the base case, "trivially a word chain with k == 1"
func longestStrChain(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	dp := make(map[string]int)  // tracks predecessors and len(prevChain) 
	longestLen := 0
	for _, word := range words {
		// assumption is to init to 1, trivial case
		dp[word] = 1
		// go through word, splicing a char out to try and make a predecessor
		// i.e., "abc" -> ["bc", "ac", "ab"]
		for i := range len(word) {
			candidate := word[:i] + word[i+1:]
			if prevChainLen, ok := dp[candidate]; ok {
				// NOTE: the +1
				dp[candidate] = max(dp[candidate], prevChainLen+1)
			}
		}
		longestLen = max(longestLen, dp[word])
	}
	return longestLen
}



// func longestStrChain(words []string) int {
// 	sort.Slice(words, func(i, j int) bool {
// 		return len(words[i]) < len(words[j])
// 	})
// 	maxLen := 0
// 	dp := make(map[string]int)
// 	for _, word := range words {
// 		dp[word] = 1
// 		// go through each letter in word, taking each out and testing if in map. This would be a predecessor, since its len is -1.
// 		// if it's (potential predecessor) is in dp, it means that the current word could have come from it.
// 		for i := range len(word) {
// 			// candidate, could have been a predecessor
// 			trimmed := word[:i] + word[i+1:]
// 			if prevLen, ok := dp[trimmed]; ok {
// 				dp[word] = max(dp[word], prevLen+1)
// 			}
// 		}
// 		maxLen = max(maxLen, dp[word])
// 	}
// 	return maxLen
//
// }

func main() {
	words := []string{"a","b","ba","bca","bda","bdca"}
	res := longestStrChain(words)
	fmt.Printf("%v\n", res)
	
}
