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

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;

public class InsertInterval57 {
    public static int[][] insert(int[][] intervals, int[] newInterval) {
        List<int[]> res = new ArrayList<>();

        // @formatter:off
        // Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
        // Output: [[1,5],[6,9]]
        // So @ each step in the iteration, we compare newInterval with
        // whatever is @ intervals[cur].
            // Case 1: prepend, so just add newInterval and the rest of list, then RETURN, since we did all the work and don't need to iterate more.
            // Case 2: res is prepopulated, NOW insert somewhere b/w [start, end] of res.
        // ie compare [1,3] and [2,5]. since you can merge, do so.
        // added to the list).

        // we need to merge any overlapping intervals (i.e., if prev[1] >= next[0], take
        // max of end times)
        int cur = 0;

        // what are our cases? (1) prepend, (2) append, (3) insert in between
        for (int i = 0; i < intervals.length; ++i) {
            var curInterval = intervals[cur];  // [1,3], new is [2,5]
            // (1): why would we prepend? 
            // if nextInterval was like [1, 50], curInterval being [55, 60]
            if (newInterval[1] < curInterval[0]) {
                // add the newInterval and insert the rest, starting with intervals[cur]
                res.add(newInterval);
                for (int current = cur; current < intervals.length; ++current) {
                    res.add(intervals[current]);
                }
                return res.stream().toArray(int[][]::new);
            } else {
                // do we need to merge? curInterval = [1,3], newInterval is [2,5]
                newInterval[0] = Math.min(curInterval[0], newInterval[0]);
                newInterval[1] = Math.max(curInterval[1], newInterval[1]);
                cur++;  // update so the current = cur inits the next interval
            }
        }
        // (3): append, since we know it wasn't added b/c of the early return
        res.add(intervals[cur]);

        return res.stream().toArray(int[][]::new);
    }

    public static void main(String[] args) {
        System.out.println();

        int[][] intervals = new int[][] { { 1, 3 }, { 6, 9 } };
        int[] newInterval = new int[] { 2, 5 };
        int[][] res = InsertInterval57.insert(intervals, newInterval);
        List<List<Integer>> resAsList = Arrays.stream(res).map(
                // then just ask, how to convert a normal int[]?
                intArray -> Arrays.stream(intArray).boxed().collect(Collectors.toList()))
                .collect(Collectors.toCollection(ArrayList::new));
        System.out.println(String.format("[[1,5],[6,9]] == %s", resAsList));

        int[][] intervals2 = new int[][]{{1,2},{3,5},{6,7},{8,10},{12,16}};
        int[] newInterval2 = new int[]{4,8};
        int[][] res2 = InsertInterval57.insert(intervals2, newInterval2);
        List<List<Integer>> res2AsList = Arrays.stream(res2).map(
            intArray -> Arrays.stream(intArray).boxed().collect(Collectors.toList())
        ).collect(Collectors.toCollection(ArrayList::new));
        System.out.println(String.format("[[1,2],[3,10],[12,16]] == %s", res2AsList));
    }
}
