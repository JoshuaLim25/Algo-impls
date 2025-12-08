/*
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.
*/

package main

import (
	"fmt"
	"strconv"
)

func letterCombinationsBAD(digits string) []string {
	// init a map of digit to possible letters
	// map[digit]slice of possible letters
	if len(digits) == 0 {
		return nil
	}
	lookup := map[int][]byte{
		2: {'a', 'b', 'c'},
		3: {'d', 'e', 'f'},
		4: {'g', 'h', 'i'},
		5: {'j', 'k', 'l'},
		6: {'m', 'n', 'o'},
		7: {'p', 'q', 'r', 's'},
		8: {'t', 'u', 'v'},
		9: {'w', 'x', 'y', 'z'},
	}
	// guaranteed digits contains nums 2-9 inclusive
	// say we had "2"
	first := digits[0] // byte val
	tmp := string(first)
	n, _ := strconv.Atoi(tmp)
	var res []string
	if n == 7 || n == 9 {
		res = make([]string, 4)
	} else {
		res = make([]string, 3)
	}

	for _, r := range digits {
		tmp := string(r)
		num, err := strconv.Atoi(tmp)
		if err != nil {
			return nil
		}
		possibilities := lookup[num] // 2-9. []byte
		for i, b := range possibilities {
			res[i] += string(b)
		}
	}

	return res
}

func letterCombinations(digits string) []string {
	// init a map of digit to possible letters
	// map[digit]slice of possible letters
	if len(digits) == 0 {
		return nil
	}
	lookup := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	res := []string{}
	var backtrack func(combo, nextDigits string)
	backtrack = func(combo, nextDigits string) {
		// goals choices, constraints on choices
		// goal: append/assign to res
		if len(nextDigits) == 0 {
			res = append(res, combo)
			return
		}
		curDigit := nextDigits[0]
		possibilities := lookup[curDigit]
		for i := range nextDigits {
			backtrack(combo+string(possibilities[i]), nextDigits[1:])
		}
	}
	backtrack("", digits)
	return res
}

func main() {
	digits := "23" // Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
	res := letterCombinations(digits)
	fmt.Printf("%+v\n", res)

}
