package sliding_window

/* 1.3 Longest Substring with K Distinct Characters \(medium\)
- Stmt: Given a string, find the length of the longest substring in it with no more than K distinct characters.
- use m := map[rune]int, numDistinct := 0
- 1st char, what are my decisions?
	- if already in map,  increment m[rune], advance high idx
	- if not, insert into map and numDistinct++
- what causes numdistinct--?  ->  when incrementing windowLeft && after checking for existence in map, deleting entry

- Example 1:
Input: String="araaci", K=2
Output: 4
Explanation: The longest substring with no more
than '2' distinct characters is "araa".

- Example 2:
Input: String="araaci", K=1
Output: 2
Explanation: The longest substring with no more
than '1' distinct characters is "aa".

- Example 3:
Input: String="cbbebi", K=3
Output: 5
Explanation: The longest substrings with no more
than '3' distinct characters are "cbbeb" & "bbebi".
*/

// TODO: find len(longestSubstring) in `s` w/no more than `k` distinct chars
// Input: String="araaci", K=2 (distinctCount <= k)
// Output: 4
// For the first character `a`, our decisions are:
// 1. Check our map for existence.
// if so: incr count and if not, add character to map
// 2. Check that count hasn't exceeded (cnt <= k) numDistinct
// if so: incr low, incr high; then check again (update count)
// if not, continue incr high
// Edge: if there's two with same len, take either

// TAKEAWAY: there's a relationship between how the left side of the window grows and the number of distinct characters you have in the string you're iterating over, which gets checked by cond `numDistinct > k`. Getting clear what makes something go into the frequencies map and what makes it get removed was important, as well as the updates you had to do depending on what action was taken.
func LongestKDistinct(s string, k int) (longest int) {
	if len(s) == 0 || k == 0 {
		return 0
	}
	low := 0 // NOTE: updates dependent on numDistinct being > k
	numDistinct := 0
	frequencies := make(map[rune]int)
	for i := range s {
		c := rune(s[i])
		// 1. check for existence. If not there, insert. If there, incr freq
		// you delete an entry when numDistinct > k. If freq == 1, delete and decrement numDistinct, otherwise freq-- and advance
		freq, isPresent := frequencies[c]
		if isPresent {
			frequencies[c]++
		} else { // not in map
			frequencies[c] = 1
			numDistinct++
		}
		if numDistinct > k {
			if freq == 1 { // if num times this char we're on appeared is once
				delete(frequencies, c)
				numDistinct--
			} else { // has appeared more than once, e.g., "ara"
				freq--
			}
			low++
		}
		if longest < i-low+1 {
			longest = i - low + 1
		}
	}
	return longest
}

// func LongestKDistinct(s string, k int) int {
// 	// distinct -> set
// 	frequencies := make(map[rune]int) //
// 	// k = 2, s = "araaci"
// 	highest := -1
// 	distinctChars := 0 // OBS: only affected by insert/delete into the map
// 	low := 0
// 	for _, c := range s { // Have we seen "a"? Is "a" in our map?
// 		val, present := frequencies[c] // lookup
// 		if present == false {
// 			val = 1
// 			distinctChars++
// 		} else {
// 			if distinctChars > k { // we've seen it, now sign to shift window
// 				low++
// 				if val == 1 { // you did see it, was it the first time?
// 					delete(frequencies, c) // it was only put in once, so remove
// 					distinctChars--        // and update distinct num chars
// 				}
// 			} else { // still have seen key, but distinct num <= k, so good
// 				val++
// 			}
// 		}
// 		if distinctChars > highest {
// 			highest = distinctChars
// 		}
// 	}
// 	return highest
// }
