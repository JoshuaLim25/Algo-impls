package main

/*
true if t is an anagram of s
An anagram is a word or phrase formed by rearranging the letters of a different word or phrase, using all the original letters exactly once.
*/


// slice impl
func isAnagram(s, t string) bool {
	if len(s) == 0 || len(t) == 0 || len(t) > len(s) {
		return false
	}
	freq := make([]int, 26)
	for _, r := range t {
		freq[r-'a']++
	}
	for _, r := range s {
		if freq[r-'a'] == 0 {
			return false
		}
		freq[r-'a']--
	}

	return true
}




// map impl
func isAnagramMap(s, t string) bool {
	// s = "anagram", t = "nagaram"
	if len(s) == 0 || len(t) > len(s) {
		return false
	}
	m := make(map[rune]int) // tracks count of each rune in s
	inTarget := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}
	// a: 2, n: 1, etc
	for _, r := range t {
		inTarget[r]++
	}
	for k, v := range m {
		if inTarget[k] != v {
			return false
		}
	}
	return true
}

func main() {

}
