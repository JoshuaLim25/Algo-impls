package main

import (
	"fmt"
	"os"
)

type Stack []rune

func (s *Stack) Push(elt rune) {
	*s = append(*s, elt)
}

func (s *Stack) Pop() rune {
	if len(*s) == 0 {
		return '0'
	}

	top := (*s)[0]
	*s = (*s)[1:]
	return top
}

func ValidateString(s string) bool {
	valid := true
	stack := Stack{}
	m := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}

	for _, char := range s {
		switch char {
		case ']':
			{
				top := stack.Pop()
				if val, _ := m[char]; val != top {
					return false
				}

			}
		case ')':
			{
				top := stack.Pop()
				if val, _ := m[char]; val != top {
					return false
				}
			}
		case '}':
			{
				top := stack.Pop()
				if val, _ := m[char]; val != top {
					return false
				}
			}
		default:
			stack.Push(char)
		}
	}
	return valid
}

func main() {
	s1 := "{}()(()){{}}"
	s2 := "{}"
	s3 := "{}()(()){}{}}"
	ans1 := ValidateString(s1)
	fmt.Fprintf(os.Stderr, "DEBUG[3]: valid_parenthesis_string.go:68: ans1=%+v\n", ans1)
	ans2 := ValidateString(s2)
	fmt.Fprintf(os.Stderr, "DEBUG[4]: valid_parenthesis_string.go:70: ans2=%+v\n", ans2)
	ans3 := ValidateString(s3)
	fmt.Fprintf(os.Stderr, "DEBUG[5]: valid_parenthesis_string.go:72: ans3=%+v\n", ans3)
}
