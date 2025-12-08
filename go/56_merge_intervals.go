/*
Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.
*/

package main

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// [1,4][1,3][2,3]
	cur := 0
	res := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		curInterval := res[cur]
		nextInterval := intervals[i]
		// check for our merge cond. If mergeable, merge. if not move on
		// [1,100], [2,1000]
		if curInterval[1] >= nextInterval[0] { // merge. Handling `=` here: [1, 2], [2, 4] means we merge
			// "merging" means forgetting nextInterval's start ever existed
			curInterval[1] = max(curInterval[1], nextInterval[1]) // two end ranges
			// NOTE: cur unmodified
		} else { // no merge.
			// "move on" means advancing appending to result and advancing cur++
			res = append(res, nextInterval)
			cur++
		}
	}
	return res
}
