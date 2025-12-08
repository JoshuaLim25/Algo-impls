/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:
    Open brackets must be closed by the same type of brackets.
    Open brackets must be closed in the correct order.
    Every close bracket has a corresponding open bracket of the same type.
*/

package main


func isValid(s string) bool {
	m := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}
	stack := []rune{}
	for _, rune := range s {
		if openingBrace, isClosing := m[rune]; isClosing {
			// opening brace *must* be there && top of stack must be the corresp. opening brace 
			if len(stack) == 0 || stack[len(stack)-1] != openingBrace {
				return false
			}   
			// it met criteria, pop off stack, discard top
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, rune)
		}
	}
	return len(stack) == 0
}

func main() {

}
