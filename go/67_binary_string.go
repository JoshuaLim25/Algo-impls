/*
Given two binary strings a and b, return their sum as a binary string.



Example 1:

Input: a = "11", b = "1"
Output: "100"

convert to []rune to index
init n1, n2 pointers starting from right (len(s)-1), as well as a carry bool
check carry, and if it is 1 and 1. if so, res is 0, carry true
*/

package main

import (
	"fmt"
	"os"
	"strings"
)

func addBinary(a, b string) string {
	// case 1: no carry, with 1+1 |-> 0 with carry
	// case 2: carry, with 1+1 |-> 1 with carry
	reprA := []rune(a)
	reprB := []rune(b)
	var sb strings.Builder
	for i := len(reprA)-1; i >= 0; i-- {
		sb.WriteRune(reprA[i])
	}
	reversedA := sb.String()
	sb.Reset()
	for i := len(reprB)-1; i >= 0; i-- {
		sb.WriteRune(reprA[i])
	}
	reversedB := sb.String()
	runesA, runesB, res := []rune(reversedA), []rune(reversedB), []rune{}
	carry := 0
	var digitA, digitB int
	for i := range max(len(runesA), len(runesB)) { // longer string
		if i < len(runesA) {
			digitA = runesA[i] - '0'  // INFO: acts like strconv.Atoi, only garbage if not a #
		} else {
			digitA = 0
		}
		if i < len(runesB) {
			digitB = runesB[i] - '0'
		} else {
			digitB = 0
		}
		total := digitA + digitB + carry
		total %= 2  // since binary base 2
		// total = 1 |-> 1
		// total = 2 |-> 0
		// total = 3 |-> 1
		res = append(res, total+'0')  // again, like strconv.Itoa()
	}

	return string(res)
}

func main() {
	s := addBinary("blah", "xd")
	fmt.Fprintf(os.Stderr, "DEBUG[1]: 67_binary_string.go:30: s=%+v\n", s)
}

