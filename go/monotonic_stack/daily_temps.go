/*
739. Daily Temperatures
Given an array of integers temperatures represents the daily temperatures, return an array answer such that answer[i] is the number of days you have to wait after the ith day to get a warmer temperature. If there is no future day for which this is possible, keep answer[i] == 0 instead.
*/

package main

import (
	"fmt"
	"os"
)

func dailyTemps(temperatures []int) []int {
	answer := make([]int, len(temperatures))
	s := []int{} // tracks indices, b/c of duplicate temps.
	// indices are what matter for daysWaited calculation anyway.

	// {73, 74, 75, 71, 69, 72, 76, 73}
	for i, temp := range temperatures {
		// BAD: s tracks indices. The top of the stack is meant to index into the temperatures slice.
		// for len(s) > 0 && temp > s[len(s)-1] {
		for len(s) > 0 && temp > temperatures[s[len(s)-1]] { // are you able to keep popping, since temp might be greater than more than one elt in the stac (i.e., temp > top && nonempty)?
			top := s[len(s)-1]
			s = s[:len(s)-1]
			daysWaited := i - top // i guaranteed > top
			// NOTE: a[i] == the number of days you have to wait *after* the ith day to get a warmer temperature.
			answer[top] = daysWaited
		}
		s = append(s, i) // any temp included in temps (in range), no check req
	}
	return answer
}

func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	answer := dailyTemps(temperatures)
	fmt.Fprintf(os.Stderr, "DEBUG[1]: daily_temps.go:18: answer=%+v\n", answer)
}
