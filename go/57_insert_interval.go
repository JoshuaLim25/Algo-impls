/*
57. Insert Interval - Medium

You are given an array of non-overlapping intervals intervals where intervals[i] = [starti, endi] represent the start and the end of the ith interval and intervals is sorted in ascending order by starti. You are also given an interval newInterval = [start, end] that represents the start and end of another interval.

Insert newInterval into intervals such that intervals is still sorted in ascending order by starti and intervals still does not have any overlapping intervals (merge overlapping intervals if necessary).

Return intervals after the insertion.

Note that you don't need to modify intervals in-place. You can make a new array and return it.



Example 1:

Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
Output: [[1,5],[6,9]]
*/

package main

import (
	"fmt"
	"os"
)

func insert(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{}
	for i, curInterval := range intervals {
		// first two are for nonoverlapping insert
		if newInterval[1] < curInterval[0] {
			res = append(res, newInterval)
			res = append(res, intervals[i:]...) // rest guaranteed >
			return res
		} else if newInterval[0] > curInterval[1] {
			res = append(res, curInterval)
			// NOTE: newInterval could overlap with future ones
		} else {
			// is overlapping, so "merge"
			newInterval[0] = min(newInterval[0], curInterval[0])
			newInterval[1] = max(newInterval[1], curInterval[1])
		}
	}
	// NOTE: 2 ways to exit: either via first case (which adds interval && returns early) or by way of merging and continuing (which doesn't add)
	res = append(res, newInterval)

	return res
}

func main() {
	intervals := [][]int{{1, 3}, {6, 9}}
	newInterval := []int{2, 5} // Output: [[1,5],[6,9]]
	output := insert(intervals, newInterval)
	fmt.Fprintf(os.Stderr, "DEBUG[1]: 57_insert_interval.go:48: output=%+v\n", output)

}
