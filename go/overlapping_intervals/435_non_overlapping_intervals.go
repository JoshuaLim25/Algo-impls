/*
Given an array of intervals intervals where intervals[i] = [starti, endi], return the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.

Note that intervals which only touch at a point are non-overlapping. For example, [1, 2] and [2, 3] are non-overlapping.
*/

package main

import (
	"sort"
	"fmt"
	"os"
)

func eraseOverlapIntervals(intervals [][]int) int {
	// KEY: sort by END TIME, unlike for merging intervals
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	// [[1 2] [2 3] [1 3] [3 4]]
	erased := 0
	cur := 0
	for i := 1; i < len(intervals); i++ {
		curInterval := intervals[cur]
		nextInterval := intervals[i]
		if curInterval[1] > nextInterval[0] {  // = means keep
			erased++
			// NOTE: don't update cur, only NO overlap
		} else {
			// no removals happened, meaning no overlap
			cur = i  // NOTE: NOT cur++. Take [[1,2], [1, 3], [1,7], [2, 3]]. You'd want cur to be on [1,7]
		}
	}
	return erased

}

// func eraseOverlapIntervals(intervals [][]int) int {
// 	// NOTE: = case is non-overlapping (so keep it)
// 	// 1. sort by start of the interval
// 	// 2. Range over intervals
// 		// 3. if it meets the overlapping criteria, erased++
// 		// 3b. if not, just keep going
// 	// 4. return erased
// 	// In this case, to be overlapping means that the **end** of the interval we're on is *strictly* > the **start** of the next one
// 	if len(intervals) == 0 {
// 		return 0
// 	}
// 	// [[1,2],[2,3],[3,4],[1,3]]
// 	fmt.Printf("%+v\n", intervals)
// 	sort.Slice(intervals, func(i, j int) bool {
// 		return intervals[i][0] < intervals[j][0] 
// 	})
// 	fmt.Printf("%+v\n", intervals)
// 	cur := 0
// 	erased := 0
// 	for next := 1; next < len(intervals); next++ {
// 		curInterval := intervals[cur]
// 		nextInterval := intervals[next]
// 		if curInterval[0] == nextInterval[0] {  // [1,2], [1,3]
// 			if curInterval[1] >= nextInterval[1] {
// 				erased++
// 				// NOTE: no cur update
// 				continue
// 			} else {
// 				erased++
// 				cur++
// 				continue
// 			}
// 		}
// 		if curInterval[1] > nextInterval[0] { // equals doesn't qualify
// 			erased++
// 		}
// 		cur++
// 		// getting here means no removals necessary
// 		// meaning you 
// 	}
//
// 	return erased
// }

func main() {
	intervals := [][]int{[]int{1,2},[]int{2,3},[]int{3,4},[]int{1,3}}
	erased := eraseOverlapIntervals(intervals)
	fmt.Fprintf(os.Stderr, "DEBUG[1]: 435_non_overlapping_intervals.go:55: erased=%+v\n", erased)
}
