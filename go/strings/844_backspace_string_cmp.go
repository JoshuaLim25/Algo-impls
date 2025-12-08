/*
844. Backspace String Compare - Easy

Given two strings s and t, return true if they are equal when both are typed into empty text editors. '#' means a backspace character.

Note that after backspacing an empty text, the text will continue empty.
*/

// package strings
package main

func buildString(s string) string {
	// var sb strings.Builder
	// for _, r := range s {
	// 	sb.WriteRune(r)
	// }
	res := []rune{}
	for _, r := range s {
		if r == '#' {
			if len(res) > 0 {
				res = res[:len(res)-1]
				continue
			}
		} else {
			res = append(res, r)
		}
	}
	return string(res)
}

func backspaceCompare(s, t string) bool {
	first := buildString(s)
	second := buildString(t)
	if first != second {
		return false
	}
	return true
}

func main() {
	s := "a#c"
	t := "b"
	backspaceCompare(s, t)
}
