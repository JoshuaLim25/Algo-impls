/*
Given an array of integers temperatures represents the daily temperatures, return an array answer such that answer[i] is the number of days you have to wait after the ith day to get a warmer temperature. If there is no future day for which this is possible, keep answer[i] == 0 instead.


[21, 34, 89, 88, 89, 90, 67]
*/
package main

func dailyTemps(temperatures []int) []int {
	if len(temperatures) == 0 {
		return []int{}
	}
	// NOTE: want strictly greater temps, <= s.top means just add
	s := []int{}       // indices of progressively decreasing temps
	res := make([]int, len(temperatures))  // no need to range: [ 0 0 0 0 0 0 ]
	for i, temp := range temperatures {
		for len(s) > 0 && temp > temperatures[s[len(s)-1]] {
			// pop
			top := s[len(s)-1]
			s = s[:len(s)-1]

			// NOTE: not assigning temps, daysWaited
			daysWaited := i - top
			res[top] = daysWaited
		}
		s = append(s, i)  // TODO: conditional?
	}
	return res
}



