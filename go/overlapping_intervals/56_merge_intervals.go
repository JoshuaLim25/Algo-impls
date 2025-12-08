/*
Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.

Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
*/

package main

import (
	"sort"
)

type Interval struct {
	Start int
	End int
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{intervals[0]}  // NOTE: after the sort
	curIdx := 0
	// [[1, 100], [2, 400]]
	for next := 1; next < len(intervals); next++ {
		curInterval := res[curIdx]
		nextInterval := intervals[next]
		// merge happens iff 
		if curInterval[1] >= nextInterval[0] {
			// merge op: max(cur and next end)
			curInterval[1] = max(curInterval[1], nextInterval[1])
		} else {
			curIdx++ // TODO: only advance when a new interval's been added
			res = append(res, nextInterval)
		}
	}

	return res
}

func coolMerge(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return intervals
	}
	res := []Interval{intervals[0]}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	// [[1, 100], [2, 3000]]
	// merge cond: cur interval's end is >= next interval's start
	curIdx := 0
	for next := 1; next < len(intervals); next++ {
		curInterval := res[curIdx]
		nextInterval := intervals[next]
		// merge cond
		if curInterval.End >= nextInterval.Start {
			// what does merging mean?
			// forget the next interval's start ever existed
			curInterval.End = max(nextInterval.End, curInterval.End)
		} else { // no merging
			res = append(res, nextInterval)
		}
		curIdx++ // TODO:
	}
	return res
}

func main() {
}
