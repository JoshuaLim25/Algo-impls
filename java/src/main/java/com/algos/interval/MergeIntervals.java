package com.algos.interval;

import java.util.ArrayList;
import java.util.Comparator;
import java.util.List;

/*
Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
 */
public class MergeIntervals {
    public static List<List<Integer>> merge(List<List<Integer>> intervals) {
        List<List<Integer>> res = new ArrayList<>();
        if (intervals.size() <= 1) return intervals;

        // intervals.sort((a, b) -> Integer.compare(a.get(0), b.get(0)));
        intervals.sort(Comparator.comparingInt((a) -> a.get(0)));

        int cur = 0;
        res.add(intervals.get(cur));

        for (int i = 1; i < intervals.size(); ++i) {
            var curInterval = res.get(cur);
            var nextInterval = intervals.get(i);

            if (curInterval.get(1) >= nextInterval.get(0)) {
                curInterval.set(1, Math.max(curInterval.get(1), nextInterval.get(1)));
            } else {
                res.add(nextInterval);
                cur++;
            }
        }
        return res;
    }

    public static void main(String[] args) {
        List<List<Integer>> intervals = new ArrayList<>(List.of(
                new ArrayList<>(List.of(1, 3)),
                new ArrayList<>(List.of(2, 6)),
                new ArrayList<>(List.of(8, 10)),
                new ArrayList<>(List.of(15, 18))
        ));
        var res = MergeIntervals.merge(intervals);
        System.out.println("Expected [[1,6],[8,10],[15,18]], got " + res);
        // Output: [[1,6],[8,10],[15,18]]

        List<List<Integer>> singleInterval = new ArrayList<>(List.of(List.of(1, 3)));
        var resSingle = MergeIntervals.merge(singleInterval);
        System.out.println("Expected [[1,3]], got " + resSingle);
        // Output: [[1,3]]
    }
}



/*
for if you need to work with int[][]:
int[][] s = new int[100][];
s[99] = new int[]{1};
System.err.println(String.format("DEBUG[1]: ContainerWithMostWater11.java:11: s=%s, s[99] = %s", Arrays.toString(s), Arrays.toString(s[99])));
 */
