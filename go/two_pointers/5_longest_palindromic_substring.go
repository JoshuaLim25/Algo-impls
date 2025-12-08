/*

*/
package main

func longestPalindrome(s string) string {
	if len(s) < 3 {
		return ""
	}
	repr := []rune(s)
	l := 0
	palindromes := [][]string{}
	for r := 1; r < len(repr); r++ {
		if s[l] == s[r] { // palindrome found
			palindromes = append(palindromes, s[l:r+1])
		} else if r+1 < len(s) {
			if s[l] == s[r] { // palindrome found
				palindromes = append(palindromes, s[l:r+1])
			}
		} // here, neither next or next next char was equal
		l++
	}
}


func main() {
	s := "babad" // "bab"

}
